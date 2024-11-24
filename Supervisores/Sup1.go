package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	brokerpb "C:\Users\sebas\Desktop\lab3-sd\proto\generated\Broker" // Ajusta esta ruta según tu configuración
	hextechpb "C:\Users\sebas\Desktop\lab3-sd\proto\generated\HextechServer"
	"google.golang.org/grpc"
)

const (
	brokerAddress1 = "localhost:50051" // Dirección del Broker
)

type Supervisor struct {
	id                   int
	lastKnownVectorClock map[string][]int
}

// connectToBroker establece conexión con el Broker
func connectToBroker(brokerAddress string) (*grpc.ClientConn, brokerpb.BrokerClient, error) {
	conn, err := grpc.Dial(brokerAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("no se pudo conectar al broker: %v", err)
	}
	client := brokerpb.NewBrokerClient(conn)
	return conn, client, nil
}

// sendCommand envía un comando al Broker y al Servidor Hextech
func (s *Supervisor) sendCommand(brokerAddress string, command string, args []string) error {
	conn, brokerClient, err := connectToBroker(brokerAddress)
	if err != nil {
		return err
	}
	defer conn.Close()

	request := &brokerpb.CommandRequest{
		Command: command,
		Args:    args,
	}
	// Enviar al Broker
	response, err := brokerClient.RouteCommand(context.Background(), request)
	if err != nil {
		return fmt.Errorf("error al enviar el comando al broker: %v", err)
	}

	serverAddress := response.ServerAddress
	connHextech, err := grpc.Dial(serverAddress, grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("no se pudo conectar al servidor Hextech: %v", err)
	}
	defer connHextech.Close()

	hextechClient := hextechpb.NewHextechServerClient(connHextech)
	resp, err := hextechClient.ExecuteCommand(context.Background(), request)
	if err != nil {
		return fmt.Errorf("error al ejecutar el comando en el servidor Hextech: %v", err)
	}

	s.lastKnownVectorClock[args[0]] = resp.VectorClock
	fmt.Printf("Supervisor %d ejecutó comando en %s. Reloj vectorial actualizado: %v\n", s.id, serverAddress, resp.VectorClock)
	return nil
}

func main() {
	supervisor := &Supervisor{
		id:                   1,
		lastKnownVectorClock: make(map[string][]int),
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Supervisor Hexgate 1 iniciado.")

	for {
		fmt.Println("\nSupervisor 1: Ingrese un comando (AgregarProducto, RenombrarProducto, ActualizarValor, BorrarProducto):")
		var command, region, product, newValue string
		fmt.Scanln(&command)

		switch command {
		case "AgregarProducto":
			fmt.Println("Ingrese <Región> <Producto> <Cantidad (opcional)>:")
			fmt.Scanln(&region, &product, &newValue)
			if newValue == "" {
				newValue = "0"
			}
			supervisor.sendCommand(brokerAddress1, command, []string{region, product, newValue})
		case "RenombrarProducto":
			fmt.Println("Ingrese <Región> <Producto> <Nuevo Nombre>:")
			fmt.Scanln(&region, &product, &newValue)
			supervisor.sendCommand(brokerAddress1, command, []string{region, product, newValue})
		case "ActualizarValor":
			fmt.Println("Ingrese <Región> <Producto> <Nuevo Valor>:")
			fmt.Scanln(&region, &product, &newValue)
			supervisor.sendCommand(brokerAddress1, command, []string{region, product, newValue})
		case "BorrarProducto":
			fmt.Println("Ingrese <Región> <Producto>:")
			fmt.Scanln(&region, &product)
			supervisor.sendCommand(brokerAddress1, command, []string{region, product})
		default:
			fmt.Println("Comando no reconocido.")
		}

		time.Sleep(1 * time.Second)
	}
}