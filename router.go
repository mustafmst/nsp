package nsp

import (
	"fmt"
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
	fmt.Println(r)
}

func (rt *basicRouter) AddPath(path string) PathNode {
	return rt.paths.AddPath(path)
}
