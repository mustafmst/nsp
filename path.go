package nsp

import (
	"net/http"
)

type Path interface {
	GetPath() string
	GetHandler() http.Handler
	Method(string)
}

type path struct {
	path    string
	handler http.Handler
	method  string
}

func (p *path) GetPath() string {
	return p.path
}
func (p *path) GetHandler() http.Handler {
	return p.handler
}
func (p *path) Method(method string) {
	if method == "GET" || method == "POST" ||
		method == "PUT" || method == "DELETE" {
		p.method = method
	}
}

func NewPath(pathUrl string, handlerFunction func(http.ResponseWriter, *http.Request)) Path {
	return &path{pathUrl, handlerFunc{handlerFunction}, "GET"}
}

type handlerFunc struct {
	function func(http.ResponseWriter, *http.Request)
}

func (h handlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.function(w, r)
}
