package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) { // Note: Receive/Send only ping channel
	for {
		s := <-ping

		// pong <- <-ping
		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	// time.Sleep(10 * time.Second) // Note: Now the go routine will last 10s cause after 10s main go routine will end for

	for {
		fmt.Print("> ")
		var userInput string

		fmt.Scanln(&userInput)

		if userInput == strings.ToLower("q") {
			break
		}

		ping <- userInput
		response := <-pong
		fmt.Println(response)
	}

	close(ping)
	close(pong)
}
