package nsp

import (
	"fmt"
	"net/http"
)

// App - main application interface
type App interface {
	DebugMode(option bool) App
	UseRouter(router http.Handler) App
	UseLogger(l Logger) App
	Run() App
}

// app main class for NSP app
type app struct {
	router       http.Handler
	logger       Logger
	debugEnabled bool
}

// NewApp - function creating new App instance
func NewApp() App {
	return new(app)
}

// DebugMode - toggle debug mode in app
func (a *app) DebugMode(option bool) App {
	a.debugEnabled = option
	return a
}

// UseRouter - provides app router
func (a *app) UseRouter(router http.Handler) App {
	a.logDebugInfo("UseBuilder")
	a.router = router
	a.safeLogInfo("Added ConfigurationBuilder")
	return a
}

// UseLogger - provides logger for app
func (a *app) UseLogger(l Logger) App {
	a.logDebugInfo("UseLogger")
	a.logger = l
	a.safeLogInfo("Added Logger")
	return a
}

// Run - initiates NSP app
func (a *app) Run() App {
	a.logDebugInfo("Run")
	if a.logger == nil {
		fmt.Println("No Logger provided! - ending app")
		return a
	}
	if a.router == nil {
		a.logger.LogError("No Router provided!")
		return a
	}
	a.logger.LogInfo("Starting NSP app.")
	a.serve()
	a.logger.LogInfo("Closing NSP app.")
	return a
}

func (a *app) logDebugInfo(funcName string) {
	if a.debugEnabled {
		if a.logger == nil {
			fmt.Println("--->", funcName)
		} else {
			a.logger.LogDebug("function: " + funcName)
		}
	}
}

func (a *app) safeLogInfo(s string) {
	if a.logger != nil {
		a.logger.LogInfo(s)
	} else {
		fmt.Println(s)
	}
}

func (a *app) serve() {
	a.logDebugInfo("serve")
	a.logger.LogInfo("Welcome to NSP!")
	http.ListenAndServe(":8000", a.router)
}
