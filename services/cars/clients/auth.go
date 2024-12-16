package clients

import (
	"github/cars/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var AuthClient auth.AuthServiceClient

func AuthClientInit() func() {
	conn, err := grpc.NewClient("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatal(err)
	}

	AuthClient = auth.NewAuthServiceClient(conn)

	return func() {
		conn.Close()
	}
}
