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

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var wg sync.WaitGroup
var db *pgxpool.Pool

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db = lib.InitDB()

	cfg := lib.Config{
		Wg:            &wg,
		Models:        model.New(db),
		ErrorChan:     make(chan error),
		ErrorChanDone: make(chan bool),
	}
	cfg.Mailer = cfg.CreateMailer()

	lib.InitCfg(&cfg)
	controller.InitCfg(&cfg)
	go cfg.ListenForMail()

	go cfg.ListenForErrors()

	go gracefulShutdown()
	server()
}

func server() {
	mux := http.NewServeMux()

	// Note: Serve Favicon
	mux.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./public/img/favicon.ico")
	})

	// Note: Serving Public Dir
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))) // Important: If a trailing / is given then all the request with prefix "public" will be handled by this handler.

	route.InitRouter(mux)

	fmt.Printf("Serving on http://localhost:%s\n", os.Getenv("PORT"))
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
	log.Println("Performing Cleanup")

	defer db.Close()
	wg.Wait() // Note: Block until wg is empty -> 0
	lib.Cfg.Mailer.DoneChan <- true
	lib.Cfg.ErrorChanDone <- true

	log.Println("Closing channels and shutting down app.")
	close(lib.Cfg.Mailer.DoneChan)
	close(lib.Cfg.Mailer.ErrorChan)
	close(lib.Cfg.Mailer.MailerChan)
	close(lib.Cfg.ErrorChan)
	close(lib.Cfg.ErrorChanDone)
}

// /opt/homebrew/opt/postgresql@16/bin/postgres -D /opt/homebrew/var/postgresql@16
// go list ./...
// go list ...
