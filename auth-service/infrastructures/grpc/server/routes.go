package server

import (
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/auth"
)

func (rpc *GrpcServer) defineRoute(provider gRPCProvider) {
	auth.RegisterAuthServiceServer(rpc.grpcServer, &provider.handlers.auth)
}
