package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	wg.Add(1)
	go updateMessage("This is my message!")
	wg.Wait()
	printMessage()

	_ = w.Close()

	data, _ := io.ReadAll(r)
	dataStr := string(data)

	os.Stdout = stdOut

	if msg != "This is my message!" {
		t.Errorf("updateMessage Failed")
	}

	if !strings.Contains(dataStr, "This is my message!") {
		t.Errorf("printMessage Failed")
	}

	if msg != strings.Trim(dataStr, "\n") {
		// Note: Multiple chars can be provided strings.Trim(dataStr, "\n -")
		t.Errorf("printMessage Failed")
	}
}
