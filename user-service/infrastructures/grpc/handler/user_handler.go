package handler

import (
	"context"

	pb "github.com/saufiroja/blog-microservice/user-service/infrastructures/grpc/rpc/pb/user"
	"github.com/saufiroja/blog-microservice/user-service/interfaces"
	"github.com/saufiroja/blog-microservice/user-service/models/dto"
	"github.com/saufiroja/blog-microservice/user-service/utils"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) FindAllUsers(ctx context.Context, req *pb.PaginationRequest) (*pb.FindAllUsersResponse, error) {
	// set pagination
	setPage := utils.SetPagination(req)

	// call service
	users, page, err := h.userService.FindAllUsers(setPage)
	if err != nil {
		return nil, err
	}

	// convert pagination
	paginationRes := &pb.Pagination{
		Limit:     page.Limit,
		Page:      page.Page,
		TotalData: page.TotalData,
		TotalPage: page.TotalPage,
	}

	// convert response
	var usersRes []*pb.FindAllUsersDTO
	for _, v := range users {
		userRes := &pb.FindAllUsersDTO{
			Id:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		}
		usersRes = append(usersRes, userRes)
	}

	// send response
	res := pb.FindAllUsersResponse{
		Code:       200,
		Message:    "success find all users",
		Pagination: paginationRes,
		Result:     usersRes,
	}

	return &res, nil
}

func (h *UserHandler) InsertUser(ctx context.Context, req *pb.InsertUserDTO) (*pb.InsertUserResponse, error) {
	// convert request
	userReq := &dto.InsertUserDTO{
		Name:      req.Name,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: req.CreatedAt,
		UpdatedAt: req.UpdatedAt,
	}

	// call service
	err := h.userService.InsertUser(userReq)
	if err != nil {
		return nil, err
	}

	// send response
	res := pb.InsertUserResponse{
		Code:    201,
		Message: "success insert user",
	}

	return &res, nil
}

func (h *UserHandler) FindUsersByEmail(ctx context.Context, req *pb.FindUsersByEmailRequest) (*pb.FindUsersByEmailResponse, error) {
	// call service
	user, err := h.userService.FindUsersByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	// convert response
	userRes := &pb.FindUsersByEmailDTO{
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
	}

	// send response
	res := pb.FindUsersByEmailResponse{
		Code:    200,
		Message: "success find user by email",
		Result:  userRes,
	}

	return &res, nil
}
