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
				PreFix: "/auth",
				Middleware: []gate.Middleware{
					middleware.Global(),
				},
				Children: []SingleRoute{
					{
						Name:       "user",
						Method:     "POST",
						Path:       "/user",
						Validation: &requests.LoginRequest{},
						Handle:     controllers.User,
						Middleware: []gate.Middleware{
							middleware.Logging(),
						},
					},
				},
			},
		},
	}
}
