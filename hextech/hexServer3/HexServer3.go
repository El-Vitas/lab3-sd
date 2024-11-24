package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"
	"time"

	pb2 "hexServer3/generated/hex_s"
	pb "hexServer3/generated/hextech"

	"google.golang.org/grpc"
)

// Estructura para representar el servidor Hextech
type server struct {
	pb.UnimplementedHextechServer
	mu       sync.Mutex
	regions  map[string][]int    // Mapa de regiones con sus relojes vectoriales
	logs     map[string][]string // Mapa para almacenar logs de cada región
	stopChan chan struct{}       // Canal para detener el servidor
	pb2.UnimplementedHexSServer
}

var servers = []string{"localhost:50051", "localhost:50052"}
var numServer = 2

// Inicializa el servidor Hextech
func newServer() *server {
	return &server{
		regions:  make(map[string][]int),
		logs:     make(map[string][]string),
		stopChan: make(chan struct{}),
	}
}

func readFileAsLines(fileName string) ([]string, error) {
	var lines []string
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

// Función para propagar cambios a otros servidores
func (s *server) propagateChanges() {
	for {
		select {
		case <-s.stopChan:
			return
		default:
			// Simula la propagación de cambios a otros servidores
			fmt.Printf("Propagando cambios a otros servidores...\n")
			for _, addr := range servers {
				conn, err := grpc.Dial(addr, grpc.WithInsecure())
				if err != nil {
					log.Printf("Error al conectar con el servidor %s: %v", addr, err)
					continue
				}
				client := pb.NewHextechClient(conn)
				fmt.Print("Propagando a ", addr, "\n")
				// Propagar los cambios de todas las regiones
				for region, reloj := range s.regions {
					s.mu.Lock()
					changes := s.logs[region]
					s.mu.Unlock()

					_, err := client.ReceiveChanges(context.Background(), &pb.ChangesRequest{
						Region:     region,
						LocalClock: convertToInt32Slice(reloj),
						Changes:    changes,
					})
					if err != nil {
						log.Printf("Error al enviar cambios a %s para la región %s: %v", addr, region, err)

					}
				}
				conn.Close()
			}

			s.logs = make(map[string][]string) // Limpiar los logs después de propagar
			time.Sleep(30 * time.Second)       // Intervalo entre propagaciones
		}
	}
}

func (s *server) propagateMerge(region string) {
	fmt.Print("Propagando merge a otros servidores...\n")
	for _, addr := range servers {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Printf("Error al conectar con el servidor %s: %v", addr, err)
			continue
		}
		client := pb.NewHextechClient(conn)

		// Propagar el merge solo para la región pasada como parámetro
		s.mu.Lock()
		localClock := s.regions[region]
		s.mu.Unlock()

		fmt.Printf("Propagando merge a %s para la reg %s\n", addr, region)
		// Realizar merge con el servidor remoto
		lines, err := readFileAsLines(region + ".txt")
		if err != nil {
			log.Printf("Error al leer el archivo %s: %v", region, err)
		}
		_, err = client.ReceiveFile(context.Background(), &pb.ReceiveFileRequest{
			Region:     region,
			LocalClock: convertToInt32Slice(localClock),
			Text:       lines,
		})
		if err != nil {
			log.Printf("Error al realizar merge con %s para la región %s: %v", addr, region, err)
		}
		s.logs[region] = make([]string, 0) // Limpiar los logs después de propagar
		conn.Close()
	}
}

