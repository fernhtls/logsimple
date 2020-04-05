package logsimple

import (
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

// Enums for the message format
const (
	// LogFormatString : Constant for initiating the logger format string to Text
	LogFormatString = iota
	// LogFormatJSON : Constant for initiating the logger format string to Text
	LogFormatJSON
)

// Logger : Logger struct to wrap the log internal package into methods
type Logger struct {
	Log *log.Logger
	// Can receive default in the string and then it will use the default
	DateFormat        string
	LogFormat         int
	logFormatInternal string
}

// Init : Initialez the logger. Logger always need to initialized
func (l *Logger) Init() {
	// Initializing a new logger without a prefix
	l.Log = log.New(os.Stdout, "", 0)
	// Initialize date format and log format
	l.setFormats()
}

func (l *Logger) setFormats() {
	var defaultLogFormat bool
	// Sets the log format
	switch l.LogFormat {
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
	ts := t.Format(l.DateFormat)
	// Transforming ts back in time to test the Format
	tts, err := time.Parse(l.DateFormat, ts)
	if t.Equal(tts) || err != nil {
		// Setting the default format if gthe format is not valid
		l.DateFormat = "2006-01-02T15:04:05.000Z"
		defaultDateFormat = true
	}
	if defaultLogFormat {
		l.Info("Using the default log format \"logFormatString\"")
		l.Info("Value passed as log format was invalid [not logFormatString or logFormatJSON]")
	}
	if defaultDateFormat {
		l.Info("Using the default date format \"2006-01-02T15:04:05.000Z\" as the date format passed was invalid")
	}
}

// Info : Logs the message as INFO level
func (l *Logger) Info(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.Log.Printf(l.logFormatInternal, time.Now().Format(l.DateFormat), "INFO", funcDetail.Name(), file, line, msg)
}

// Warning : Logs the msg as WARN level
func (l *Logger) Warning(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.Log.Printf(l.logFormatInternal, time.Now().Format(l.DateFormat), "WARN", funcDetail.Name(), file, line, msg)
}

// Error : Logs the message as ERROR level - but does not kill or error the process
func (l *Logger) Error(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.Log.Printf(l.logFormatInternal, time.Now().Format(l.DateFormat), "ERROR", funcDetail.Name(), file, line, msg)
}

// Fatal : Logs the message as FATAL and kills the process
func (l *Logger) Fatal(abort bool, msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	if abort {
		l.Log.Fatalf(l.logFormatInternal, time.Now().Format(l.DateFormat), "FATAL", funcDetail.Name(), file, line, msg)
	}
	l.Log.Printf(l.logFormatInternal, time.Now().Format(l.DateFormat), "FATAL", funcDetail.Name(), file, line, msg)
}
