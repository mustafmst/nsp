package main

import (
	"github.com/mustafmst/nsp"
	"github.com/mustafmst/nsp/controllers"
	"github.com/mustafmst/nsp/sample"
)

func main() {
	nsp.NewApp().
		UseLogger(&sample.Logger{}).
		UseBuilder(sample.NewBasicBuilder(getControllersMap())).
		Run()
}

func getControllersMap() controllers.ControllersMap {
	cm := controllers.NewControllersMap()
	cm.AddController(sample.NewHomeController())
	return cm
}