// Función para recibir el archivo y actualizar el reloj
func (s *server) ReceiveFile(ctx context.Context, req *pb.ReceiveFileRequest) (*pb.ReceiveFileResponse, error) {
	// Bloquear acceso concurrente
	s.mu.Lock()
	defer s.mu.Unlock()

	// Extraer la región, el reloj y el contenido del archivo del request
	region := req.Region
	remoteClock := req.LocalClock
	fileContent := req.Text // Suponiendo que Text es el contenido del archivo como un string

	fmt.Printf("Recibiendo archivo de merge para la región %s\n", region)
	// Actualizar el archivo de la región
	fileName := fmt.Sprintf("%s.txt", region)
	content := strings.Join(fileContent, "\n")
	err := os.WriteFile(fileName, []byte(content), 0644) // Crear o sobrescribir el archivo
	if err != nil {
		log.Printf("Error al escribir el archivo para la región %s: %v", region, err)
		return nil, fmt.Errorf("error al escribir el archivo: %v", err)
	}

	// Actualizar el reloj de la región
	s.regions[region] = convertToIntSlice(remoteClock)

	fmt.Printf("Archivo recibido y reloj actualizado para la región %s\n", region)
	fmt.Printf("Numero de reloj: %v\n", s.regions[region])

	s.logs[region] = make([]string, 0) // Limpiar los logs después de recibir el archivo
	return &pb.ReceiveFileResponse{
		Region:      region,
		VectorClock: convertToInt32Slice(s.regions[region]),
	}, nil
}

func (s *server) ReceiveChanges(ctx context.Context, req *pb.ChangesRequest) (*pb.ChangesResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	region := req.Region
	remoteClock := req.LocalClock
	changes := req.Changes

	localClock := s.getOrCreateRegion(region)

	// Detectar conflictos en el reloj vectorial
	conflict := false
	cont := 0
	for i := range localClock {
		if remoteClock[i] > int32(localClock[i]) {
			if cont < 0 {
				conflict = true
				break
			}
			cont++
		} else if remoteClock[i] < int32(localClock[i]) {
			if cont > 0 {
				conflict = true
				break
			}
			cont--
		}
	}

	if conflict {
		fmt.Printf("Conflicto detectado en la región %s\n", region)
		// // Realizar merge si hay conflictos
		// conn, err := grpc.Dial(dominante, grpc.WithTransportCredentials(insecure.NewCredentials()))
		// if err != nil {
		// 	log.Fatalf("Error al conectar con el nodo dominante: %v", err)
		// }
		// defer conn.Close()
		// client := pb.NewHextechClient(conn) // Reemplaza con el nombre adecuado del cliente
		// _, err = client.MergeChanges(context.Background(), &pb.MergeChangesRequest{
		// 	Region:      region,
		// 	LocalClockA: convertToInt32Slice(localClock),
		// 	LocalClockB: remoteClock,
		// 	LogA:        s.logs[region],
		// 	LogB:        changes,
		// })
		// if err != nil {
		// 	log.Printf("Error al realizar merge con el nodo dominante: %v", err)
		// }

		_, err := s.MergeChanges(context.Background(), &pb.MergeChangesRequest{
			Region:      region,
			LocalClockA: convertToInt32Slice(localClock),
			LocalClockB: remoteClock,
			LogA:        s.logs[region],
			LogB:        changes,
		})
		if err != nil {
			log.Printf("Error al realizar merge con el nodo dominante: %v", err)
		}
	} else if cont > 0 { // Si el reloj remoto es mayor aplicamos lo cambios directamente, caso contrario ignoramos
		fmt.Printf("Aplicando cambios directamente en la región %s\n", region)
		// Aplicar cambios directamente si no hay conflictos
		for range changes {
			s.applyChange(region, remoteClock, changes)
		}
	}

	return &pb.ChangesResponse{Acknowledged: true}, nil
}

// Función para verificar si un producto existe en el archivo de registros
func (s *server) recordExists(region string, record string) bool {
	// Abrir el archivo correspondiente a la región
	fileName := fmt.Sprintf("%s.txt", region) // Cambia la ruta de acuerdo a tu estructura

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error al abrir el archivo para la región %s: %v", region, err)
		return false
	}
	defer file.Close()

	// Leer línea por línea del archivo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, record) { // Buscar si el registro (producto) está en la línea
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error al leer el archivo para la región %s: %v", region, err)
	}

	return false
}

// Función para hacer merge en caso de conflictos

