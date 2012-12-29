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

	if logger.Output != buf {
		t.Error("Returned logger should have the correct writer.")
	}

	if len(loggerCache) != 1 {
		t.Errorf("Expected the logger cache to contain %d loggers but had %d.", 1, len(loggerCache))
	}
}
