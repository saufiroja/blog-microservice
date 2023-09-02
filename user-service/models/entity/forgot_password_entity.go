package entity

type ForgetPassword struct {
	ID         string `json:"id" gorm:"primaryKey"`
	OTP        string `json:"otp"`
	ExpiredOtp int64  `json:"expired_otp"`
	UserID     string `json:"user_id"`
	User       User   `json:"user"`
	CreatedAt  int64  `json:"created_at"`
	UpdatedAt  int64  `json:"updated_at"`
}
