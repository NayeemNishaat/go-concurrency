package main

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")

	os.Exit(exitVal)
}

func TestUpdateMessage(t *testing.T) {
	myMsg := "my text"

	wg.Add(1)
	go updateMessage(myMsg)
	wg.Wait()

	if msg != myMsg {
		t.Errorf("updateMessage Failed")
	}
}

func TestPrintMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	printMessage()

	_ = w.Close()

	data, _ := io.ReadAll(r)
	dataStr := string(data)

	os.Stdout = stdOut

	if !strings.Contains(dataStr, msg) {
		t.Errorf("printMessage Failed")
	}

}

func TestPrintUpdateMessage(t *testing.T) {
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

func TestMainFunc(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()

	data, _ := io.ReadAll(r)
	dataStr := string(data)

	os.Stdout = stdOut

	if !strings.Contains(dataStr, "Hello, universe!") {
		t.Errorf("'Hello, universe!' not found!")
	}

	if !strings.Contains(dataStr, "Hello, cosmos!") {
		t.Errorf("'Hello, cosmos!' not found!")
	}

	if !strings.Contains(dataStr, "Hello, world!") {
		t.Errorf("'Hello, world!' not found!")
	}
}
