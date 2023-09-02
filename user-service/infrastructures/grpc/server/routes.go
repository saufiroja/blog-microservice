package server

import "github.com/saufiroja/blog-microservice/user-service/infrastructures/grpc/rpc/pb/user"

func (rpc *GrpcServer) defineRoute(provider gRPCProvider) {
	user.RegisterUserServiceServer(rpc.grpcServer, &provider.handlers.user)
}
