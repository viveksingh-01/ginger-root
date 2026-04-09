package auth

type SignupRequest struct {
	Phone    string `json:"phone" binding:"required,min=10"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Phone    string `json:"phone" binding:"required,min=10"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type AuthResponse struct {
	User  *UserResponse `json:"user"`
	Token string        `json:"token"`
}

type Response struct {
	StatusCode    int           `json:"statusCode"`
	StatusMessage string        `json:"statusMessage"`
	Data          *AuthResponse `json:"data,omitempty"`
}
