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
	"os"
)

var (
	grpcAuthServerEndpoint  = ":3000"
	grpcCarsServerEndpoint  = ":3001"
	grpcOrderServerEndpoint = ":3002"
)

func run() error {
	authHost := os.Getenv("AUTH_HOST")
	carsHost := os.Getenv("CARS_HOST")
	orderHost := os.Getenv("ORDER_HOST")

	apiServerHost := os.Getenv("API_SERVER_HOST")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := auth.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, authHost+grpcAuthServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = cars.RegisterCarsServiceHandlerFromEndpoint(ctx, mux, carsHost+grpcCarsServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	err = order.RegisterOrderServiceHandlerFromEndpoint(ctx, mux, orderHost+grpcOrderServerEndpoint, opts)
	if err != nil {
		log.Fatal(err)
	}

	return http.ListenAndServe(apiServerHost+":8000", mux)
}

func main() {
	if err := run(); err != nil {
		grpclog.Fatal(err)
	}
}
