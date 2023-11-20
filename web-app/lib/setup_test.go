package lib

import (
	"encoding/gob"
	"os"
	"sync"
	"testing"
	"web/model"
)

var TestConfig Config

func TestMain(m *testing.M) {
	gob.Register(model.User{})

	TestConfig = Config{
		Wg:            &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		Mailer:        &Mail{},
		Models:        model.Models{},
	}

	os.Exit(m.Run())
}

// Note: Logger
// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// infoLog.Printf("%s", "str")
