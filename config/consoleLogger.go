package config

import (
	"fmt"
	"time"

	"github.com/mustafmst/nsp/core"
)

// ConsoleLogger - logs info to console
type ConsoleLogger struct{}

// NewConsoleLogger - creates new console logger instance
func NewConsoleLogger() core.Logger {
	return &ConsoleLogger{}
}

// LogInfo - info
func (c *ConsoleLogger) LogInfo(s string) {
	fmt.Println(CreateLogMsg(s, "Info "))
}

// LogDebug - debug
func (c *ConsoleLogger) LogDebug(s string) {
	fmt.Println(CreateLogMsg(s, "Debug"))
}

// LogError - error
func (c *ConsoleLogger) LogError(s string) {
	fmt.Println(CreateLogMsg(s, "Error"))
}

// CreateLogMsg - create formatted message
func CreateLogMsg(msg, msgType string) string {
	return time.Now().UTC().Format(time.UnixDate) + " | " + msgType + " | " + msg
}
