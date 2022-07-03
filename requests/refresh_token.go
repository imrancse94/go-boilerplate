package requests

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" valid:"required~refresh_token is required"`
}