func (s *server) MergeChanges(ctx context.Context, req *pb.MergeChangesRequest) (*pb.MergeChangesResponse, error) {
	region := req.Region
	localClockA := req.LocalClockA
	localClockB := req.LocalClockB
	logA := req.LogA
	logB := req.LogB

	log.Printf("Realizando merge para la región %s", region)

	if len(s.regions[region]) == 0 {
		// Inicializar el slice con el tamaño necesario si está vacío
		s.regions[region] = make([]int, len(localClockA))
	}

	// Realizar el merge de los relojes vectoriales
	for i := range localClockA {
		if localClockB[i] > localClockA[i] {
			if localClockB[i] > int32(s.regions[region][i]) {
				s.regions[region][i] = int(localClockB[i])
			}
		} else {
			if localClockA[i] > int32(s.regions[region][i]) {
				s.regions[region][i] = int(localClockA[i])
			}
		}
	}

	// Procesar los cambios de A
	if len(logA) > 0 { // Validar si logA no está vacío
		for _, change := range logA {
			parts := strings.Fields(change)
			if len(parts) < 2 {
				log.Printf("Cambio inválido en la región %s: %s", region, change)
				continue
			}

			action := parts[0]
			switch action {
			case "AgregarProducto":
				if len(parts) < 3 {
					log.Printf("Cambio inválido para 'AgregarProducto' en la región %s: %s", region, change)
					continue
				}
				record := strings.Join(parts[2:], " ")
				if !s.recordExists(region, record) {
					err := add(region, record)
					if err != nil {
						log.Printf("Error al agregar registro en la región %s: %v", region, err)
					}
				}

			case "BorrarProducto":
				if len(parts) < 3 {
					log.Printf("Cambio inválido para 'BorrarProducto' en la región %s: %s", region, change)
					continue
				}
				product := parts[2]
				if s.recordExists(region, product) {
					err := delete(region, product)
					if err != nil {
						log.Printf("Error al borrar producto en la región %s: %v", region, err)
					}
				} else {
					log.Printf("El producto no existe en la región %s: %s", region, product)
				}
			}
		}
	} else {
		log.Printf("No se encontraron cambios en logA para la región %s", region)
	}

	// Procesar los cambios de B
	if len(logB) > 0 { // Validar si logB no está vacío
		for _, change := range logB {
			parts := strings.Fields(change)
			if len(parts) < 2 {
				log.Printf("Cambio inválido en la región %s: %s", region, change)
				continue
			}

			action := parts[0]
			switch action {
			case "AgregarProducto":
				if len(parts) < 3 {
					log.Printf("Cambio inválido para 'AgregarProducto' en la región %s: %s", region, change)
					continue
				}
				record := strings.Join(parts[2:], " ")
				if !s.recordExists(region, record) {
					err := add(region, record)
					if err != nil {
						log.Printf("Error al agregar registro en la región %s: %v", region, err)
					}
				}

			case "BorrarProducto":
				if len(parts) < 3 {
					log.Printf("Cambio inválido para 'BorrarProducto' en la región %s: %s", region, change)
					continue
				}
				product := parts[2]
				if s.recordExists(region, product) {
					err := delete(region, product)
					if err != nil {
						log.Printf("Error al borrar producto en la región %s: %v", region, err)
					}
				} else {
					log.Printf("El producto no existe en la región %s: %s", region, product)
				}

			default:
				log.Printf("Acción no reconocida en la región %s: %s", region, action)
			}
		}
	} else {
		log.Printf("No se encontraron cambios en logB para la región %s", region)
	}

	fmt.Print("Merge realizado con éxito\n")
	fmt.Printf("Numero de reloj: %v\n", s.regions[region])
	// Llamar a propagateChanges después de realizar el merge
	go s.propagateMerge(region)

	// Devolver respuesta
	return &pb.MergeChangesResponse{
		Region:      region, // O alguna lógica para elegir la región principal
		VectorClock: convertToInt32Slice(s.regions[region]),
	}, nil
}

