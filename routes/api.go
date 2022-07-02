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
				Middleware: []gate.Middleware{
					//middleware.Logging(),
				},
			},
		},

		GroupRoutes: []GroupRoute{
			{
				Name:   "User",
				PreFix: "/user",
				Middleware: []gate.Middleware{
					middleware.Global(),
				},
				Children: []SingleRoute{
					{
						Name:   "user",
						Method: "GET",
						Path:   "/auth-data",
						Handle: controllers.AuthData,
						Middleware: []gate.Middleware{
							middleware.Auth(),
						},
					},
				},
			},
		},
	}
}
