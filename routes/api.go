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
				Name:   "Test",
				Method: "GET",
				Path:   "/test",
				Handle: controllers.Test,
			},
			{
				Name:   "Test",
				Method: "POST",
				Path:   "/ed",
				Handle: controllers.TestEncryptDecrypt,
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
					{
						Name:       "roles",
						Method:     "GET",
						Path:       "/roles",
						Handle:     controllers.GetRoles,
						Middleware: []gate.Middleware{},
					},
					{
						Name:       "add-roles",
						Method:     "POST",
						Path:       "/add/roles",
						Handle:     controllers.AddRole,
						Validation: &requests.AddRoleRequest{},
						Middleware: []gate.Middleware{},
					},
				},
			},
		},
	}
}
