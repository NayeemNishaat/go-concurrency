package controller

import (
	"log"
	"os"
	"testing"
	"web/lib"

	"github.com/joho/godotenv"
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	lib.InitTestConfig()

	os.Exit(m.Run())
}

// Note: Logger
// infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
// log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
// infoLog.Printf("%s", "str")
