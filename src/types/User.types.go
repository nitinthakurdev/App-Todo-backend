package types

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Message  string `json:"message"`
	Token    string `json:"token"`
}
