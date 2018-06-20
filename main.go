package main

import (
	"github.com/mustafmst/nsp/config"
	"github.com/mustafmst/nsp/core"
)

func main() {
	Start(core.NewApp())
}

// Start - starts whole app
func Start(app core.AppInterface) {
	app.
		DebugMode(true).
		UseLogger(config.NewConsoleLogger()).
		UseBuilder(config.NewBasicBuilder()).
		Run()
}
