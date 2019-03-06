package wechat

import (
	"log"
)

// Logger specifies the interface for all log operations.
type Logger interface {
	Printf(format string, v ...interface{})
}

// DefaultLogger DefaultLogger
type DefaultLogger struct{}

// NewDefaultLogger NewDefaultLogger
func NewDefaultLogger() *DefaultLogger {
	return &DefaultLogger{}
}

// Printf Printf
func (l *DefaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
