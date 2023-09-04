package routes

import (
	"go-boilerplate/pkg/gate"
	"net/http"
)

type SingleRoute struct {
	Name       string
	Method     string
	Path       string
	Validation interface{}
	Handle     http.HandlerFunc
	Middleware []gate.Middleware
}

type GroupRoute struct {
	Name       string
	PreFix     string
	Middleware []gate.Middleware
	Children   []SingleRoute
}

type Register struct {
	SingleRoutes []SingleRoute
	GroupRoutes  []GroupRoute
}
