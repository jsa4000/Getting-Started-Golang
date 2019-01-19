package logrus

import (
	"io"
	"webapp/core/logging"

	"github.com/sirupsen/logrus"
)

// Logger struct logrus placeholder
type Logger struct {
	Log *logrus.Logger
}

// New Creates Logrus new instance
func New() *Logger {
	return &Logger{
		Log: logrus.New(),
	}
}

// SetFormatter sets the standard logger formatter.
func (l *Logger) SetFormatter(format logging.Format) {
	switch format {
	case logging.TextFormat:
		l.Log.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
	case logging.JSONFormat:
		l.Log.SetFormatter(&logrus.JSONFormatter{})
	}
}

// SetLevel sets the standard logger level.
func (l *Logger) SetLevel(level logging.Level) {
	l.Log.SetLevel(logrus.Level(level))
}

// SetOutput sets the standard logger output.
func (l *Logger) SetOutput(out io.Writer) {
	l.Log.SetOutput(out)
}

// Debugf log
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.Log.Debugf(format, args...)
}

// Infof log
func (l *Logger) Infof(format string, args ...interface{}) {
	l.Log.Infof(format, args...)
}

// Printf log
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Log.Printf(format, args...)
}

// Warnf log
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.Log.Warnf(format, args...)
}

// Warningf log
func (l *Logger) Warningf(format string, args ...interface{}) {
	l.Log.Warningf(format, args...)
}

// Errorf log
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.Log.Errorf(format, args...)
}

// Fatalf log
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.Log.Fatalf(format, args...)
}

// Panicf log
func (l *Logger) Panicf(format string, args ...interface{}) {
	l.Log.Panicf(format, args...)
}

// Debug log
func (l *Logger) Debug(args ...interface{}) {
	l.Log.Debug(args...)
}

// Info log
func (l *Logger) Info(args ...interface{}) {
	l.Log.Info(args...)
}

// Print log
func (l *Logger) Print(args ...interface{}) {
	l.Log.Print(args...)
}

// Warn log
func (l *Logger) Warn(args ...interface{}) {
	l.Log.Warn(args...)
}

// Warning log
func (l *Logger) Warning(args ...interface{}) {
	l.Log.Warning(args...)
}

// Error log
func (l *Logger) Error(args ...interface{}) {
	l.Log.Error(args...)
}

// Fatal log
func (l *Logger) Fatal(args ...interface{}) {
	l.Log.Fatal(args...)
}

// Panic log
func (l *Logger) Panic(args ...interface{}) {
	l.Log.Panic(args...)
}
