package requests

type LoginRequest struct {
	Email    string `json:"email" valid:"required~Email is required"`
	Password string `json:"password" valid:"required~Password is required"`
}
