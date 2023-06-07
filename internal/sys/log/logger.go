package log

import (
	"log"
	"os"
	"strings"
)

type Level int

const (
	Debug Level = iota
	Info
	Error
)

type (
	Logger interface {
		Debug(v ...interface{})
		Info(v ...interface{})
		Error(v ...interface{})
	}

	BaseLogger struct {
		debugLogger *log.Logger
		infoLogger  *log.Logger
		errorLogger *log.Logger
		logLevel    Level
	}
)

func NewLogger(level string) *BaseLogger {
	logLevel := toValidLevel(level)
	return &BaseLogger{
		debugLogger: log.New(os.Stdout, "[DBG] ", log.LstdFlags),
		infoLogger:  log.New(os.Stdout, "[INF] ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "[ERR] ", log.LstdFlags),
		logLevel:    logLevel,
	}
}

func toValidLevel(level string) Level {
	level = strings.ToLower(level)

	switch level {
	case "debug", "dbg":
		return Debug
	case "info", "inf":
		return Info
	case "error", "err":
		return Error
	default:
		return Error
	}
}

func (l *BaseLogger) SetLogLevel(level Level) {
	l.logLevel = level
}

func (l *BaseLogger) Debug(v ...interface{}) {
	if l.logLevel <= Debug {
		l.debugLogger.Println(v...)
	}
}

func (l *BaseLogger) Info(v ...interface{}) {
	if l.logLevel <= Info {
		l.infoLogger.Println(v...)
	}
}

func (l *BaseLogger) Error(v ...interface{}) {
	if l.logLevel <= Error {
		l.errorLogger.Println(v...)
	}
}
