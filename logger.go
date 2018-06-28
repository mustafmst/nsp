package nsp

import (
	"fmt"
	"time"
)

// Logger interface
type Logger interface {
	LogInfo(string)
	LogDebug(string)
	LogError(string)
}

type logger struct{}

func NewLogger() Logger {
	return &logger{}
}

func (l *logger) LogInfo(msg string) {
	log("INFO ", msg)
}
func (l *logger) LogDebug(msg string) {
	log("DEBUG", msg)
}
func (l *logger) LogError(msg string) {
	log("ERROR", msg)
}

func log(level, msg string) {
	fmt.Printf("%s | %s | %s", time.Now(), level, msg)
}
