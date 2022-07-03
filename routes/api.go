package routes

import (
	"go-boilerplate/controllers"
	"go-boilerplate/gate"
	"go-boilerplate/middleware"
	"go-boilerplate/requests"
)

func Api() Register {

	return Register{
		SingleRoutes: []SingleRoute{
			{
				Name:       "Hello",
				Method:     "POST",
				Path:       "/login",
				Validation: &requests.LoginRequest{},
				Handle:     controllers.Login,
			},
			{
				Name:   "RefreshToken",
				Method: "POST",
				Path:   "/refresh-token",
				Handle: controllers.RefreshToken,
				Middleware: []gate.Middleware{
					middleware.ValidateRefreshToken(),
				},
			},
		},

		GroupRoutes: []GroupRoute{
			{
				Name:   "User",
				PreFix: "/user",
				Middleware: []gate.Middleware{
					middleware.Auth(),
				},
				Children: []SingleRoute{
					{
						Name:       "user",
						Method:     "GET",
						Path:       "/auth-data",
						Handle:     controllers.AuthData,
						Middleware: []gate.Middleware{},
					},
				},
			},
		},
	}
}
