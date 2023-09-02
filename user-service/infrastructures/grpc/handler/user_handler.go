package handler

import (
	"context"

	"github.com/saufiroja/blog-microservice/user-service/infrastructures/grpc/rpc/pb/user"
	"github.com/saufiroja/blog-microservice/user-service/interfaces"
	"github.com/saufiroja/blog-microservice/user-service/utils"
)

type UserHandler struct {
	user.UnimplementedUserServiceServer
	userService interfaces.UserService
}

func NewUserHandler(userService interfaces.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) FindAllUsers(ctx context.Context, req *user.PaginationRequest) (*user.FindAllUsersResponse, error) {
	// set pagination
	setPage := utils.SetPagination(req)

	// call service
	users, page, err := h.userService.FindAllUsers(setPage)
	if err != nil {
		return nil, err
	}

	// convert pagination
	paginationRes := &user.Pagination{
		Limit:     page.Limit,
		Page:      page.Page,
		TotalData: page.TotalData,
		TotalPage: page.TotalPage,
	}

	// convert response
	var usersRes []*user.FindAllUsersDTO
	for _, v := range users {
		userRes := &user.FindAllUsersDTO{
			Id:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			CreatedAt: v.CreatedAt,
		}
		usersRes = append(usersRes, userRes)
	}

	// send response
	res := user.FindAllUsersResponse{
		Code:       200,
		Message:    "success find all users",
		Pagination: paginationRes,
		Result:     usersRes,
	}

	return &res, nil
}
