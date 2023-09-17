package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/auth"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/user"
	"github.com/saufiroja/blog-microservice/auth-service/interfaces"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	authService interfaces.AuthService
}

func NewAuthHandler(authService interfaces.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
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

	err := h.authService.Register(input)
	if err != nil {
		return nil, err
	}

	res := &auth.RegisterResponse{
		Code:    201,
		Message: "success register user",
	}

	return res, nil
}

func (h *AuthHandler) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	// find user by email
	input := &user.FindUsersByEmailRequest{
		Email: req.Email,
	}

	res, err := h.authService.Login(input, req.Password)
	if err != nil {
		return nil, err
	}

	// set response
	result := &auth.LoginResponse{
		Code:    200,
		Message: "success login",
		Result:  res,
	}

	return result, nil
}
