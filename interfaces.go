package nsp

import (
	"net/http"
)

// Logger interface
type Logger interface {
	LogInfo(string)
	LogDebug(string)
	LogError(string)
}

// AppInterface - main application interface
type AppInterface interface {
	DebugMode(option bool) AppInterface
	UseRouter(router http.Handler) AppInterface
	UseLogger(l Logger) AppInterface
	Run() AppInterface
}
