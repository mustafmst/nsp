package controllers

import (
	"net/http"
)

type ControllersMap interface {
	GetControllerMethod(string, string) func(w http.ResponseWriter, r *http.Request)
}

type controllersMap struct {
	controllers map[string]Controller
}

func (cm *controllersMap) addController(name string, controller Controller) {
	cm.controllers[name] = controller
}

func (cm *controllersMap) getController(name string) (Controller, bool) {
	c, ok := cm.controllers[name]
	if ok {
		return c, true
	}
	return c, false
}

func (cm *controllersMap) GetControllerMethod(controller string, method string) func(w http.ResponseWriter, r *http.Request) {
	c, ok := cm.getController(controller)
	if ok {
		return c.GetMethod(method)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("There is no " + controller + " controller!"))
	}
}

func NewControllersMap() ControllersMap {
	cm := &controllersMap{make(map[string]Controller)}
	cm.addController(NewHomeController())
	return cm
}
