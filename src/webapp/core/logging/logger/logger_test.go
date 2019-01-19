package logger

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"webapp/core/logging"
)

var logger *Logger

func init() {
	logger = New()
	logger.SetLevel(logging.DebugLevel)
	logger.SetFormatter(logging.TextFormat)
	// Default std writer in logrus
	logger.SetOutput(os.Stderr)
}

func TestJSONFormat(t *testing.T) {
	logger.SetFormatter(logging.JSONFormat)
	logger.Debug("This is a debug log")
	logger.SetFormatter(logging.TextFormat)
}

func TestDebug(t *testing.T) {
	logger.Debug("This is a debug log")
}

func TestDebugInLevelInfo(t *testing.T) {
	logger.SetLevel(logging.InfoLevel)
	logger.Debug("This is a debug log that must not be traced")
	logger.SetLevel(logging.DebugLevel)
}

func TestStdoutOutput(t *testing.T) {
	// Set Stdout writter
	logger.SetOutput(os.Stdout)
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger.Print("This is a debug log to stdout")
	logger.Error("This is an error log to stdout")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	fmt.Printf("Captured: %s\n", out)
	// Default std logrus
	logger.SetOutput(os.Stderr)
}

func logAll(logLevel string) {
	logger.Debug("This is a debug log", logLevel)
	logger.Info("This is an info log", logLevel)
	logger.Warn("This is a warn log", logLevel)
	logger.Error("This is an error log", logLevel)
	logger.Fatal("This is a fatal log", logLevel)
	//logger.Panic("This is an info log", logLevel)
}

func TestLevels(t *testing.T) {
	logger.SetLevel(logging.DebugLevel)
	logAll("(DebugLevel)")

	logger.SetLevel(logging.InfoLevel)
	logAll("(InfoLevel)")

	logger.SetLevel(logging.WarnLevel)
	logAll("(WarnLevel)")

	logger.SetLevel(logging.ErrorLevel)
	logAll("(ErrorLevel)")

	logger.SetLevel(logging.FatalLevel)
	logAll("(FatalLevel)")

	logger.SetLevel(logging.PanicLevel)
	logAll("(PanicLevel)")

	// To defaults
	logger.SetLevel(logging.DebugLevel)
}

func TestDebugf(t *testing.T) {
	logger.Debugf("This is a debug log %s", "with formatting")
}

func TestInfo(t *testing.T) {
	logger.Info("This is an info log")
}

func TestInfof(t *testing.T) {
	logger.Infof("This is an info log %s", "with formatting")
}

func TestWarn(t *testing.T) {
	logger.Warn("This is a warn log")
}

func TestWarnf(t *testing.T) {
	logger.Warnf("This is a warn log %s", "with formatting")
}

func TestError(t *testing.T) {
	logger.Error("This is an error log")
}

func TestErrorf(t *testing.T) {
	logger.Errorf("This is an error log %s", "with formatting")
}

func TestPanic(t *testing.T) {
	// Quits as expected
	//logger.Panic("This is an panic log")
}

func TestPanicf(t *testing.T) {
	// Quits as expected
	//logger.Panicf("This is an panic log %s", "with formatting")
}