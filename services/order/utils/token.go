package utils

import (
	"context"
	"github/order/clients"
	"github/order/gen/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"log"
)

func VerifyToken(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "No headers")
	}

	var token string
	if values := md.Get("Barear"); len(values) > 0 {
		token = values[0]
		log.Printf("Authorization token: %s", token)
	} else {
		return status.Error(codes.Unauthenticated, "No authorizatio header")
	}

	_, err := clients.AuthClient.ValidateToken(context.Background(), &auth.TokenRequest{
		Token: token,
	})

	if err != nil {
		return status.Error(codes.Unauthenticated, "No authorization")
	}

	return nil
}
