package requests

type UserRequest struct {
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password"`
	Role     string `json:"role"`
	// Email    string `json:"email" binding:"required,email,unique"`
}
