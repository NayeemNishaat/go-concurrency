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
			Wg:            &sync.WaitGroup{},
			ErrorChan:     make(chan error),
			ErrorChanDone: make(chan bool),
			Mailer: &lib.Mail{
				MailerChan: make(chan lib.Message),
				ErrorChan:  make(chan error),
				DoneChan:   make(chan bool),
			},
			Models: model.Models{User: &model.TestUser{}, Plan: &model.TestPlan{}},
		},
	}

	os.Exit(m.Run())
}

// Note: Logger
// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// infoLog.Printf("%s", "str")
