package dto

type CreateUserRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type UserLoginResponse struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}