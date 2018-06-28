package nsp

import (
	"net/http"
)

// Router interface
type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	AddPath(string) PathNode
}

type basicRouter struct {
	paths PathNode
}

// NewRouter returns new empty router instance
func NewRouter() Router {
	return &basicRouter{NewPathNode()}
}

func (rt *basicRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rt.paths.ServeHTTP(w, r)
}

func (rt *basicRouter) AddPath(path string) PathNode {
	return rt.paths.AddPath(path)
}
