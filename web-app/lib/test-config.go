package lib

import (
	"sync"
	"web/model"
)

var TestConfig Config

func InitTestConfig() {
	// gob.Register(model.User{})
	TestConfig = Config{
		Wg:            &sync.WaitGroup{},
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
		Mailer:        &Mail{},
		Models:        model.Models{},
	}
}
