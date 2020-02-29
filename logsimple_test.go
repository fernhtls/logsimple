package logsimple

import (
	"log"
	"testing"
)

// TestLogSimple : Simple test for Logging
func TestLogSimple(t *testing.T) {
	// Logger with String format
	var loggerS Logger
	loggerS.DateFormat = "2006-01-02T15:04:05.000Z"
	loggerS.Init()
	loggerS.Info("Test INFO message")
	loggerS.Warning("Test WARN message")
	loggerS.Error("Test ERROR message")
	// Logger with JSON format
	var loggerJ Logger
	loggerJ.DateFormat = "2006-01-02T15:04:05.000Z"
	loggerJ.LogFormat = LogFormatJSON
	loggerJ.Init()
	loggerJ.Info("Hey I'm logging info to my app!")
	loggerJ.Warning("Test WARN message")
	loggerJ.Error("Test ERROR message")
	loggerS.Fatal("Test ERROR message", false)
	// Getting the default log Format
	// Logger with String format as the default
	var loggerD Logger
	loggerD.DateFormat = "2006-01-02T15:04:05.000Z"
	loggerD.LogFormat = 10
	loggerD.Init()
	loggerD.Info("Hey I'm logging info to my app!")
	loggerD.Warning("Test WARN message")
	loggerD.Error("Test ERROR message")
	// Testing the setDateFormat function
	var loggerDFError Logger
	loggerDFError.DateFormat = "2006-01-02T15:04:AA.000Z"
	// Calling setDateFormat
	log.Println("Calling setDateFormat")
	loggerDFError.setDateFormat()
	// Not including the FATAL message on unittest as it aborts the execution
	t.Log("No errors on execution of the test")
}
