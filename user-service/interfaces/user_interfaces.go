package interfaces

import "github.com/saufiroja/blog-microservice/user-service/models/dto"

type UserRepository interface {
	FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, error)
	CountAllUsers() (int32, error)
	InsertUser(user *dto.InsertUserDTO) error
	FindUsersByEmail(email string) (*dto.FindUsersDTO, error)
	FindUsersByID(id string) (*dto.FindUsersDTO, error)
	UpdateUser(id string, user *dto.UpdateUserDTO) error
	DeleteUser(id string) error
}

type UserService interface {
	FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, *dto.Pagination, error)
	InsertUser(user *dto.InsertUserDTO) error
	FindUsersByEmail(email string) (*dto.FindUsersDTO, error)
	FindUsersByID(id string) (*dto.FindUsersDTO, error)
	UpdateUser(id string, user *dto.UpdateUserDTO) error
	DeleteUser(id string) error
}
