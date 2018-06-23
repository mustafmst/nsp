package controllers

import (
	"net/http"
)

func NewHomeController() (string, Controller) {
	c := NewController("home")
	c.AddMethod("index", index)
	c.AddMethod("info", info)
	return c.GetName(), c
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home#index"))
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home#info"))
}
