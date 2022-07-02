package Services

import (
	"go-boilerplate/models"
	"go-boilerplate/requests"
	"golang.org/x/crypto/bcrypt"
)

// Token jwt Standard Claim Object
/*type Token struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.StandardClaims
}*/

// Create a dummy local db instance as a key value pair
var userdb = map[string]string{
	"imrancse94@gmail.com": "Nop@ss123411",
}

// Login user login function
func Login(input requests.LoginRequest) (data interface{}, error string) {

	user := models.GetUserByEmail(input.Email)

	if !CheckPasswordHash(input.Password, user.Password) {
		return user, "Invalid email or password"
	}

	jwt := Jwt{}
	token, err := jwt.CreateToken(user)
	if err != nil {
		return nil, "Failed to create token"
	}

	return token, "Successfully created token"
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
