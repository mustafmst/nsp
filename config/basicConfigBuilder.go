package config

import (
	"net/http"

	"github.com/mustafmst/nsp/controllers"

	"github.com/mustafmst/nsp/core"
)

// BasicBuilder - builds app config
type BasicBuilder struct {
	useMiddleware bool
	middleware    []func(w http.ResponseWriter, r *http.Request)
}

// ConfigRoutes - configure all routes for app
func (b *BasicBuilder) ConfigRoutes(l core.Logger) {
	c := controllers.NewControllersMap()

	http.HandleFunc(
		b.getHandler("/", c.GetControllerMethod("home", "index"), l))
	http.HandleFunc(
		b.getHandler("/info", c.GetControllerMethod("home", "info"), l))
}

// UseMiddleware - enables using middleware
func (b *BasicBuilder) UseMiddleware() {
	b.useMiddleware = true
	b.middleware = make([]func(w http.ResponseWriter, r *http.Request), 0)
}

// AddMiddlewareFunc - Adds function to be rub with middleware
func (b *BasicBuilder) AddMiddlewareFunc(f func(w http.ResponseWriter, r *http.Request)) {
	b.middleware = append(b.middleware, f)
}

// NewBasicBuilder - creates new builder instance
func NewBasicBuilder() *BasicBuilder {
	return &BasicBuilder{false, nil}
}

func (b *BasicBuilder) getHandler(route string, handler func(w http.ResponseWriter, r *http.Request), l core.Logger) (string, func(w http.ResponseWriter, r *http.Request)) {
	// l.LogInfo("route -> " + route)
	if b.useMiddleware == false {
		return route, handler
	}
	return route, func(w http.ResponseWriter, r *http.Request) {
		l.LogDebug("aaaaaaaaaaaaaaaaaaaaaaaa")
		handler(w, r)
	}
}
