package nsp

import (
	"fmt"
	"net/http"
)

type Router interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	HandleFunc(string, func(http.ResponseWriter, *http.Request)) Path
}

type basicRouter struct {
	paths []Path
}

func NewRouter() Router {
	return &basicRouter{make([]Path, 0)}
}

func (rt *basicRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
}

func (rt *basicRouter) HandleFunc(path string, f func(http.ResponseWriter, *http.Request)) Path {
	p := NewPath(path, f)
	rt.paths = append(rt.paths, p)
	return p
}
