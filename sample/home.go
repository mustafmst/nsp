package sample

import (
	"net/http"

	"github.com/mustafmst/nsp/controllers"
)

func NewHomeController() controllers.Controller {
	c := controllers.NewController("home")
	c.AddMethod("index", index)
	c.AddMethod("info", info)
	return c
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home#index"))
}

func info(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("home#info"))
}
