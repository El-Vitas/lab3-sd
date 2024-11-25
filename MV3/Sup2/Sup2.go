package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	brokerpb "Sup2/generated/broker-s"
	hexpb "Sup2/generated/hex_s"

	"google.golang.org/grpc"
)

const (
	brokerAddress = "dist101:50054" // Dirección del servidor Broker
)

type Supervisor struct {
	id                   int
	lastKnownVectorClock map[string][]int32
}

// connectToBroker establece conexión con el servidor Broker
func connectToBroker(address string) (*grpc.ClientConn, brokerpb.BrokerClient, error) {
	fmt.Printf("Conectando al servidor Broker en %s...\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("no se pudo conectar al broker: %v", err)
	}
	client := brokerpb.NewBrokerClient(conn)
	return conn, client, nil
}

// getServerAddress obtiene la dirección de un servidor Hextech a través del broker
func getServerAddress(brokerAddress string) (string, error) {
	conn, brokerClient, err := connectToBroker(brokerAddress)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	// Realiza la solicitud al broker
	fmt.Println("Obteniendo dirección del servidor Hextech desde el broker...")
	response, err := brokerClient.RouteCommand(context.Background(), &brokerpb.CommandRequest{})
	if err != nil {
		return "", fmt.Errorf("error al obtener dirección del servidor desde el broker: %v", err)
	}
	return response.ServerAddress, nil
}

// connectToHextechServer establece conexión con el servidor Hextech
func connectToHextechServer(address string) (*grpc.ClientConn, hexpb.HexSClient, error) {
	fmt.Printf("Conectando al servidor Hextech en %s...\n", address)
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("no se pudo conectar al servidor Hextech: %v", err)
	}
	client := hexpb.NewHexSClient(conn)
	return conn, client, nil
}

// sendCommand envía un comando al Servidor Hextech
func (s *Supervisor) sendCommand(brokerAddress, command string, args []string) error {
	// Obtener la dirección del servidor a través del broker
	serverAddress, err := getServerAddress(brokerAddress)
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

	var vectorClock []int32

	// Manejamos cada comando por separado
	switch command {
	case "AgregarProducto":
		request := &hexpb.AddRecordRequest{
			Region:  args[0],
			Product: args[1],
			Value:   args[2],
		}
		fmt.Printf("Supervisor %d envía comando '%s' a Hextech: %v\n", s.id, command, request)
		response, err := hexClient.AddRecord(context.Background(), request)
		if err != nil {
			return fmt.Errorf("error al agregar producto: %v", err)
		}
		vectorClock = response.VectorClock

	case "RenombrarProducto":
		request := &hexpb.RenameRecordRequest{
			Region:     args[0],
			OldProduct: args[1],
			NewProduct: args[2],
		}
		fmt.Printf("Supervisor %d envía comando '%s' a Hextech: %v\n", s.id, command, request)
		response, err := hexClient.RenameRecord(context.Background(), request)
		if err != nil {
			return fmt.Errorf("error al renombrar producto: %v", err)
		}
		vectorClock = response.VectorClock

	case "ActualizarValor":
		request := &hexpb.UpdateRecordRequest{
			Region:   args[0],
			Product:  args[1],
			NewValue: args[2],
		}
		fmt.Printf("Supervisor %d envía comando '%s' a Hextech: %v\n", s.id, command, request)
		response, err := hexClient.UpdateRecordValue(context.Background(), request)
		if err != nil {
			return fmt.Errorf("error al actualizar valor: %v", err)
		}
		vectorClock = response.VectorClock

	case "BorrarProducto":
		request := &hexpb.DeleteRecordRequest{
			Region:  args[0],
			Product: args[1],
		}
		fmt.Printf("Supervisor %d envía comando '%s' a Hextech: %v\n", s.id, command, request)
		response, err := hexClient.DeleteRecord(context.Background(), request)
		if err != nil {
			return fmt.Errorf("error al borrar producto: %v", err)
		}
		vectorClock = response.VectorClock

	default:
		return fmt.Errorf("comando no reconocido: %s", command)
	}

	// Actualizar reloj vectorial
	s.lastKnownVectorClock[args[0]] = vectorClock
	fmt.Printf("Supervisor %d ejecutó comando '%s'. Reloj vectorial actualizado: %v\n", s.id, command, vectorClock)
	return nil
}

func main() {
	supervisor := &Supervisor{
		id:                   2,
		lastKnownVectorClock: make(map[string][]int32),
	}

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Supervisor Hexgate 2 iniciado.")

	for {
		fmt.Println("\nSupervisor 2: Ingrese un comando (AgregarProducto, RenombrarProducto, ActualizarValor, BorrarProducto):")
		var command, region, product, newValue string
		fmt.Scanln(&command)

		switch command {
		case "AgregarProducto":
			fmt.Println("Ingrese <Región> <Producto> <Cantidad (opcional)>:")
			fmt.Scanln(&region, &product, &newValue)
			if newValue == "" {
				newValue = "0"
			}
			supervisor.sendCommand(brokerAddress, command, []string{region, product, newValue})
		case "RenombrarProducto":
			fmt.Println("Ingrese <Región> <Producto> <Nuevo Nombre>:")
			fmt.Scanln(&region, &product, &newValue)
			supervisor.sendCommand(brokerAddress, command, []string{region, product, newValue})
		case "ActualizarValor":
			fmt.Println("Ingrese <Región> <Producto> <Nuevo Valor>:")
			fmt.Scanln(&region, &product, &newValue)
			supervisor.sendCommand(brokerAddress, command, []string{region, product, newValue})
		case "BorrarProducto":
			fmt.Println("Ingrese <Región> <Producto>:")
			fmt.Scanln(&region, &product)
			supervisor.sendCommand(brokerAddress, command, []string{region, product})
		default:
			fmt.Println("Comando no reconocido.")
		}

		time.Sleep(1 * time.Second)
	}
}
