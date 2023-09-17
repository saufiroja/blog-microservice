package interfaces

import (
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/auth"
	"github.com/saufiroja/blog-microservice/auth-service/infrastructures/grpc/rpc/pb/user"
)

type AuthService interface {
	Register(input *user.InsertUserDTO) error
	Login(input *user.FindUsersByEmailRequest, password string) (*auth.Token, error)
}
