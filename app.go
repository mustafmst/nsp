package nsp

import (
	"fmt"
	"net/http"
)

// App main class for NSP app
type App struct {
	configBuilder ConfigurationBuilder
	logger        Logger
	debugEnabled  bool
}

// NewApp - function creating new App instance
func NewApp() *App {
	return new(App)
}

// DebugMode - toggle debug mode in app
func (a *App) DebugMode(option bool) AppInterface {
	a.debugEnabled = option
	return a
}

// UseBuilder - provides configuration builder for app
func (a *App) UseBuilder(b ConfigurationBuilder) AppInterface {
	a.logDebugInfo("UseBuilder")
	a.configBuilder = b
	a.safeLogInfo("Added ConfigurationBuilder")
	return a
}

// UseLogger - provides logger for app
func (a *App) UseLogger(l Logger) AppInterface {
	a.logDebugInfo("UseLogger")
	a.logger = l
	a.safeLogInfo("Added Logger")
	return a
}

// Run - initiates NSP app
func (a *App) Run() AppInterface {
	a.logDebugInfo("Run")
	if a.logger == nil {
		fmt.Println("No Logger provided! - ending app")
		return a
	}
	if a.configBuilder == nil {
		a.logger.LogError("No ConfigurationBuilder provided!")
		return a
	}
	a.logger.LogInfo("Starting NSP app.")
	a.configBuilder.ConfigRoutes(a.logger)
	a.serve()
	a.logger.LogInfo("Closing NSP app.")
	return a
}

func (a *App) logDebugInfo(funcName string) {
	if a.debugEnabled {
		if a.logger == nil {
			fmt.Println("--->", funcName)
		} else {
			a.logger.LogDebug("function: " + funcName)
		}
	}
}

func (a *App) safeLogInfo(s string) {
	if a.logger != nil {
		a.logger.LogInfo(s)
	} else {
		fmt.Println(s)
	}
}

func (a *App) serve() {
	a.logDebugInfo("serve")
	a.logger.LogInfo("Welcome to NSP!")
	http.ListenAndServe(":8000", nil)
}
