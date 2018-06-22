package config

import (
	"encoding/json"
	"net/http"

	"github.com/mustafmst/nsp/core"
)

// BasicBuilder - builds app config
type BasicBuilder struct {
	useMiddleware bool
	middleware    []func(w http.ResponseWriter, r *http.Request)
}

// ConfigRoutes - configure all routes for app
func (b *BasicBuilder) ConfigRoutes(l core.Logger) {
	http.HandleFunc("/", b.getHandler(func(w http.ResponseWriter, r *http.Request) {
		l.LogInfo("Get -> /")
		json.NewEncoder(w).Encode("Hello world!")
	}))
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

func (b *BasicBuilder) getHandler(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	if b.useMiddleware == false {
		return handler
	}
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
