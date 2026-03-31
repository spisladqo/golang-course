package main

import (
	pb "collector/api/proto/v1"
	"collector/internal/drivers"
	"collector/internal/handlers"
	"collector/internal/usecases"
	"log"
	"net"
	"net/http"
	"time"

	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	client := &http.Client{Timeout: 30 * time.Second}
	driver := drivers.NewRepositoryDriver(client)
	usecases := usecases.NewRepositoryUsecases(driver)
	handler := handlers.NewRepositoryHandler(usecases)

	lis, err := net.Listen("tcp", address)
	if err != nil {
		return
	}
	s := grpc.NewServer()

	pb.RegisterRepositoryServiceServer(s, handler)

	log.Println("Serving gRPC on", address)
	log.Fatalln(s.Serve(lis))
}
