//Copyright (c) 2012 Nova Roma. All rights reserved.

package log

import (
	"io"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
)

// A Logger is a configurable object which logs to a writer. It has one or more log levels; each of which can be 
// configured individually.
type Logger struct {
	// Name is a unique identifier for the Logger.
	Name   string
	levels map[int]*logLevel
}

type logLevel struct {
	level  int
	output io.Writer
}

var loggerCache map[string]*Logger

// CreateLogger allocates a new logger object and adds it to the cache. 
func CreateLogger(name string, out io.Writer) *Logger {
	l := &Logger{
		Name: name,
		levels: map[int]*logLevel{
			DEBUG: &logLevel{DEBUG, out},
			INFO:  &logLevel{INFO, out},
			WARN:  &logLevel{WARN, out},
			ERROR: &logLevel{ERROR, out},
		},
	}

	if loggerCache == nil {
		loggerCache = make(map[string]*Logger)
	}
	loggerCache[name] = l
	return l
}
