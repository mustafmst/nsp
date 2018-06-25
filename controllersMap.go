package nsp

import (
	"net/http"
)

// ControllersMap provides interface for controllers managment
type ControllersMap interface {
	GetControllerMethod(string, string) func(w http.ResponseWriter, r *http.Request)
	AddController(Controller)
}

type controllersMap struct {
	controllers map[string]Controller
}

func (cm *controllersMap) AddController(controller Controller) {
	cm.controllers[controller.GetName()] = controller
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

// NewControllersMap creates new empty controllers map
func NewControllersMap() ControllersMap {
	cm := &controllersMap{make(map[string]Controller)}
	return cm
}
