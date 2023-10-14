package main

import (
	"fmt"
	"sync"
)

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done() // Note: Defer means don't go forward until this line is completed

	fmt.Println(s)
}

// Note: Main function itself is a Go Routine, it runs on a very light weight thread
func main() {
	// Chapter: Weight Group
	// go printSomething("Print it!")

	var wg sync.WaitGroup

	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
	}

	wg.Add(len(words))

	for i, x := range words {
		go printSomething(fmt.Sprintf("%d: %s", i, x), &wg) // Warning: Pass Weight Group always by reference, don't copy them.
	}

	wg.Wait()

	// time.Sleep(1 * time.Second) // Important: Without this the program will exit so fast that the go routine spowned above won't get time to be executed and will just die. We are basically waiting a second to give enough time to the go routine to finish its job.

	wg.Add(1) // Note: Preventing negative wg. It's in 0 at this point.
	printSomething("Print it too!", &wg)
}
