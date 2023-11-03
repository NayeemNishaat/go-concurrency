package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"web/controller"
	"web/lib"
	"web/model"
	"web/route"

	"github.com/joho/godotenv"
)

var wg sync.WaitGroup

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := lib.InitDB()
	// db.Ping(context.Background())

	cfg := controller.Config{
		Wg:     &wg,
		Models: model.New(db),
	}
	controller.InitCfg(&cfg)

	go gracefulShutdown()
	server()
}

func server() {
	mux := http.NewServeMux()
	route.InitRouter(mux)

	err := http.ListenAndServe(fmt.Sprintf(":%s", os.Getenv("PORT")), mux)

	if err != nil {
		log.Panic(err)
	}
}

func gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit // Note: Waiting for interrupt signal

	shutdown()
	os.Exit(0)
}

func shutdown() {
	log.Println("Perform Cleanup")

	wg.Wait() // Note: Block until wg is empty -> 0

	log.Println("Cleaning up and shutting down app.")
}

// /opt/homebrew/opt/postgresql@16/bin/postgres -D /opt/homebrew/var/postgresql@16
