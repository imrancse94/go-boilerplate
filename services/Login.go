package Services

import (
	"go-boilerplate/models"
	"go-boilerplate/requests"
)

// Login user login function
func Login(input requests.LoginRequest) (data interface{}, error string) {

	user := models.GetUserByEmail(input.Email)

	if !CheckPasswordHash(input.Password, user.Password) {
		return user, "Invalid email or password"
	}

	return CreateTokenDataByUser(user)
}

func CreateTokenDataByUser(user models.User) (data interface{}, error string) {
	jwt := Jwt{}
	token, err := jwt.CreateToken(user)
	if err != nil {
		return nil, "Failed to create token"
	}

	return token, "Successfully created token"
}
