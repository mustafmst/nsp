package sample

import (
	"net/http"

	"github.com/mustafmst/nsp"
	"github.com/mustafmst/nsp/controllers"
)

// basicBuilder - builds app config
type basicBuilder struct {
	useMiddleware  bool
	middleware     []func(w http.ResponseWriter, r *http.Request)
	controllersMap controllers.ControllersMap
	configFunc     func(controllers.ControllersMap, nsp.Logger)
}

// ConfigRoutes - configure all routes for app
func (b *basicBuilder) ConfigRoutes(l nsp.Logger) {
	http.HandleFunc(
		b.getHandler("/", b.controllersMap.GetControllerMethod("home", "index"), l))
	http.HandleFunc(
		b.getHandler("/info", b.controllersMap.GetControllerMethod("home", "info"), l))
}

// UseMiddleware - enables using middleware
func (b *basicBuilder) UseMiddleware() {
	b.useMiddleware = true
	b.middleware = make([]func(w http.ResponseWriter, r *http.Request), 0)
}

// AddMiddlewareFunc - Adds function to be rub with middleware
func (b *basicBuilder) AddMiddlewareFunc(f func(w http.ResponseWriter, r *http.Request)) {
	b.middleware = append(b.middleware, f)
}

// NewBasicBuilder - creates new builder instance
func NewBasicBuilder(controllersMap controllers.ControllersMap) *basicBuilder {
	return &basicBuilder{false, nil, controllersMap, nil}
}

func (b *basicBuilder) getHandler(route string, handler func(w http.ResponseWriter, r *http.Request), l nsp.Logger) (string, func(w http.ResponseWriter, r *http.Request)) {
	// l.LogInfo("route -> " + route)
	if b.useMiddleware == false {
		return route, handler
	}
	return route, func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
