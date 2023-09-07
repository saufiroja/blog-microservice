package main

import (
	"github.com/saufiroja/blog-microservice/auth-service/config"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/server"
)

func main() {
	port := config.NewAppConfig().Grpc.Port
	host := config.NewAppConfig().Grpc.Host
	grpcServer := server.NewGrpcServer(host, port)
	grpcServer.Start()
}
