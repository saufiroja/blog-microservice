package interfaces

import "github.com/saufiroja/blog-microservice/user-service/models/dto"

type UserRepository interface {
	FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, error)
	CountAllUsers() (int32, error)
}

type UserService interface {
	FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, *dto.Pagination, error)
}
