package bootstrap

import (
	"go-boilerplate/gate"
	"go-boilerplate/middleware"
	"go-boilerplate/requests"
	"go-boilerplate/routes"
	"net/http"
)

func Init() {

	Register := routes.Api()

	for _, route := range Register.SingleRoutes {
		if route.Middleware != nil {
			http.HandleFunc(route.Path, Initialize(route.Handle, middleware.Method(route.Method), route.Middleware, requests.Validation(route.Validation)))
		} else {
			http.HandleFunc(route.Path, Initialize(route.Handle, middleware.Method(route.Method), nil, requests.Validation(route.Validation)))
		}
	}

	for _, route := range Register.GroupRoutes {
		for _, child := range route.Children {
			if child.Middleware != nil {
				currentMiddleware := child.Middleware
				if route.Middleware != nil {
					currentMiddleware = append(currentMiddleware, route.Middleware...)
				}
				http.HandleFunc(route.PreFix+child.Path, Initialize(child.Handle, middleware.Method(child.Method), currentMiddleware, requests.Validation(child.Validation)))
			} else {
				http.HandleFunc(route.PreFix+child.Path, Initialize(child.Handle, middleware.Method(child.Method), nil, requests.Validation(child.Validation)))
			}
		}
	}

}

func Initialize(f http.HandlerFunc, MethodMiddleware gate.Middleware, middlewares []gate.Middleware, Validation gate.Middleware) http.HandlerFunc {

	for _, m := range middlewares {
		f = m(f)
	}
	f = Validation(f)
	f = MethodMiddleware(f)
	return f
}
