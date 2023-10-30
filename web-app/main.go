package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"web/lib"
	"web/route"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db := lib.InitDB()
	db.Ping(context.Background())

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

// /opt/homebrew/opt/postgresql@16/bin/postgres -D /opt/homebrew/var/postgresql@16
