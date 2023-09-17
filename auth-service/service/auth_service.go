package service

import (
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/client"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/auth"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/user"
	"github.com/saufiroja/blog-microservice/auth-service/interfaces"
	"github.com/saufiroja/blog-microservice/auth-service/models/dto"
	"github.com/saufiroja/blog-microservice/auth-service/utils"
)

type AuthService struct {
	UserClient *client.UserServerClient
	bcrpyt     *utils.Password
	token      utils.GenerateToken
}

func NewAuthService(userClient *client.UserServerClient, bcrpyt *utils.Password, token utils.GenerateToken) interfaces.AuthService {
	return &AuthService{
		UserClient: userClient,
		bcrpyt:     bcrpyt,
		token:      token,
	}
}

func (s *AuthService) Register(input *user.InsertUserDTO) error {
	hashPassword, err := s.bcrpyt.Hash(input.Password)
	if err != nil {
		return err
	}

	input.Password = hashPassword
	_, err = s.UserClient.InsertUser(input)
	if err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Login(input *user.FindUsersByEmailRequest, password string) (*auth.Token, error) {
	findUserByEmailRes, err := s.UserClient.FindUsersByEmail(input)
	if err != nil {
		return nil, err
	}

	// compare password
	err = s.bcrpyt.Compare(findUserByEmailRes.Result.Password, password)
	if err != nil {
		return nil, err
	}

	// generate token
	payload := &dto.GenerateTokenDTO{
		Id:    findUserByEmailRes.Result.Id,
		Name:  findUserByEmailRes.Result.Name,
		Email: findUserByEmailRes.Result.Email,
	}

	accessToken, expAccessToken, err := s.token.GenerateAccessToken(payload)
	if err != nil {
		return nil, err
	}

	refreshToken, expRefreshToken, err := s.token.GenerateRefreshToken(payload)
	if err != nil {
		return nil, err
	}

	res := &auth.Token{
		AccessToken:         accessToken,
		RefreshToken:        refreshToken,
		ExpiredAccessToken:  int32(expAccessToken),
		ExpiredRefreshToken: int32(expRefreshToken),
	}

	return res, nil
}
