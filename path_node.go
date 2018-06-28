package nsp

import (
	"net/http"
	"strings"
)

// PathNode holds tree representation of api paths
type PathNode interface {
	AddMethod(string, func(http.ResponseWriter, *http.Request)) PathNode
	addPath([]string) PathNode
	AddPath(string) PathNode
	ServeHTTP(http.ResponseWriter, *http.Request)
}
type pathNode struct {
	childNodes map[string]PathNode
	methods    map[string]http.Handler
}

func (p *pathNode) AddMethod(method string, handler func(http.ResponseWriter, *http.Request)) PathNode {
	if _, ok := p.methods[method]; !ok {
		correctMethod := false
		switch method {
		case "GET":
			fallthrough
		case "POST":
			fallthrough
		case "PUT":
			fallthrough
		case "DELETE":
			correctMethod = true
		}
		if correctMethod {
			p.methods[method] = &pathMethodHandler{handler}
		}
	}
	return p
}

func (p *pathNode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: get proper node for request path
	if m, ok := p.methods[r.Method]; ok {
		m.ServeHTTP(w, r)
	}
	noSuchMethodInPath(w, r)
}

func (p *pathNode) addPath(path []string) PathNode {
	if len(path) != 1 {
		if _, ok := p.childNodes[path[1]]; !ok {
			newNode := NewPathNode()
			p.childNodes[path[1]] = newNode
		}
		return p.childNodes[path[1]].addPath(path[1:len(path)])
	}
	return p
}

func (p *pathNode) AddPath(path string) PathNode {
	return p.addPath(strings.Split(path, "/"))
}

// NewPathNode return new empty PathNode
func NewPathNode() PathNode {
	return &pathNode{make(map[string]PathNode), make(map[string]http.Handler)}
}

type pathMethodHandler struct {
	f func(http.ResponseWriter, *http.Request)
}

func (pmh *pathMethodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pmh.f(w, r)
}
