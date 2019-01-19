package logrus

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"webapp/core/logging"
)

var log *Logger

func init() {
	log = New()
	log.SetLevel(logging.DebugLevel)
	log.SetFormatter(logging.TextFormat)
	// Default std writer in logrus
	log.SetOutput(os.Stderr)
}

func TestJSONFormat(t *testing.T) {
	log.SetFormatter(logging.JSONFormat)
	log.Debug("This is a debug log")
	log.SetFormatter(logging.TextFormat)
}

func TestDebug(t *testing.T) {
	log.Debug("This is a debug log")
}

func TestDebugInLevelInfo(t *testing.T) {
	log.SetLevel(logging.InfoLevel)
	log.Debug("This is a debug log that must not be traced")
	log.SetLevel(logging.DebugLevel)
}

func TestStdoutOutput(t *testing.T) {
	// Set Stdout writter
	log.SetOutput(os.Stdout)
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	log.Print("This is a debug log to stdout")
	log.Error("This is an error log to stdout")

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout
	fmt.Printf("Captured: %s\n", out)
	// Default std logrus
	log.SetOutput(os.Stderr)
}

func TestDebugf(t *testing.T) {
	log.Debugf("This is a debug log %s", "with formatting")
}

func TestInfo(t *testing.T) {
	log.Info("This is an info log")
}

func TestInfof(t *testing.T) {
	log.Infof("This is an info log %s", "with formatting")
}

func TestWarn(t *testing.T) {
	log.Warn("This is a warn log")
}

func TestWarnf(t *testing.T) {
	log.Warnf("This is a warn log %s", "with formatting")
}

func TestError(t *testing.T) {
	log.Error("This is an error log")
}

func TestErrorf(t *testing.T) {
	log.Errorf("This is an error log %s", "with formatting")
}

func TestPanic(t *testing.T) {
	// Quits as expected
	//log.Panic("This is an panic log")
}

func TestPanicf(t *testing.T) {
	// Quits as expected
	//log.Panicf("This is an panic log %s", "with formatting")
}
