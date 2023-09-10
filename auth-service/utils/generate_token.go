package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/saufiroja/blog-microservice/auth-service/models/dto"
)

type GenerateToken interface {
	GenerateAccessToken(input *dto.GenerateTokenDTO) (string, int64, error)
	GenerateRefreshToken(input *dto.GenerateTokenDTO) (string, int64, error)
}

type GenerateTokenImpl struct {
}

func NewGenerateToken() GenerateToken {
	return &GenerateTokenImpl{}
}

func (g *GenerateTokenImpl) GenerateAccessToken(input *dto.GenerateTokenDTO) (string, int64, error) {
	expired := time.Now().Add(time.Minute * 15).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    input.Id,
		"name":  input.Name,
		"email": input.Email,
		"exp":   expired,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expired, nil
}

func (g *GenerateTokenImpl) GenerateRefreshToken(input *dto.GenerateTokenDTO) (string, int64, error) {
	expired := time.Now().Add(time.Hour * 24 * 7).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    input.Id,
		"name":  input.Name,
		"email": input.Email,
		"exp":   expired,
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", 0, err
	}

	return tokenString, expired, nil
}