func (s *server) applyChange(region string, remoteClock []int32, changes []string) {
	// Sincronizar el reloj vectorial
	for i := range s.regions[region] {
		if remoteClock[i] > int32(s.regions[region][i]) {
			s.regions[region][i] = int(remoteClock[i])
		}
	}
	// Ajustar el índice local
	s.regions[region][numServer]--

	// Procesar cada cambio recibido
	for _, change := range changes {
		parts := strings.Fields(change)
		if len(parts) < 2 {
			log.Printf("Cambio inválido: %s", change)
			continue
		}

		action := parts[0]
		switch action {
		case "AgregarProducto":
			if len(parts) < 3 {
				log.Printf("Cambio inválido para 'AgregarProducto': %s", change)
				continue
			}
			record := strings.Join(parts[2:], " ")
			err := add(region, record)
			if err != nil {
				log.Printf("Error al agregar registro: %v", err)
			}

		case "BorrarProducto":
			if len(parts) < 3 {
				log.Printf("Cambio inválido para 'BorrarProducto': %s", change)
				continue
			}
			product := parts[2]
			err := delete(region, product)
			if err != nil {
				log.Printf("Error al borrar producto: %v", err)
			}

		case "RenombrarProducto":
			if len(parts) < 4 {
				log.Printf("Cambio inválido para 'RenombrarProducto': %s", change)
				continue
			}
			oldProduct := parts[2]
			newProduct := parts[3]
			err := rename(region, oldProduct, newProduct)
			if err != nil {
				log.Printf("Error al renombrar producto: %v", err)
			}

		case "ActualizarValor":
			if len(parts) < 4 {
				log.Printf("Cambio inválido para 'ActualizarValor': %s", change)
				continue
			}
			product := parts[2]
			newValue := parts[3]
			err := updateValue(region, product, newValue)
			if err != nil {
				log.Printf("Error al actualizar valor: %v", err)
			}

		default:
			log.Printf("Acción no reconocida: %s", action)
		}
	}
}

// Función para obtener o inicializar un reloj vectorial para una región
func (s *server) getOrCreateRegion(region string) []int {
	fmt.Println("Obteniendo o creando reloj vectorial y archivo para la región", region)
	if _, exists := s.regions[region]; !exists {
		// Inicializa un reloj vectorial [0, 0, 0] para la nueva región
		s.regions[region] = []int{0, 0, 0}
		// Crea un archivo para la región
		file, err := os.Create(fmt.Sprintf("%s.txt", region))
		if err != nil {
			log.Printf("Error al crear el archivo para la región %s: %v", region, err)
		}
		defer file.Close()
	}

	return s.regions[region]
}

// Función que se encarga de actualizar el reloj vectorial y registrar la operación en el log
func (s *server) AddRecord(ctx context.Context, req *pb2.AddRecordRequest) (*pb2.AddRecordResponse, error) {
	region := req.Region
	product := req.Product
	value := req.Value
	s.mu.Lock()
	defer s.mu.Unlock()

	// Obtiene el reloj vectorial asociado a la región
	reloj := s.getOrCreateRegion(region)

	// Actualiza el reloj vectorial
	reloj[numServer]++

	// Registra la operación en el log
	s.logs[region] = append(s.logs[region], fmt.Sprintf("AgregarProducto %s %s %s", region, product, value))
	log.Printf("Registro agregado: en la región %s", region)

	// Llama a la función add para escribir el registro en el archivo
	if err := add(region, fmt.Sprintf("%s %s", product, value)); err != nil {
		return nil, err
	}
	fmt.Printf("Numero de reloj: %v\n", reloj)
	// Devuelve el reloj vectorial actualizado
	return &pb2.AddRecordResponse{
		VectorClock: convertToInt32Slice(reloj),
	}, nil
}

