package services

import (
	"math"

	"github.com/google/uuid"
	"github.com/saufiroja/blog-microservice/user-service/interfaces"
	"github.com/saufiroja/blog-microservice/user-service/models/dto"
)

type UserService struct {
	userRepo interfaces.UserRepository
}

func NewUserService(userRepo interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (s *UserService) FindAllUsers(pagination *dto.Pagination) ([]dto.FindAllUsersDTO, *dto.Pagination, error) {
	count, err := s.userRepo.CountAllUsers()
	if err != nil {
		return nil, nil, err
	}

	pagination.TotalData = count
	totalPage := math.Ceil(float64(count) / float64(pagination.Limit))
	pagination.TotalPage = int32(totalPage)

	users, err := s.userRepo.FindAllUsers(pagination)
	if err != nil {
		return nil, nil, err
	}

	return users, pagination, nil
}

func (s *UserService) InsertUser(user *dto.InsertUserDTO) error {
	user.ID = uuid.New().String()
	err := s.userRepo.InsertUser(user)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) FindUsersByEmail(email string) (*dto.FindUsersByEmailDTO, error) {
	user, err := s.userRepo.FindUsersByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}
