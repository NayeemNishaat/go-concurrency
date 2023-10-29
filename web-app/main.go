package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var env map[string]string

func main() {
	env, _ = godotenv.Read(".env")
	db := initDB()
	db.Ping(context.Background())

	server()
}

func initDB() *pgx.Conn {
	conn := connectToDB()

	if conn == nil {
		log.Panic("Can't connect to the DB.")
	}

	return conn
}

func connectToDB() *pgx.Conn {
	count := 0

	dsn := env["DSN"]

	for {
		conn, err := openDB(dsn)

		if err != nil {
			log.Println("DB is not yet ready.")
		} else {
			log.Print("Connected to DB.")
			return conn
		}

		if count > 5 {
			return nil
		}

		log.Print("Backing off for 1 second.")
		time.Sleep(1 * time.Second)

		count++
		continue
	}
}

func openDB(dsn string) (*pgx.Conn, error) {
	// db, err := sql.Open("pgx", dsn)
	db, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())

	if err != nil {
		return nil, err
	}

	defer db.Close(context.Background())

	return db, nil
}

func server() {
	r := &Router{}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", env["PORT"]),
		Handler: r,
	}

	err := srv.ListenAndServe()

	if err != nil {
		log.Panic(err)
	}
}

// /opt/homebrew/opt/postgresql@16/bin/postgres -D /opt/homebrew/var/postgresql@16
