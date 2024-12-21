package clients

import (
	"github/order/gen/cars"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var CarsClient cars.CarsServiceClient

func CarsClientInit() func() {
	carsClientHost := os.Getenv("CARS_CLIENT_HOST")
	conn, err := grpc.NewClient(carsClientHost+":3001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("clientAuthError: " + err.Error())
	}

	CarsClient = cars.NewCarsServiceClient(conn)

	return func() {
		conn.Close()
	}
}
