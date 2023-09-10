package handler

import (
	"context"

	"github.com/google/uuid"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/client"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/auth"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/user"
	"github.com/saufiroja/blog-microservice/auth-service/models/dto"
	"github.com/saufiroja/blog-microservice/auth-service/utils"
)

type AuthHandler struct {
	auth.UnimplementedAuthServiceServer
	UserClient *client.UserServerClient
	bcrpyt     *utils.Password
	token      utils.GenerateToken
}

func NewAuthHandler(userClient *client.UserServerClient, bcrpyt *utils.Password, token utils.GenerateToken) *AuthHandler {
	return &AuthHandler{
		UserClient: userClient,
		bcrpyt:     bcrpyt,
		token:      token,
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
	// hash
	hashPassword, err := h.bcrpyt.Hash(req.Password)
	if err != nil {
		return nil, err
	}

	input.Password = hashPassword
	_, err = h.UserClient.InsertUser(input)
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
	findUserByEmailReq := &user.FindUsersByEmailRequest{
		Email: req.Email,
	}

	findUserByEmailRes, err := h.UserClient.FindUsersByEmail(findUserByEmailReq)
	if err != nil {
		return nil, err
	}

	// compare password
	err = h.bcrpyt.Compare(findUserByEmailRes.Result.Password, req.Password)
	if err != nil {
		return nil, err
	}

	// generate token
	input := &dto.GenerateTokenDTO{
		Id:    findUserByEmailRes.Result.Id,
		Name:  findUserByEmailRes.Result.Name,
		Email: findUserByEmailRes.Result.Email,
	}

	accessToken, expAccessToken, err := h.token.GenerateAccessToken(input)
	if err != nil {
		return nil, err
	}

	refreshToken, expRefreshToken, err := h.token.GenerateRefreshToken(input)
	if err != nil {
		return nil, err
	}

	res := &auth.LoginResponse{
		Code:    200,
		Message: "success login",
		Result: &auth.Token{
			AccessToken:         accessToken,
			RefreshToken:        refreshToken,
			ExpiredAccessToken:  int32(expAccessToken),
			ExpiredRefreshToken: int32(expRefreshToken),
		},
	}

	return res, nil
}
