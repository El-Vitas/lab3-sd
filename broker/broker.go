package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"time"

	broker "broker/generated/broker-s"

	"google.golang.org/grpc"
	// Importa el paquete generado por el archivo .proto
)

// Implementación del servidor Broker
type BrokerServer struct {
	broker.UnimplementedBrokerServer
}

var servers = []string{"localhost:50051", "localhost:50052", "localhost:50053"}

// Maneja la solicitud RouteCommand
func (s *BrokerServer) RouteCommand(ctx context.Context, req *broker.CommandRequest) (*broker.CommandResponse, error) {
	// Aquí se pueden manejar los comandos específicos, por ejemplo:

	// Selecciona un servidor aleatorio
	rand.Seed(time.Now().UnixNano())
	serverIndex := rand.Intn(len(servers))

	// Devuelve la dirección del servidor seleccionado
	response := &broker.CommandResponse{
		ServerAddress: servers[serverIndex],
	}

	return response, nil

}

func main() {
	// Configura el servidor gRPC
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	// Crea el servidor gRPC
	s := grpc.NewServer()

	// Registra el servicio Broker
	brokerServer := &BrokerServer{}
	broker.RegisterBrokerServer(s, brokerServer)

	// Inicia el servidor
	log.Println("Servidor Broker corriendo en el puerto :50054")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
