package client

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/saufiroja/blog-microservice/auth-service/config"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/user"
	"google.golang.org/grpc"
)

type UserServerClient struct {
	UserClient user.UserServiceClient
}

func NewUserServerClient() UserServerClient {
	conf := config.NewAppConfig()
	url := fmt.Sprintf("%s:%s", conf.UserService.Host, conf.UserService.Port)

	conn, err := grpc.Dial(url, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	log.Println("connected to user service")

	return UserServerClient{
		UserClient: user.NewUserServiceClient(conn),
	}
}

func (c *UserServerClient) InsertUser(req *user.InsertUserDTO) (*user.InsertUserResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.UserClient.InsertUser(ctx, req)
}

func (c *UserServerClient) FindUsersByEmail(req *user.FindUsersByEmailRequest) (*user.FindUsersByEmailResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return c.UserClient.FindUsersByEmail(ctx, req)
}
