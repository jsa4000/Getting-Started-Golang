package logging

import (
	"io"
)

// Log Global logger
var Log Logger

// Level type
type Level uint32

// Format type
type Format uint

// These are the different logging levels.
const (
	TextFormat Format = iota
	JSONFormat
)

// These are the different logging levels.
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// Fields type, used to pass to `WithFields`.
type Fields map[string]interface{}

// Logger The Logger interface generalizes the Entry and Logger types
type Logger interface {
	SetFormatter(format Format)
	SetLevel(level Level)
	SetOutput(out io.Writer)

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}
