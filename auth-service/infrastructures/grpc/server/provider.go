package server

import (
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/client"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/handler"
)

type gRPCProvider struct {
	handlers struct {
		auth handler.AuthHandler
	}
}

func (rpc *GrpcServer) provide() gRPCProvider {
	provider := gRPCProvider{}

	userClient := client.NewUserServerClient()

	provider.handlers.auth = *handler.NewAuthHandler(&userClient)

	return provider
}
