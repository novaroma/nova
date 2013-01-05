//Copyright (c) 2012 Nova Roma. All rights reserved.

package log

import (
	"io"
	"log"
	"os"
)

const (
	LogLevelDebug = iota
	LogLevelInfo
	LogLevelWarn
	LogLevelError

	PrefixLogLevelDebug = "[DEBUG] "
	PrefixLogLevelInfo  = "[INFO] "
	PrefixLogLevelWarn  = "[WARN] "
	PrefixLogLevelError = "[ERROR] "
)

var loggerCache map[string]*Logger

func init() {
	loggerCache = make(map[string]*Logger)
}

// A Logger is a configurable object which logs to a writer. It has one or more log levels; each of which can be 
// configured individually.
type Logger struct {
	// Name is a unique identifier for the Logger.
	Name   string
	levels map[int]*logLevel
}

// Log writes to the log level given by the first argument; or to the standard logger if the given log level does not 
// exist. Other arguments are handled in the manner of fmt.Print.
func (logger *Logger) Log(logLevel int, v ...interface{}) {
	level := logger.levels[logLevel]
	if level != nil {
		level.logger.Print(v...)
	} else {
		log.Print(v...)
	}
}

// Logf writes to the log level given by the first argument; or to the standard logger if the given log level does not
// exist. Other arguments are handled in the manner of fmt.Printf.
func (logger *Logger) Logf(logLevel int, format string, v ...interface{}) {
	level := logger.levels[logLevel]
	if level != nil {
		level.logger.Printf(format, v...)
	} else {
		log.Printf(format, v...)
	}
}

// Logln writes to the log level given by the first argument. Other arguments are handled in the manner of fmt.Println.
func (logger *Logger) Logln(logLevel int, v ...interface{}) {
	level := logger.levels[logLevel]
	if level != nil {
		level.logger.Println(v...)
	} else {
		log.Println(v...)
	}
}

type logLevel struct {
	level  int
	logger *log.Logger
}

// CreateLogger allocates a new logger object and adds it to the cache. 
func CreateLogger(name string, out io.Writer) *Logger {
	l := &Logger{
		Name: name,
		levels: map[int]*logLevel{
			LogLevelDebug: &logLevel{LogLevelDebug, log.New(out, PrefixLogLevelDebug, log.LstdFlags)},
			LogLevelInfo:  &logLevel{LogLevelInfo, log.New(out, PrefixLogLevelInfo, log.LstdFlags)},
			LogLevelWarn:  &logLevel{LogLevelWarn, log.New(out, PrefixLogLevelWarn, log.LstdFlags)},
			LogLevelError: &logLevel{LogLevelError, log.New(out, PrefixLogLevelError, log.LstdFlags)},
		},
	}

	loggerCache[name] = l
	return l
}

// GetLogger gets the logger with the specified name from the cache. If it does not already exist, then it is created.
func GetLogger(name string) *Logger {
	if logger, ok := loggerCache[name]; ok {
		return logger
	}

	return CreateLogger(name, os.Stdout)
}
