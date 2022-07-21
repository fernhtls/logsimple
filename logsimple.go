 package logsimple

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

// Constants with the log Formats
const (
	logFormatString string = `[%s] [%s] [%s] [%s:%v] [%s]`
	logFormatJSON   string = `{"timestamp":"%s","level":"%s","function":"%s","file":"%s:%v","msg":"%s"}`
)

type logFormat int

// Enums for the message format
const (
	// LogFormatString : Constant for initiating the logger format string to Text
	LogFormatString logFormat = iota
	// LogFormatJSON : Constant for initiating the logger format string to Text
	LogFormatJSON
)

// Logger : Logger struct to wrap the log internal package into methods
type Logger struct {
	log *log.Logger
	// Can receive default in the string and then it will use the default
	dateFormat        string
	logFormat         logFormat
	logFormatInternal string
}

// NewLogger : returns a new non initialised logger
func NewLogger() *Logger {
	return &Logger{}
}

// SetDateFormat : Sets the log format for the builder
func (l *Logger) SetDateFormat(dateFormat string) *Logger {
	l.dateFormat = dateFormat
	return l
}

// SetLogFormat : Sets the log format for the builder
func (l *Logger) SetLogFormat(logFormat logFormat) *Logger {
	l.logFormat = logFormat
	return l
}

// GetDateFormat : Returns dateFormat
func (l *Logger) GetDateFormat() string {
	return l.dateFormat
}

// GetDateFormat : Returns dateFormat
func (l *Logger) GetLogFormat() string {
	switch l.logFormat {
	case LogFormatString:
		return "LogFormatString"
	case LogFormatJSON:
		return "LogFormatJSON"
	default: // unknown format
		return ""
	}
}

// Init : Initialises the logger. Logger always need to initialized
func (l *Logger) Init() *Logger {
	// Initializing a new logger without a prefix
	l.log = log.New(os.Stdout, "", 0)
	// Initialize date format and log format
	l.setFormats()
	return l
}

func (l *Logger) setFormats() {
	var defaultLogFormat bool
	// Sets the log format
	switch l.logFormat {
	case LogFormatString:
		l.logFormatInternal = logFormatString
	case LogFormatJSON:
		l.logFormatInternal = logFormatJSON
	default:
		l.logFormatInternal = logFormatString
		defaultLogFormat = true
	}
	// Date Formats SET
	var defaultDateFormat bool
	t := time.Now()
	ts := t.Format(l.dateFormat)
	// Transforming ts back in time to test the Format
	tts, err := time.Parse(l.dateFormat, ts)
	if tts.Round(time.Second).Before(t.Round(time.Second)) || err != nil {
		// Setting the default format if gthe format is not valid
		l.dateFormat = "2006-01-02T15:04:05.000 MST"
		defaultDateFormat = true
	}
	if defaultLogFormat {
		l.Info("Using the default log format \"logFormatString\"")
		l.Info("Value passed as log format was invalid [not logFormatString or logFormatJSON]")
	}
	if defaultDateFormat {
		l.Info(fmt.Sprintf("Using the default date format \"%v\" as the date format passed was invalid or blank", l.dateFormat))
	}
}

// Info : Logs the message as INFO level
func (l *Logger) Info(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.log.Printf(l.logFormatInternal, time.Now().Format(l.dateFormat), "INFO", funcDetail.Name(), file, line, msg)
}

// Warning : Logs the msg as WARN level
func (l *Logger) Warning(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.log.Printf(l.logFormatInternal, time.Now().Format(l.dateFormat), "WARN", funcDetail.Name(), file, line, msg)
}

// Error : Logs the message as ERROR level - but does not kill or error the process
func (l *Logger) Error(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.log.Printf(l.logFormatInternal, time.Now().Format(l.dateFormat), "ERROR", funcDetail.Name(), file, line, msg)
}

// Fatal : Logs the message as FATAL and kills the process
func (l *Logger) Fatal(abort bool, msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	if abort {
		l.log.Fatalf(l.logFormatInternal, time.Now().Format(l.dateFormat), "FATAL", funcDetail.Name(), file, line, msg)
	}
	l.log.Printf(l.logFormatInternal, time.Now().Format(l.dateFormat), "FATAL", funcDetail.Name(), file, line, msg)
}