// Función que se encarga de escribir el registro en el archivo correspondiente
func add(region, record string) error {
	// Abre el archivo en modo lectura para verificar si el producto ya existe
	file, err := os.Open(fmt.Sprintf("%s.txt", region)) // Solo lectura
	if err != nil {
		log.Printf("Error al abrir el archivo para la región %s: %v", region, err)
		return err
	}

	product := strings.Fields(record)[0]
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 && parts[0] == product {
			// Si el producto ya existe, no lo agregamos
			log.Printf("El registro ya existe en la región %s: %s", region, record)
			return nil
		}
	}
	file.Close()

	// Si no existe, abrimos el archivo nuevamente en modo escritura
	file, err = os.OpenFile(fmt.Sprintf("%s.txt", region), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error al abrir el archivo para la región %s: %v", region, err)
		return err
	}
	defer file.Close()

	// Escribir el nuevo registro al final del archivo
	recordWithNewLine := record + "\n"
	if _, err := file.WriteString(recordWithNewLine); err != nil {
		log.Printf("Error al escribir en el archivo para la región %s: %v", region, err)
		return err
	}

	fmt.Printf("Producto agregado: %s en la región %s\n", record, region)
	return nil
}

func (s *server) DeleteRecord(ctx context.Context, req *pb2.DeleteRecordRequest) (*pb2.DeleteRecordResponse, error) {
	region := req.Region
	product := req.Product
	s.mu.Lock()
	defer s.mu.Unlock()

	// Obtiene el reloj vectorial asociado a la región
	reloj := s.getOrCreateRegion(region)

	// Actualiza el reloj vectorial
	reloj[numServer]++

	log.Printf("Borrando producto: %s en la región %s", product, region)
	// Llama a la función 'delete' para borrar el registro
	err := delete(region, product)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Numero de reloj: %v\n", reloj)
	// Devuelve el reloj vectorial actualizado
	return &pb2.DeleteRecordResponse{
		VectorClock: convertToInt32Slice(reloj),
	}, nil
}

func delete(region string, product string) error {
	// Lee el archivo para la región
	file, err := os.OpenFile(fmt.Sprintf("%s.txt", region), os.O_RDWR, 0644)
	if err != nil {
		log.Printf("Error al abrir el archivo para la región %s: %v", region, err)
		return err
	}
	defer file.Close()

	// Lee el contenido y filtra el producto a borrar
	lines, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error al leer el archivo para la región %s: %v", region, err)
		return err
	}

	// Filtra las líneas que no contengan el producto a borrar
	var updatedLines []string
	for _, line := range strings.Split(string(lines), "\n") {
		if !strings.Contains(line, product) {
			updatedLines = append(updatedLines, line)
		}
	}

	// Vuelve a escribir el archivo con las líneas actualizadas
	file.Seek(0, 0)
	file.Truncate(0)
	for _, line := range updatedLines {
		file.WriteString(line + "\n")
	}

	// Registra la operación en el log
	log.Printf("Producto borrado: %s en la región %s", product, region)
	return nil
}

func (s *server) RenameRecord(ctx context.Context, req *pb2.RenameRecordRequest) (*pb2.RenameRecordResponse, error) {
	region := req.Region
	oldProduct := req.OldProduct
	newProduct := req.NewProduct
	s.mu.Lock()
	defer s.mu.Unlock()

	// Obtiene el reloj vectorial asociado a la región
	reloj := s.getOrCreateRegion(region)

	// Actualiza el reloj vectorial
	reloj[numServer]++

	fmt.Printf("Renombrando producto: %s a %s en la región %s\n", oldProduct, newProduct, region)
	// Llama a la función 'rename' para renombrar el producto
	err := rename(region, oldProduct, newProduct)
	if err != nil {
		return nil, err
	}
	fmt.Printf("Numero de reloj: %v\n", reloj)
	// Devuelve el reloj vectorial actualizado
	return &pb2.RenameRecordResponse{
		VectorClock: convertToInt32Slice(reloj),
	}, nil
}

