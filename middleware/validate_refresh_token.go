package middleware

import (
	"go-boilerplate/Helper"
	"go-boilerplate/models"
	"go-boilerplate/pkg/gate"
	"go-boilerplate/requests"
	Services "go-boilerplate/services"
	"net/http"
	"strconv"
	"strings"
)

var RefreshTokenRequest requests.RefreshTokenRequest

func ValidateRefreshToken() gate.Middleware {

	// Create a new gate
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			bearerToken := strings.Split(authHeader, " ")

			//token := bearerToken[1]
			if len(bearerToken) < 2 {
				unauthenticatedResponseMessage(w)
				return
			} else {

				Helper.Request(r, &RefreshTokenRequest)

				service := Services.Jwt{}
				user, err := service.ValidateRefreshToken(models.Token{
					AccessToken:  bearerToken[1],
					RefreshToken: RefreshTokenRequest.RefreshToken,
				})

				if err != nil {
					unauthenticatedResponseMessage(w)
					return
				}

				r.Header.Set("auth_id", strconv.Itoa(user.ID))
				r.Header.Set("email", user.Email)

			}
			f(w, r)
		}
	}
}
