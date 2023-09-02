package utils

import (
	"github.com/saufiroja/blog-microservice/user-service/infrastructures/grpc/rpc/pb/user"
	"github.com/saufiroja/blog-microservice/user-service/models/dto"
)

func SetPagination(req *user.PaginationRequest) *dto.Pagination {
	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	offset := (req.Page - 1) * req.Limit

	return &dto.Pagination{
		Page:   req.Page,
		Limit:  req.Limit,
		Offset: offset,
	}
}
