package dto

type Pagination struct {
	TotalData int32 `json:"total_data"`
	TotalPage int32 `json:"total_page"`
	Page      int32 `json:"page"`
	Limit     int32 `json:"limit"`
	Offset    int32 `json:"-"`
}

type PaginationRequest struct {
	Page  int32 `json:"page"`
	Limit int32 `json:"limit"`
}
