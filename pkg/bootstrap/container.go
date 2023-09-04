package bootstrap

import (
	"github.com/gorilla/mux"
	"go-boilerplate/middleware"
	"go-boilerplate/pkg/gate"
	"go-boilerplate/requests"
	"go-boilerplate/routes"
	"net/http"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
		return
	})
}

func Init() *mux.Router {

	Register := routes.Api()
	r := mux.NewRouter()
	r.Use(CORS)
	for _, route := range Register.SingleRoutes {
		if route.Middleware != nil {
			r.HandleFunc(route.Path, Initialize(route.Handle, middleware.Method(route.Method), route.Middleware, requests.Validation(route.Validation)))
		} else {
			r.HandleFunc(route.Path, Initialize(route.Handle, middleware.Method(route.Method), nil, requests.Validation(route.Validation)))
		}
	}

	for _, route := range Register.GroupRoutes {
		for _, child := range route.Children {
			if child.Middleware != nil {
				currentMiddleware := child.Middleware
				if route.Middleware != nil {
					currentMiddleware = append(currentMiddleware, route.Middleware...)
				}
				r.HandleFunc(route.PreFix+child.Path, Initialize(child.Handle, middleware.Method(child.Method), currentMiddleware, requests.Validation(child.Validation)))
			} else {
				r.HandleFunc(route.PreFix+child.Path, Initialize(child.Handle, middleware.Method(child.Method), nil, requests.Validation(child.Validation)))
			}
		}
	}

	http.Handle("/", r)
	return r
}

func Initialize(f http.HandlerFunc, MethodMiddleware gate.Middleware, middlewares []gate.Middleware, Validation gate.Middleware) http.HandlerFunc {

	for _, m := range middlewares {
		f = m(f)
	}
	f = Validation(f)
	f = MethodMiddleware(f)
	return f
}