func rename(region string, oldProduct string, newProduct string) error {
	// Lee el archivo para la región
	file, err := os.OpenFile(fmt.Sprintf("%s.txt", region), os.O_RDWR, 0644)
	if err != nil {
		log.Printf("Error al abrir el archivo para la región %s: %v", region, err)
		return err
	}
	defer file.Close()

	// Lee el contenido y renombra el producto
	lines, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error al leer el archivo para la región %s: %v", region, err)
		return err
	}

	var updatedLines []string
	for _, line := range strings.Split(string(lines), "\n") {
		if strings.Contains(line, oldProduct) {
			line = strings.Replace(line, oldProduct, newProduct, 1)
		}
		updatedLines = append(updatedLines, line)
	}

	// Vuelve a escribir el archivo con las líneas actualizadas
	file.Seek(0, 0)
	file.Truncate(0)
	for _, line := range updatedLines {
		file.WriteString(line + "\n")
	}

	// Registra la operación en el log
	log.Printf("Producto renombrado de %s a %s en la región %s", oldProduct, newProduct, region)

	return nil
}

func (s *server) UpdateRecordValue(ctx context.Context, req *pb2.UpdateRecordRequest) (*pb2.UpdateRecordResponse, error) {
	region := req.Region
	product := req.Product
	newValue := req.NewValue
	s.mu.Lock()
	defer s.mu.Unlock()

	// Obtiene el reloj vectorial asociado a la región
	reloj := s.getOrCreateRegion(region)

	// Actualiza el reloj vectorial
	reloj[numServer]++
	fmt.Printf("Actualizando valor del producto: %s a %s en la región %s\n", product, newValue, region)
	// Llama a la función 'updateValue' para actualizar el valor del producto
	err := updateValue(region, product, fmt.Sprintf("%f", newValue))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Numero de reloj: %v\n", reloj)
	// Devuelve el reloj vectorial actualizado
	return &pb2.UpdateRecordResponse{
		VectorClock: convertToInt32Slice(reloj),
	}, nil
}

func updateValue(region string, product string, newValue string) error {
	// Lee el archivo para la región
	file, err := os.OpenFile(fmt.Sprintf("%s.txt", region), os.O_RDWR, 0644)
	if err != nil {
		log.Printf("Error al abrir el archivo para la región %s: %v", region, err)
		return err
	}
	defer file.Close()

	// Lee el contenido y actualiza el valor del producto
	lines, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error al leer el archivo para la región %s: %v", region, err)
		return err
	}

	var updatedLines []string
	for _, line := range strings.Split(string(lines), "\n") {
		if strings.Contains(line, product) {
			// Asume que el valor está al final de la línea, por ejemplo: "Vino 25"
			parts := strings.Fields(line)
			if len(parts) > 1 {
				parts[len(parts)-1] = newValue // Actualiza el valor
				line = strings.Join(parts, " ")
			}
		}
		updatedLines = append(updatedLines, line)
	}

	// Vuelve a escribir el archivo con las líneas actualizadas
	file.Seek(0, 0)
	file.Truncate(0)
	for _, line := range updatedLines {
		file.WriteString(line + "\n")
	}

	// Registra la operación en el log
	log.Printf("Valor del producto %s actualizado a %s en la región %s", product, newValue, region)

	return nil
}

// Función para convertir []int a []int32
func convertToIntSlice(int32Slice []int32) []int {
	intSlice := make([]int, len(int32Slice))
	for i, v := range int32Slice {
		intSlice[i] = int(v)
	}
	return intSlice
}

// Función para convertir []int a []int32
func convertToInt32Slice(intSlice []int) []int32 {
	int32Slice := make([]int32, len(intSlice))
	for i, v := range intSlice {
		int32Slice[i] = int32(v)
	}
	return int32Slice
}

// Función principal para configurar y ejecutar el servidor
func main() {
	lis, err := net.Listen("tcp", ":50053") // Cambia el puerto si es necesario
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}

	s := grpc.NewServer()
	hexServer := newServer()

	// Registro del servidor gRPC
	pb.RegisterHextechServer(s, hexServer) // Servicio para sincronización entre servidores
	pb2.RegisterHexSServer(s, hexServer)

	// Ejecuta la propagación de cambios en una rutina separada
	go hexServer.propagateChanges()

	log.Println("Servidor Hextech corriendo en el puerto :50053")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error al servir: %v", err)
	}
}
