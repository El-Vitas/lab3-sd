package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	hextechpb "path/to/generated/HextechServer"
	brokerpb "path/to/generated/broker-s"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Jayce maintains known data and ensures Monotonic Reads.
type Jayce struct {
	knownData      map[string]ProductData // Known data for each product
	lastServerUsed map[string]string      // Last server used for each product/region
	brokerClient   brokerpb.BrokerClient  // gRPC client to interact with the Broker
}

type ProductData struct {
	Quantity    int32
	VectorClock []int32
}

// Connects to the Broker and returns a gRPC client
func connectToBroker(brokerAddress string) (*grpc.ClientConn, brokerpb.BrokerClient, error) {
	conn, err := grpc.Dial(brokerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Broker: %v", err)
	}
	client := brokerpb.NewBrokerClient(conn)
	return conn, client, nil
}

// Connects to a Hextech server
func connectToHextech(serverAddress string) (*grpc.ClientConn, hextechpb.HextechServerClient, error) {
	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to Hextech server: %v", err)
	}
	client := hextechpb.NewHextechServerClient(conn)
	return conn, client, nil
}

// Checks if the new vector clock is more recent or equal to the last known clock
func (j *Jayce) isMoreRecent(newClock, knownClock []int32) bool {
	for i := range newClock {
		if newClock[i] < knownClock[i] {
			return false
		}
	}
	return true
}

// Fetches the product data ensuring Monotonic Reads
func (j *Jayce) fetchProduct(brokerAddress, region, product string) error {
	// Get server address from the Broker
	serverAddress, err := j.getServerAddress(brokerAddress, region, product)
	if err != nil {
		return err
	}

	// Connect to the Hextech server
	connHextech, hextechClient, err := connectToHextech(serverAddress)
	if err != nil {
		return err
	}
	defer connHextech.Close()

	// Request the product data
	resp, err := hextechClient.GetData(context.Background(), &hextechpb.GetDataRequest{Key: product})
	if err != nil {
		return fmt.Errorf("error fetching product data from Hextech server: %v", err)
	}

	// Ensure Monotonic Reads by checking the vector clock
	key := region + ":" + product
	if known, exists := j.knownData[key]; exists {
		if !j.isMoreRecent(resp.VectorClock, known.VectorClock) {
			return fmt.Errorf("inconsistent data received; Monotonic Reads violated")
		}
	}

	// Update Jayce's known data
	j.knownData[key] = ProductData{
		Quantity:    resp.Quantity,
		VectorClock: resp.VectorClock,
	}
	j.lastServerUsed[key] = serverAddress

	fmt.Printf("Product: %s, Region: %s, Quantity: %d, VectorClock: %v\n",
		product, region, resp.Quantity, resp.VectorClock)
	return nil
}

// Gets the server address for a product from the Broker
func (j *Jayce) getServerAddress(brokerAddress, region, product string) (string, error) {
	request := &brokerpb.CommandRequest{
		Command: "GetServer",
		Args:    []string{region, product},
	}

	// Call RouteCommand on the Broker
	response, err := j.brokerClient.RouteCommand(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf("error getting server address from Broker: %v", err)
	}

	return response.ServerAddress, nil
}

func main() {
	// Address of the Broker
	brokerAddress := "localhost:50052"

	// Connect to the Broker
	connBroker, brokerClient, err := connectToBroker(brokerAddress)
	if err != nil {
		log.Fatalf("Failed to connect to Broker: %v", err)
	}
	defer connBroker.Close()

	// Initialize Jayce
	jayce := &Jayce{
		knownData:      make(map[string]ProductData),
		lastServerUsed: make(map[string]string),
		brokerClient:   brokerClient,
	}

	// Simulate product queries
	rand.Seed(time.Now().UnixNano())
	for {
		fmt.Println("Enter a region and product to query:")
		var region, product string
		fmt.Scanln(&region, &product)

		err := jayce.fetchProduct(brokerAddress, region, product)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		time.Sleep(2 * time.Second) // Simulate delay between queries
	}
}
