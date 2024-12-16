package main

import (
	"auth/db"
	auth "auth/gen"
	"auth/server"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db.Init()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 3000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authServer := &server.AuthServiceServer{}

	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authServer)

	log.Printf("server listening at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
