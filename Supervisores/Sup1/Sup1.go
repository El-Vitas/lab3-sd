package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	hexpb "Supervisores/generated/hex_s"

	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:50054" // Dirección del servidor Hextech
)

type Supervisor struct {
	id                   int
	lastKnownVectorClock map[string][]int32
}

// connectToHextechServer establece conexión con el servidor Hextech
func connectToHextechServer(address string) (*grpc.ClientConn, hexpb.HexSClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("no se pudo conectar al servidor Hextech: %v", err)
	}
	client := hexpb.NewHexSClient(conn)
	return conn, client, nil
}

// sendCommand envía un comando al Servidor Hextech
func (s *Supervisor) sendCommand(serverAddress string, command string, args []string) error {
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
		id:                   1,
		lastKnownVectorClock: make(map[string][]int32),
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
			supervisor.sendCommand(serverAddress, command, []string{region, product, newValue})
		case "RenombrarProducto":
			fmt.Println("Ingrese <Región> <Producto> <Nuevo Nombre>:")
			fmt.Scanln(&region, &product, &newValue)
			supervisor.sendCommand(serverAddress, command, []string{region, product, newValue})
		case "ActualizarValor":
			fmt.Println("Ingrese <Región> <Producto> <Nuevo Valor>:")
			fmt.Scanln(&region, &product, &newValue)
			supervisor.sendCommand(serverAddress, command, []string{region, product, newValue})
		case "BorrarProducto":
			fmt.Println("Ingrese <Región> <Producto>:")
			fmt.Scanln(&region, &product)
			supervisor.sendCommand(serverAddress, command, []string{region, product})
		default:
			fmt.Println("Comando no reconocido.")
		}

		time.Sleep(1 * time.Second)
	}
}
