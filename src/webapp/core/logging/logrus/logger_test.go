package logrus

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"webapp/core/logging"

	"github.com/stretchr/testify/assert"
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
	expectedMessage := "This is a log to stdout"

	// Set Stdout writer
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	// set the stdout as the pipe created
	os.Stdout = w
	// Set the log output as well
	log.SetOutput(w)

	log.Info(expectedMessage)

	w.Close()
	out, _ := ioutil.ReadAll(r)
	// Restore again the stdout as before
	os.Stdout = rescueStdout
	// Default std logrus
	log.SetOutput(os.Stderr)

	assert.True(t, strings.Contains(string(out), expectedMessage))
	assert.True(t, strings.Contains(strings.ToLower(string(out)), "info"))
}

func TestDebugf(t *testing.T) {
	log.Debugf("This is a debug log %s", "with formatting")
}

func TestDebugWf(t *testing.T) {
	log.DebugWf(logging.Fields{"name": "test", "version": "1.11"},
		"This is an debug log ", "with fields")
}

func TestInfo(t *testing.T) {
	log.Info("This is an info log")
}

func TestInfof(t *testing.T) {
	log.Infof("This is an info log %s", "with formatting")
}

func TestInfoWf(t *testing.T) {
	log.InfoWf(logging.Fields{"name": "test", "version": "1.11"},
		"This is an info log ", "with fields")
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
