package main

import (
	"testing"

	"github.com/mustafmst/nsp/core"
)

// App main class for NSP app
type app struct {
	debug   bool
	builder bool
	logger  bool
	run     bool
}

func newAppMock() *app {
	a := new(app)
	a.debug = false
	a.builder = false
	a.logger = false
	a.run = false
	return a
}

// DebugMode - toggle debug mode in app
func (a *app) DebugMode(option bool) core.AppInterface {
	a.debug = true
	return a
}

// UseBuilder - provides configuration builder for app
func (a *app) UseBuilder(b core.ConfigurationBuilder) core.AppInterface {
	a.builder = true
	return a
}

// UseLogger - provides logger for app
func (a *app) UseLogger(l core.Logger) core.AppInterface {
	a.logger = true
	return a
}

// Run - initiates NSP app
func (a *app) Run() core.AppInterface {
	a.run = true
	return a
}

func TestStart(t *testing.T) {
	assertConfiguration := func(configName string, wasSet bool) {
		if wasSet == false {
			t.Error(configName, "was not set!")
		}
	}

	t.Run("Evaluate startup configuration", func(t *testing.T) {
		a := newAppMock()
		Start(a)
		assertConfiguration("debugMode", a.debug)
		assertConfiguration("builder", a.builder)
		assertConfiguration("logger", a.logger)
		assertConfiguration("run", a.run)
	})
}
