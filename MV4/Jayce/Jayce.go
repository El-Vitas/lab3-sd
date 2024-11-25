package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	brokerpb "Jayce/generated/broker-s"
	hexpb "Jayce/generated/hex_s"

	"google.golang.org/grpc"
)

type Jayce struct {
	knownData      map[string]ProductData // Known data for each product
	lastServerUsed map[string]string      // Last server used for each product/region
	brokerClient   brokerpb.BrokerClient  // gRPC client to interact with the Broker
}

type ProductData struct {
	Quantity    int32
	VectorClock []int32
}

func connectToBroker(address string) (*grpc.ClientConn, brokerpb.BrokerClient, error) {
	fmt.Printf("Conectando al servidor Broker en %s...\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("no se pudo conectar al broker: %v", err)
	}
	client := brokerpb.NewBrokerClient(conn)
	return conn, client, nil
}

func connectToHextechServer(address string) (*grpc.ClientConn, hexpb.HexSClient, error) {
	fmt.Printf("Conectando al servidor Hextech en %s...\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("no se pudo conectar al servidor Hextech: %v", err)
	}
	client := hexpb.NewHexSClient(conn)
	return conn, client, nil
}

func getServerAddress(brokerAddress, region, product string) (string, error) {
	conn, brokerClient, err := connectToBroker(brokerAddress)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Realiza la solicitud al broker
	request := &brokerpb.CommandRequest{
		Command: "GetServer",
		Args:    []string{region, product},
	}
	fmt.Println("Obteniendo dirección del servidor Hextech desde el broker...")
	response, err := brokerClient.RouteCommand(context.Background(), request)
	if err != nil {
		return "", fmt.Errorf("error al obtener dirección del servidor desde el broker: %v", err)
	}
	return response.ServerAddress, nil
}

func (j *Jayce) sendCommand(brokerAddress, region, product string) error {
	// Obtener la dirección del servidor a través del broker
	serverAddress, err := getServerAddress(brokerAddress, region, product)
	if err != nil {
		return err
	}
	fmt.Printf("Servidor asignado por el broker: %s\n", serverAddress)

	// Conectarse al servidor Hextech
	conn, hexClient, err := connectToHextechServer(serverAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	//var vectorClock []int32

	request := &hexpb.GetDataRequest{
		Region:  region,
		Product: product,
	}

	resp, err := hexClient.GetData(context.Background(), request)
	if err != nil {
		fmt.Println("Producto no encontrado")
		return nil
	}

	key := region + ":" + product
	/*if known, exists := j.knownData[key]; exists {
		if !j.isMoreRecent(resp.VectorClock, known.VectorClock) {
			return fmt.Errorf("inconsistent data received; Monotonic Reads violated")
		}
	}*/

	j.knownData[key] = ProductData{
		Quantity:    resp.Quantity,
		VectorClock: resp.VectorClock,
	}
	j.lastServerUsed[key] = serverAddress

	fmt.Printf("Product: %s, Region: %s, Quantity: %d, VectorClock: %v\n",
		product, region, resp.Quantity, resp.VectorClock)

	return nil
}

func main() {
	// Address of the Broker
	brokerAddress := "dist101:50054"

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
		fmt.Println("Ingresa una región y un producto:")
		fmt.Println("<Region> <Producto>")
		var region, product string
		fmt.Scanln(&region, &product)

		err := jayce.sendCommand(brokerAddress, region, product)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}

		time.Sleep(2 * time.Second) // Simulate delay between queries
	}
}
