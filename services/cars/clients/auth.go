package clients

import (
	"github/cars/gen/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

var AuthClient auth.AuthServiceClient

func AuthClientInit() func() {

	authClientHost := os.Getenv("AUTH_CLIENT_HOST")
	conn, err := grpc.NewClient(authClientHost+":3000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("clientAuthError: " + err.Error())
	}

	AuthClient = auth.NewAuthServiceClient(conn)

	return func() {
		conn.Close()
	}
}
