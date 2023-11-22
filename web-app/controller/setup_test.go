package controller

import (
	"log"
	"os"
	"sync"
	"testing"
	"web/lib"
	"web/model"

	"github.com/joho/godotenv"
)

var TestConfig Config

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// lib.InitTestConfig()
	TestConfig = Config{
		&lib.Config{
			Wg:            &sync.WaitGroup{}, // Note: Always create and initialize wg and chan when testing
			ErrorChan:     make(chan error),
			ErrorChanDone: make(chan bool),
			Mailer: &lib.Mail{
				MailerChan: make(chan lib.Message, 100), // Important: Must be a buffered chan else it will wait/hang until the message sent to the channel is received (channel becomes empty).
				ErrorChan:  make(chan error),
				DoneChan:   make(chan bool),
			},
			Models: model.Models{User: &model.TestUser{}, Plan: &model.TestPlan{}},
		},
	}

	TestConfig.Mailer.Wait = TestConfig.Wg

	go func() {
		select {
		case <-TestConfig.ErrorChan: // Important: Must need to listen to the created channels
		case <-TestConfig.ErrorChanDone:
			return
		}
	}()

	go func() {
		for {
			select {
			case <-TestConfig.Mailer.MailerChan:
			case <-TestConfig.Mailer.ErrorChan:
			case <-TestConfig.Mailer.DoneChan:
				return
			}
		}
	}()

	os.Exit(m.Run())
}

// Note: Logger
// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// infoLog.Printf("%s", "str")

// go test .
// go test . -v
// go test ./...
// go test -coverprofile=coverage.out
// go test -coverprofile=coverage.out
// go tool cover -html=coverage.out
