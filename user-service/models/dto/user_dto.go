package dto

type FindAllUsersDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	CreatedAt int32  `json:"created_at"`
}
