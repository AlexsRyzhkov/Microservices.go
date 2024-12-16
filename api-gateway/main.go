package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github/api-gateway/gen/auth"
	"github/api-gateway/gen/cars"
	"github/api-gateway/gen/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"log"
	"net/http"
)

var (
	grpcAuthServerEndpoint  = "localhost:3000"
	grpcCarsServerEndpoint  = "localhost:3001"
	grpcOrderServerEndpoint = "localhost:3002"
)

func run() error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, grpcAuthServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = cars.RegisterCarsServiceHandlerFromEndpoint(ctx, mux, grpcCarsServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, grpcOrderServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	return http.ListenAndServe(":8000", mux)
}

func main() {
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
