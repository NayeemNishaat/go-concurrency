package main

import (
	"encoding/gob"
	"os"
	"sync"
	"testing"
	"web/lib"
	"web/model"
)

var testConfig lib.Config

func TestMain(m *testing.M) {
	gob.Register(model.User{})

	testConfig = lib.Config{
		Wg:            &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		Mailer:        &lib.Mail{},
	}

	os.Exit(m.Run())
}

// Note: Logger
// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// infoLog.Printf("%s", "str")
