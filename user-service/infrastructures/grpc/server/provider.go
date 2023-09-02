package server

import (
	"github.com/saufiroja/blog-microservice/user-service/config"
	"github.com/saufiroja/blog-microservice/user-service/infrastructures/database"
	"github.com/saufiroja/blog-microservice/user-service/infrastructures/grpc/handler"
	"github.com/saufiroja/blog-microservice/user-service/repositories"
	"github.com/saufiroja/blog-microservice/user-service/services"
)

type gRPCProvider struct {
	handlers struct {
		user handler.UserHandler
	}
}

func (rpc *GrpcServer) provide() gRPCProvider {
	provider := gRPCProvider{}

	conf := config.NewAppConfig()
	db := database.NewPostgres(conf)

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	provider.handlers.user = *handler.NewUserHandler(userService)

	return provider
}
