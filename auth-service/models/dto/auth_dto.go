package dto

type RegisterDTO struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int32  `json:"created_at"`
	UpdatedAt int32  `json:"updated_at"`
}

type RegisterResponse struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	ExpiredAccessToken  int32  `json:"expired_access_token"`
	ExpiredRefreshToken int32  `json:"expired_refresh_token"`
}
