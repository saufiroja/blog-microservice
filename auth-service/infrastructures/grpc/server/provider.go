package server

import (
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/client"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/handler"
	"github.com/saufiroja/blog-microservice/auth-service/utils"
)

type gRPCProvider struct {
	handlers struct {
		auth handler.AuthHandler
	}
}

func (rpc *GrpcServer) provide() gRPCProvider {
	provider := gRPCProvider{}

	userClient := client.NewUserServerClient()
	bcrypt := utils.NewPassword()
	token := utils.NewGenerateToken()

	provider.handlers.auth = *handler.NewAuthHandler(&userClient, bcrypt, token)

	return provider
}
