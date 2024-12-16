package clients

import (
	"github/order/gen/cars"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var CarsClient cars.CarsServiceClient

func CarsClientInit() func() {
	conn, err := grpc.NewClient("localhost:3001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	CarsClient = cars.NewCarsServiceClient(conn)

	return func() {
		conn.Close()
	}
}
