package middleware

import (
	"go-boilerplate/constant"
	"go-boilerplate/gate"
	"go-boilerplate/response"
	Services "go-boilerplate/services"
	"net/http"
	"strconv"
	"strings"
)

func Auth() gate.Middleware {

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

				signedToken := bearerToken[1]
				service := Services.Jwt{}
				user, err := service.ValidateToken(signedToken)
				//fmt.Println("user", user)
				if err != nil {
					unauthenticatedResponseMessage(w)
					return
					//fmt.Fprintf(w, err.Error())
				}

				r.Header.Set("auth_id", strconv.Itoa(user.ID))
				r.Header.Set("email", user.Email)
			}

			f(w, r)
		}
	}
}

func unauthenticatedResponseMessage(w http.ResponseWriter) {
	res := response.Response{
		StatusCode: constant.Status("UNAUTHORIZED"),
		Message:    "Unauthenticated",
		Data:       "",
	}
	response.SuccessRespond(res, w)
}
