package logging

import (
	"io"
)

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

	DebugWf(fields Fields, args ...interface{})
	InfoWf(fields Fields, args ...interface{})
	PrintWf(fields Fields, args ...interface{})
	WarnWf(fields Fields, args ...interface{})
	WarningWf(fields Fields, args ...interface{})
	ErrorWf(fields Fields, args ...interface{})
	FatalWf(fields Fields, args ...interface{})
	PanicWf(fields Fields, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Print(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
}

// Log Global logger
var logger Logger

// SetGlobal sets the Global Logger (singletone)
func SetGlobal(l Logger) {
	logger = l
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(format Format) {
	logger.SetFormatter(format)
}

// SetLevel sets the standard logger level.
func SetLevel(level Level) {
	logger.SetLevel(level)
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	logger.SetOutput(out)
}

// Debugf log
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Infof log
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Printf log
func Printf(format string, args ...interface{}) {
	logger.Printf(format, args...)
}

// Warnf log
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Warningf log
func Warningf(format string, args ...interface{}) {
	logger.Warningf(format, args...)
}

// Errorf log
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatalf log
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panicf log
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// DebugWf log
func DebugWf(fields Fields, args ...interface{}) {
	logger.DebugWf(fields, args...)
}

// InfoWf log
func InfoWf(fields Fields, args ...interface{}) {
	logger.InfoWf(fields, args...)
}

// PrintWf log
func PrintWf(fields Fields, args ...interface{}) {
	logger.PrintWf(fields, args...)
}

// WarnWf log
func WarnWf(fields Fields, args ...interface{}) {
	logger.WarnWf(fields, args...)
}

// WarningWf log
func WarningWf(fields Fields, args ...interface{}) {
	logger.WarningWf(fields, args...)
}

// ErrorWf log
func ErrorWf(fields Fields, args ...interface{}) {
	logger.ErrorWf(fields, args...)
}

// FatalWf log
func FatalWf(fields Fields, args ...interface{}) {
	logger.FatalWf(fields, args...)
}

// PanicWf log
func PanicWf(fields Fields, args ...interface{}) {
	logger.PanicWf(fields, args...)
}

// Debug log
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Info log
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Print log
func Print(args ...interface{}) {
	logger.Print(args...)
}

// Warn log
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warning log
func Warning(args ...interface{}) {
	logger.Warning(args...)
}

// Error log
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Fatal log
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Panic log
func Panic(args ...interface{}) {
	logger.Panic(args...)
}
