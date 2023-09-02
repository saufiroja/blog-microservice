package main

import (
	"github.com/saufiroja/blog-microservice/user-service/config"
	"github.com/saufiroja/blog-microservice/user-service/infrastructures/grpc/server"
)

func main() {
	port := config.NewAppConfig().Grpc.Port
	host := config.NewAppConfig().Grpc.Host
	grpcServer := server.NewGrpcServer(host, port)
	grpcServer.Start()
}
