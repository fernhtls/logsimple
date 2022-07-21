package logsimple

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestLogSimple : Simple test for Logging
func TestLogSimple(t *testing.T) {
	// Logger no DateFormat
	// Logger with String format
	loggerND := NewLogger().Init()
	// loggerS.DateFormat = "2006-01-02T15:04:05.000 MST"
	loggerND.Info("Test INFO message")
	loggerND.Warning("Test WARN message")
	loggerND.Error("Test ERROR message")
	// Logger with String format
	loggerS := NewLogger().
		SetDateFormat("2006-01-02T15:04:05.000 MST").
		Init()
	loggerS.Info("Test INFO message")
	loggerS.Warning("Test WARN message")
	loggerS.Error("Test ERROR message")
	//// Logger with JSON format
	loggerJ := NewLogger().
		SetDateFormat("2006-01-02T15:04:05.000 MST").
		SetLogFormat(LogFormatJSON).
		Init()
	loggerJ.Info("Hey I'm logging info to my app!")
	loggerJ.Warning("Test WARN message")
	loggerJ.Error("Test ERROR message")
	loggerJ.Fatal(false, "Test ERROR message")
	//// Getting the default log Format
	//// Logger with String format as the default
	loggerD := NewLogger().
		SetDateFormat("2006-01-02T15:04:05.000 MST").
		SetLogFormat(10).
		Init()
	loggerD.Info("Hey I'm logging info to my app!")
	loggerD.Warning("Test WARN message")
	loggerD.Error("Test ERROR message")
	//// Testing the setDateFormat function
	loggerDFError := NewLogger().
		SetDateFormat("2006-01-02T15:04:AA.000 MST").
		Init()
	// Calling setDateFormat
	loggerDFError.Info("Test incorret Date format")
	t.Log("No errors on execution of the test")
}

func TestAbort(t *testing.T) {
	// Env var used to run abort only in the subprocess
	// only when it reaches cmd.Run()
	// Otherwise abort would run in the beginning of the test
	if os.Getenv("RUN_ABORT") == "1" {
		loggerFatal := NewLogger().Init()
		loggerFatal.Fatal(true, "Abort!")
		return
	}
	// os.Args[0] - temp go build file when running the test
	// /tmp/go-build986654860/b001/logsimple.test
	cmd := exec.Command(os.Args[0], "-test.run=TestAbort")
	// Sets RUN_ABORT to 1 before cmd.Run()
	cmd.Env = append(os.Environ(), "RUN_ABORT=1")
	err := cmd.Run()
	ex, ok := err.(*exec.ExitError)
	assert.Equal(t, strings.Contains(ex.String(), "exit status 1"), true)
	assert.Equal(t, ok, true)
	t.Log("Application was aborted correctly with message : \"", ex.String(), "\"")
}
