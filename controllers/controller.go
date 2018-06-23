package controllers

import (
	"net/http"
)

type Controller interface {
	GetMethod(string) func(w http.ResponseWriter, r *http.Request)
	AddMethod(string, func(w http.ResponseWriter, r *http.Request))
	GetName() string
}

type controller struct {
	name    string
	methods map[string]func(w http.ResponseWriter, r *http.Request)
}

func (c *controller) GetMethod(name string) func(w http.ResponseWriter, r *http.Request) {
	m, ok := c.methods[name]
	if ok {
		return m
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("There is no " + name + " method in " + c.GetName() + " controller!"))
	}
}

func (c *controller) AddMethod(name string, method func(w http.ResponseWriter, r *http.Request)) {
	c.methods[name] = method
}

func (c *controller) GetName() string {
	return c.name
}

func NewController(name string) Controller {
	return &controller{name, make(map[string]func(w http.ResponseWriter, r *http.Request))}
}
