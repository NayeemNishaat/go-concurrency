package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMainFunc(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	w.Close()

	res, _ := io.ReadAll(r)
	out := string(res)

	os.Stdout = stdOut

	if !strings.Contains(out, "34320") {
		t.Errorf("Wrong Balance")
	}
}
