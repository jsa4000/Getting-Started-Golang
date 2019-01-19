package logger

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"webapp/core/logging"
)

// Logger struct to store all the logrus data
type Logger struct {
	Log   *log.Logger
	Level logging.Level
}

// New Creates new Default Logger instance
func New() *Logger {
	return &Logger{
		Log:   log.New(os.Stderr, "", 0),
		Level: logging.DebugLevel,
	}
}

// SetFormatter sets the standard logger formatter.
func (l *Logger) SetFormatter(format logging.Format) {
}

// SetLevel sets the standard logger level.
func (l *Logger) SetLevel(level logging.Level) {
	l.Level = level
}

// SetOutput sets the standard logger output.
func (l *Logger) SetOutput(out io.Writer) {
	l.Log = log.New(out, "", 0)
}

func getLevelString(level logging.Level) string {
	switch level {
	case logging.DebugLevel:
		return "DEBUG"
	case logging.InfoLevel:
		return "INFO"
	case logging.WarnLevel:
		return "WARN"
	case logging.ErrorLevel:
		return "ERROR"
	case logging.FatalLevel:
		return "FATAL"
	case logging.PanicLevel:
		return "PANIC"
	default:
		return "UNKNOWN"
	}
}

func getPrefix(level logging.Level) string {
	return fmt.Sprintf("%s [%s]:", time.Now().Format(time.RFC3339), getLevelString(level))
}

func getPrefixWithFields(level logging.Level, fields logging.Fields) string {
	return fmt.Sprintf("%s [%s]: %s, ", time.Now().Format(time.RFC3339), getLevelString(level), fields)
}

// Debugf log
func (l *Logger) Debugf(format string, args ...interface{}) {
	if l.Level < logging.DebugLevel {
		return
	}
	l.Log.Printf(getPrefix(logging.DebugLevel)+format, args...)
}

// Infof log
func (l *Logger) Infof(format string, args ...interface{}) {
	if l.Level < logging.InfoLevel {
		return
	}
	l.Log.Printf(getPrefix(logging.InfoLevel)+format, args...)
}

// Printf log
func (l *Logger) Printf(format string, args ...interface{}) {
	l.Log.Printf(getPrefix(0)+format, args...)
}

// Warnf log
func (l *Logger) Warnf(format string, args ...interface{}) {
	if l.Level < logging.WarnLevel {
		return
	}
	l.Log.Printf(getPrefix(logging.WarnLevel)+format, args...)
}

// Warningf log
func (l *Logger) Warningf(format string, args ...interface{}) {
	if l.Level < logging.WarnLevel {
		return
	}
	l.Log.Printf(getPrefix(logging.WarnLevel)+format, args...)
}

// Errorf log
func (l *Logger) Errorf(format string, args ...interface{}) {
	if l.Level < logging.ErrorLevel {
		return
	}
	l.Log.Printf(getPrefix(logging.ErrorLevel)+format, args...)
}

// Fatalf log
func (l *Logger) Fatalf(format string, args ...interface{}) {
	if l.Level < logging.FatalLevel {
		return
	}
	l.Log.Printf(getPrefix(logging.FatalLevel)+format, args...)
}

// Panicf log
func (l *Logger) Panicf(format string, args ...interface{}) {
	if l.Level < logging.PanicLevel {
		return
	}
	l.Log.Printf(getPrefix(logging.PanicLevel)+format, args...)
	panic(fmt.Sprintf(getPrefix(logging.PanicLevel)+format, args...))
}

// DebugWf log
func (l *Logger) DebugWf(fields logging.Fields, args ...interface{}) {
	if l.Level < logging.DebugLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefixWithFields(logging.DebugLevel, fields)}, args...))
}

// InfoWf log
func (l *Logger) InfoWf(fields logging.Fields, args ...interface{}) {
	if l.Level < logging.InfoLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefixWithFields(logging.InfoLevel, fields)}, args...))
}

// PrintWf log
func (l *Logger) PrintWf(fields logging.Fields, args ...interface{}) {
	l.Log.Println(append([]interface{}{getPrefixWithFields(0, fields)}, args...))
}

// WarnWf log
func (l *Logger) WarnWf(fields logging.Fields, args ...interface{}) {
	if l.Level < logging.WarnLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefixWithFields(logging.WarnLevel, fields)}, args...))
}

// WarningWf log
func (l *Logger) WarningWf(fields logging.Fields, args ...interface{}) {
	if l.Level < logging.WarnLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefixWithFields(logging.WarnLevel, fields)}, args...))
}

// ErrorWf log
func (l *Logger) ErrorWf(fields logging.Fields, args ...interface{}) {
	if l.Level < logging.ErrorLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefixWithFields(logging.ErrorLevel, fields)}, args...))
}

// FatalWf log
func (l *Logger) FatalWf(fields logging.Fields, args ...interface{}) {
	if l.Level < logging.FatalLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefixWithFields(logging.FatalLevel, fields)}, args...))
}

// PanicWf log
func (l *Logger) PanicWf(fields logging.Fields, args ...interface{}) {
	if l.Level < logging.PanicLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefixWithFields(logging.PanicLevel, fields)}, args...))
}

// Debug log
func (l *Logger) Debug(args ...interface{}) {
	if l.Level < logging.DebugLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefix(logging.DebugLevel)}, args...))
}

// Info log
func (l *Logger) Info(args ...interface{}) {
	if l.Level < logging.InfoLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefix(logging.InfoLevel)}, args...))
}

// Print log
func (l *Logger) Print(args ...interface{}) {
	l.Log.Println(append([]interface{}{getPrefix(0)}, args...))
}

// Warn log
func (l *Logger) Warn(args ...interface{}) {
	if l.Level < logging.WarnLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefix(logging.WarnLevel)}, args...))
}

// Warning log
func (l *Logger) Warning(args ...interface{}) {
	if l.Level < logging.WarnLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefix(logging.WarnLevel)}, args...))
}

// Error log
func (l *Logger) Error(args ...interface{}) {
	if l.Level < logging.ErrorLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefix(logging.ErrorLevel)}, args...))
}

// Fatal log
func (l *Logger) Fatal(args ...interface{}) {
	if l.Level < logging.FatalLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefix(logging.FatalLevel)}, args...))
}

// Panic log
func (l *Logger) Panic(args ...interface{}) {
	if l.Level < logging.PanicLevel {
		return
	}
	l.Log.Println(append([]interface{}{getPrefix(logging.PanicLevel)}, args...))
	panic(fmt.Sprintln(append([]interface{}{getPrefix(logging.PanicLevel)}, args...)))
}
