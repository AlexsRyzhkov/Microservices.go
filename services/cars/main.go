package main

import (
	"fmt"
	"github/cars/clients"
	"github/cars/db"
	cars "github/cars/gen"
	"github/cars/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db.Init()
	closeClient := clients.AuthClientInit()
	defer closeClient()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3001))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	cars.RegisterCarsServiceServer(grpcServer, &server.CarsServiceServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
