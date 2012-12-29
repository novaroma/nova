//Copyright (c) 2012 Nova Roma. All rights reserved.

package log

import (
	"bytes"
	"testing"
)

func TestCreateLogger(t *testing.T) {
	buf := bytes.NewBufferString("")

	logger := CreateLogger("test", buf)

	if logger == nil {
		t.Error("Expected the returned logger to not be nil.")
	}

	if logger.Name != "test" {
		t.Errorf("Expected the returned logger to have name '%s' but was '%s'.", "test", logger.Name)
	}

	if len(loggerCache) != 1 {
		t.Errorf("Expected the logger cache to contain %d loggers but had %d.", 1, len(loggerCache))
	}

	if len(logger.levels) != 4 {
		t.Errorf("Expected the logger to have %d log levels but had %d.", 4, len(logger.levels))
	}

	for _, level := range logger.levels {
		if level.output != buf {
			t.Error("Expected the log levels to have the correct output writer.")
		}
	}
}
