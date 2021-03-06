//Copyright (c) 2012 Nova Roma. All rights reserved.

package log

import (
	"fmt"
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

// DisableLevel causes the log level specified to not output.
func (logger *Logger) DisableLevel(l int) {
	if level, ok := logger.levels[l]; ok {
		level.enabled = false
	}
}

// EnableLevel causes the log level specified to resume output.
func (logger *Logger) EnableLevel(l int) {
	if level, ok := logger.levels[l]; ok {
		level.enabled = true
	}
}

// Log writes to the log level given by the first argument. Other arguments are handled in the manner of fmt.Print. 
// It does nothing if the log level is disabled; and returns an error if the log level does not exist.
func (logger *Logger) Log(logLevel int, v ...interface{}) error {
	level := logger.levels[logLevel]
	if level != nil {
		if level.enabled {
			level.logger.Print(v...)
		}
		return nil
	}

	return fmt.Errorf("No log level '%d' exists in this logger.", logLevel)
}

// Debug writes to the debug log level. It handles arguments in the manner of fmt.Print.
// It does nothing if debug is disabled. It panics if the debug log level does not exist.
func (logger *Logger) Debug(v ...interface{}) {
	if logger.Log(LogLevelDebug, v...) != nil {
		panic("Logger must have Debug log level.")
	}
}

// Info writes to the info log level. It handles arguments in the manner of fmt.Print.
// It does nothing if info is disabled. It panics if the info log level does not exist.
func (logger *Logger) Info(v ...interface{}) {
	if logger.Log(LogLevelInfo, v...) != nil {
		panic("Logger must have Info log level.")
	}
}

// Warn writes to the warn log level. It handles arguments in the manner of fmt.Print.
// It does nothing if warn is disabled. It panics if the warn log level does not exist.
func (logger *Logger) Warn(v ...interface{}) {
	if logger.Log(LogLevelWarn, v...) != nil {
		panic("Logger must have Warn log level.")
	}
}

// Error writes to the error log level. It handles arguments in the manner of fmt.Print.
// It does nothing if error is disabled. It panics if the error log level does not exist.
func (logger *Logger) Error(v ...interface{}) {
	if logger.Log(LogLevelError, v...) != nil {
		panic("Logger must have Error log level.")
	}
}

// Panics writes to the error log level, then panics with the formatted string. It handles arguments in the manner of 
// fmt.Print.
func (logger *Logger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	logger.Log(LogLevelError, s)
	panic(s)
}

// Fatal writes to the error log level; then exits with os.Exit(1). It handles arguments in the manner of fmt.Print.
func (logger *Logger) Fatal(v ...interface{}) {
	logger.Log(LogLevelError, v...)
	os.Exit(1)
}

// Logf writes to the log level given by the first argument. Other arguments are handled in the manner of fmt.Printf.
// Logf does nothing if the log level specified is disabled; it returns an error if the log level does not exist.
func (logger *Logger) Logf(logLevel int, format string, v ...interface{}) error {
	level := logger.levels[logLevel]
	if level != nil {
		if level.enabled {
			level.logger.Printf(format, v...)
		}
		return nil
	}

	return fmt.Errorf("No log level '%d' exists in this logger.", logLevel)
}

// Debugf writes to the debug log level. It handles arguments in the manner of fmt.Printf.
// It does nothing if debug is disabled. It panics if the debug log level does not exist.
func (logger *Logger) Debugf(format string, v ...interface{}) {
	if logger.Logf(LogLevelDebug, format, v...) != nil {
		panic("Logger must have Debug log level.")
	}
}

// Infof writes to the info log level. It handles arguments in the manner of fmt.Printf.
// It does nothing if info is disabled. It panics if the info log level does not exist.
func (logger *Logger) Infof(format string, v ...interface{}) {
	if logger.Logf(LogLevelInfo, format, v...) != nil {
		panic("Logger must have Info log level.")
	}
}

// Warnf writes to the warn log level. It handles arguments in the manner of fmt.Printf.
// It does nothing if warn is disabled. It panics if the warn log level does not exist.
func (logger *Logger) Warnf(format string, v ...interface{}) {
	if logger.Logf(LogLevelWarn, format, v...) != nil {
		panic("Logger must have Warn log level.")
	}
}

// Errorf writes to the error log level. It handles arguments in the manner of fmt.Printf.
// It does nothing if error is disabled. It panics if the error log level does not exist.
func (logger *Logger) Errorf(format string, v ...interface{}) {
	if logger.Logf(LogLevelError, format, v...) != nil {
		panic("Logger must have Error log level.")
	}
}

// Panicf writes to the error log level, then panics with the formatted string. It handles arguments in the manner of 
// fmt.Printf.
func (logger *Logger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	logger.Logf(LogLevelError, s)
	panic(s)
}

// Fatalf writes to the error log level; then exits with os.Exit(1). It handles arguments in the manner of fmt.Printf.
func (logger *Logger) Fatalf(format string, v ...interface{}) {
	logger.Logf(LogLevelError, format, v...)
	os.Exit(1)
}

// Logln writes to the log level given by the first argument. Other arguments are handled in the manner of fmt.Println.
func (logger *Logger) Logln(logLevel int, v ...interface{}) error {
	level := logger.levels[logLevel]
	if level != nil {
		if level.enabled {
			level.logger.Println(v...)
		}
		return nil
	}

	return fmt.Errorf("No log level '%d' exists in this logger.", logLevel)
}

// Debugln writes to the debug log level. It handles arguments in the manner of fmt.Println.
// It does nothing if debug is disabled. It panics if the debug log level does not exist.
func (logger *Logger) Debugln(v ...interface{}) {
	if logger.Logln(LogLevelDebug, v...) != nil {
		panic("Logger must have Debug log level.")
	}
}

// Infoln writes to the info log level. It handles arguments in the manner of fmt.Println.
// It does nothing if info is disabled. It panics if the info log level does not exist.
func (logger *Logger) Infoln(v ...interface{}) {
	if logger.Logln(LogLevelInfo, v...) != nil {
		panic("Logger must have Info log level.")
	}
}

// Warnln writes to the warn log level. It handles arguments in the manner of fmt.Println.
// It does nothing if warn is disabled. It panics if the warn log level does not exist.
func (logger *Logger) Warnln(v ...interface{}) {
	if logger.Logln(LogLevelWarn, v...) != nil {
		panic("Logger must have Warn log level.")
	}
}

// Errorln writes to the error log level. It handles arguments in the manner of fmt.Println.
// It does nothing if error is disabled. It panics if the error log level does not exist.
func (logger *Logger) Errorln(v ...interface{}) {
	if logger.Logln(LogLevelError, v...) != nil {
		panic("Logger must have Error log level.")
	}
}

// Panicln writes to the error log level, then panics with the formatted string. It handles arguments in the manner of 
// fmt.Println.
func (logger *Logger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	logger.Logln(LogLevelError, s)
	panic(s)
}

// Fatalln writes to the error log level; then exits with os.Exit(1). It handles arguments in the manner of fmt.Println.
func (logger *Logger) Fatalln(v ...interface{}) {
	logger.Logln(LogLevelError, v...)
	os.Exit(1)
}

type logLevel struct {
	level   int
	logger  *log.Logger
	enabled bool
}

// CreateLogger allocates a new logger object and adds it to the cache. 
func CreateLogger(name string, out io.Writer) *Logger {
	l := &Logger{
		Name: name,
		levels: map[int]*logLevel{
			LogLevelDebug: &logLevel{LogLevelDebug, log.New(out, PrefixLogLevelDebug, log.LstdFlags), true},
			LogLevelInfo:  &logLevel{LogLevelInfo, log.New(out, PrefixLogLevelInfo, log.LstdFlags), true},
			LogLevelWarn:  &logLevel{LogLevelWarn, log.New(out, PrefixLogLevelWarn, log.LstdFlags), true},
			LogLevelError: &logLevel{LogLevelError, log.New(out, PrefixLogLevelError, log.LstdFlags), true},
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
