package main

import (
	"fmt"
	"github/order/clients"
	"github/order/db"
	order "github/order/gen"
	"github/order/server"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	db.Init()
	closeAuthClient := clients.AuthClientInit()
	defer closeAuthClient()

	closeCarsClient := clients.CarsClientInit()
	defer closeCarsClient()

	serverHost := os.Getenv("SERVER_HOST")
	lis, err := net.Listen("tcp", fmt.Sprintf(serverHost+":%d", 3002))
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	order.RegisterOrderServiceServer(grpcServer, &server.OrderServiceServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Printf("failed to serve: %v", err)
	}
}
