//Copyright (c) 2012 Nova Roma. All rights reserved.

package log

import (
	"bytes"
	"math/rand"
	"os"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix() >> 0x4f83bca)
}

func TestCreateLogger(t *testing.T) {
	defer tearDown()
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

func BenchmarkCreateLogger(b *testing.B) {
	b.StopTimer()
	buf := bytes.NewBufferString("")
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		CreateLogger("test", buf)

		b.StopTimer()
		tearDown()
		b.StartTimer()
	}
}

func TestGetLogger_ExistingLogger(t *testing.T) {
	defer tearDown()
	buf := bytes.NewBufferString("")
	expect := CreateLogger("test", buf)
	if expect == nil {
		t.FailNow()
	}

	actual := GetLogger("test")
	if actual != expect {
		t.Error("Expected the logger returned by GetLogger to be a reference to an existing logger.")
	}
}

func TestGetLogger_NewLogger(t *testing.T) {
	defer tearDown()
	actual := GetLogger("test")

	if actual == nil {
		t.Error("Expected the returned logger to not be nil.")
	}

	for _, level := range actual.levels {
		if level.output != os.Stdout {
			t.Errorf("Expected the outputs of the log levels to be '%p' but was '%p'.", os.Stdout, level.output)
		}
	}
}

func BenchmarkGetLogger_NewLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetLogger("test")

		b.StopTimer()
		tearDown()
		b.StartTimer()
	}
}

func BenchmarkGetLogger_LookupLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		name := randomString()
		b.StartTimer()

		GetLogger(name)
	}
}

func randomString() string {
	buf := bytes.NewBufferString("")
	numChars := rand.Intn(300)
	for i := 0; i < numChars; i++ {
		runeInt := rand.Int31()
		buf.WriteRune(rune(runeInt))
	}

	return buf.String()
}

func tearDown() {
	loggerCache = make(map[string]*Logger)
}
