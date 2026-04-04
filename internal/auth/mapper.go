package auth

func ToUserResponse(user *User) *UserResponse {
	return &UserResponse{
		ID:    user.ID.Hex(),
		Email: user.Email,
		Name:  user.Name,
		Phone: user.Phone,
	}
}
