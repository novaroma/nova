//Copyright (c) 2012 Nova Roma. All rights reserved.

package log

import (
	"io"
)

type Logger struct {
	Name   string
	Output io.Writer
}

var loggerCache map[string]*Logger

func CreateLogger(name string, out io.Writer) *Logger {
	l := &Logger{
		Name:   name,
		Output: out,
	}

	if loggerCache == nil {
		loggerCache = make(map[string]*Logger)
	}
	loggerCache[name] = l
	return l
}
