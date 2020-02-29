package logsimple

import (
	"log"
	"os"
	"runtime"
	"time"
)

// Internal package variables that will be initialized on Init
var logFormat string

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
	DateFormat string
	LogFormat  int
}

// Init : Initialez the logger. Logger always need to initialized
func (l *Logger) Init() {
	// Initializing a new logger without a prefix
	l.Log = log.New(os.Stdout, "", 0)
	// Initialize date format
	l.setDateFormat()
	// Initialize format
	l.setLogFormat()
}

func (l *Logger) setLogFormat() {
	switch l.LogFormat {
	case LogFormatString:
		logFormat = logFormatString
	case LogFormatJSON:
		logFormat = logFormatJSON
	default:
		log.Println("You didn't pass a valid LogFormat [LogFormatString, LogFormatJSON]")
		log.Println("LogFormatString will be used as default")
		logFormat = logFormatString
	}
}

func (l *Logger) setDateFormat() {
	t := time.Now()
	ts := t.Format(l.DateFormat)
	// Transforming ts back in time to test the Format
	tts, err := time.Parse(l.DateFormat, ts)
	if tts != t || err != nil {
		log.Println("Invalid date format, check the format:", l.DateFormat)
		// Setting the default format
		l.DateFormat = "2006-01-02T15:04:05.000Z"
	}
}

// Info : Logs the message as INFO level
func (l *Logger) Info(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.Log.Printf(logFormat, time.Now().Format(l.DateFormat), "INFO", funcDetail.Name(), file, line, msg)
}

// Warning : Logs the message as WARN level
func (l *Logger) Warning(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.Log.Printf(logFormat, time.Now().Format(l.DateFormat), "WARN", funcDetail.Name(), file, line, msg)
}

// Error : Logs the message as ERROR level - but does not kill or error the process
func (l *Logger) Error(msg interface{}) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	l.Log.Printf(logFormat, time.Now().Format(l.DateFormat), "ERROR", funcDetail.Name(), file, line, msg)
}

// Fatal : Logs the message as FATAL and kills the process
func (l *Logger) Fatal(msg interface{}, abort bool) {
	pc, file, line, _ := runtime.Caller(1)
	funcDetail := runtime.FuncForPC(pc)
	if abort {
		// can't be included in unit test as it will abort execution
		l.Log.Fatalf(logFormat, time.Now().Format(l.DateFormat), "FATAL", funcDetail.Name(), file, line, msg)
	}
	l.Log.Printf(logFormat, time.Now().Format(l.DateFormat), "ERROR", funcDetail.Name(), file, line, msg)
}
