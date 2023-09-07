package handler

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/client"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/auth"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/user"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	UserClient *client.UserServerClient
}

func NewAuthHandler(userClient *client.UserServerClient) *AuthHandler {
	return &AuthHandler{
		UserClient: userClient,
	}
}

func (h *AuthHandler) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	input := &user.InsertUserDTO{
		Id:        uuid.New().String(),
		Email:     req.Email,
		Name:      req.Name,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}
	fmt.Println(input)
	_, err := h.UserClient.InsertUser(input)
	if err != nil {
		return nil, err
	}

	res := &auth.RegisterResponse{
		Code:    201,
		Message: "success register user",
	}

	return res, nil
}
