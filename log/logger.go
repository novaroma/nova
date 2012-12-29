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

type Logger struct {
	Name   string
	levels map[int]*logLevel
}

type logLevel struct {
	level  int
	output io.Writer
}

var loggerCache map[string]*Logger

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
