package dto

type RegisterRequest struct {
	Name            string            `json:"name" binding:"required"`
	Email           string            `json:"email" binding:"required,email"`
	Password        string            `json:"password" binding:"required,min=8"`
	PasswordConfirm string            `json:"password_confirm" binding:"required,min=8"`
	Gender          string `json:"gender" binding:"required,oneof=pria wanita"`
}
