package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	envFile, _ := godotenv.Read(".env")
	fmt.Println(envFile["PORT"])
}
