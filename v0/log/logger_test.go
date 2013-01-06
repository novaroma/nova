//Copyright (c) 2012 Nova Roma. All rights reserved.

package log

import (
	"bytes"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"
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
	b.StopTimer()
	name := randomString()
	GetLogger(name)
	defer func() {
		b.StopTimer()
		tearDown()
		b.StartTimer()
	}()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		GetLogger(name)
	}
}

func TestLog(t *testing.T) {
	buf := bytes.NewBufferString("")
	defer tearDown()

	logger := CreateLogger("test", buf)
	logger.Log(LogLevelDebug, "ASDF", 1234, "pqf", "afg")

	out := buf.String()
	if !strings.Contains(out, PrefixLogLevelDebug) {
		t.Errorf("Expected the logger output '%s' to contain the prefix '%s'.", out, PrefixLogLevelDebug)
	}

	formatContent := "ASDF1234pqfafg"
	if !strings.Contains(out, formatContent) {
		t.Errorf("Expected the logger output '%s' to contain the content '%s'.", out, formatContent)
	}
}

func TestLog_MissingLevel(t *testing.T) {
	buf := bytes.NewBufferString("")
	defer tearDown()

	logger := CreateLogger("test", buf)
	err := logger.Log(42, "ASDF", 1234, "pqf", "afg")

	if err == nil {
		t.Errorf("Expected an error to be returend but was nil.")
	}
}

func BenchmarkLog(b *testing.B) {
	b.StopTimer()
	logger := CreateLogger("test", ioutil.Discard)
	defer func() {
		b.StopTimer()
		tearDown()
		b.StartTimer()
	}()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Log(LogLevelDebug, "ASDLFKAJDLFKAJSDF")
	}
}

func BenchmarkStdLog(b *testing.B) {
	b.StopTimer()
	log.SetOutput(ioutil.Discard)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		log.Print("ASDLFKAJDLFKAJSDF")
	}
}

func TestLogf(t *testing.T) {
	buf := bytes.NewBufferString("")
	defer tearDown()

	logger := CreateLogger("test", buf)
	logger.Logf(LogLevelDebug, "This is a test with %d numbers, a '%s' string, and a complex %v", 4, "ASDFLKJASDLFKJADS", complex(-1, 1))

	actual := buf.String()
	if !strings.Contains(actual, PrefixLogLevelDebug) {
		t.Errorf("Expected the logger output '%s' to contain the prefix '%s'", actual, PrefixLogLevelDebug)
	}

	formatContent := "This is a test with 4 numbers, a 'ASDFLKJASDLFKJADS' string, and a complex (-1+1i)"
	if !strings.Contains(actual, formatContent) {
		t.Errorf("Expected the logger output '%s' to contain the content '%s'", actual, formatContent)
	}
}

func TestLogf_MissingLevel(t *testing.T) {
	buf := bytes.NewBufferString("")
	defer tearDown()

	logger := CreateLogger("test", buf)
	err := logger.Logf(42, "This is a test with %d numbers, a '%s' string, and a complex %v", 4, "ASDFLKJASDLFKJADS", complex(-1, 1))

	if err == nil {
		t.Errorf("Expected an error to be returend but was nil.")
	}
}

func BenchmarkLogf(b *testing.B) {
	b.StopTimer()
	logger := CreateLogger("test", ioutil.Discard)
	defer func() {
		b.StopTimer()
		tearDown()
		b.StartTimer()
	}()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Logf(LogLevelDebug, "ASDLFKAJDLFKAJSDF%d", 3)
	}
}

func BenchmarkStdLogf(b *testing.B) {
	b.StopTimer()
	log.SetOutput(ioutil.Discard)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		log.Printf("ASDLFKAJDLFKAJSDF%d", 3)
	}
}

func TestLogln(t *testing.T) {
	buf := bytes.NewBufferString("")
	defer tearDown()

	logger := CreateLogger("test", buf)
	logger.Logln(LogLevelDebug, "this", "is", 2, complex(-1, 1))

	actual := buf.String()
	if !strings.Contains(actual, PrefixLogLevelDebug) {
		t.Errorf("Expected the logger output '%s' to contain the prefix '%s'", actual, PrefixLogLevelDebug)
	}

	formatContent := "this is 2 (-1+1i)\n"
	if !strings.Contains(actual, formatContent) {
		t.Errorf("Expected the logger output '%s' to contain the content '%s'", actual, formatContent)
	}
}

func TestLogln_MissingLevel(t *testing.T) {
	buf := bytes.NewBufferString("")
	defer tearDown()

	logger := CreateLogger("test", buf)
	err := logger.Logln(42, "this", "is", 2, complex(-1, 1))

	if err == nil {
		t.Errorf("Expected an error to be returend but was nil.")
	}
}

func BenchmarkLogln(b *testing.B) {
	b.StopTimer()
	logger := CreateLogger("test", ioutil.Discard)
	defer func() {
		b.StopTimer()
		tearDown()
		b.StartTimer()
	}()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		logger.Logln(LogLevelDebug, "ASDLFKAJDLFKAJSDF", 3)
	}
}

func BenchmarkStdLogln(b *testing.B) {
	b.StopTimer()
	log.SetOutput(ioutil.Discard)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		log.Println("ASDLFKAJDLFKAJSDF", 3)
	}
}

func TestDisableLevel(t *testing.T) {
	buf := bytes.NewBufferString("")
	defer tearDown()

	logger := CreateLogger("test", buf)
	logger.DisableLevel(LogLevelDebug)

	logger.Debug("This is a test.")

	out := buf.String()

	if out != "" {
		t.Errorf("Expected the output to be '%s' but was '%s'.", "", out)
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
