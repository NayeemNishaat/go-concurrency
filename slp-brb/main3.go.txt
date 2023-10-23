package main

import (
	"fmt"
	"time"
)

func listen(ch chan int) {
	for { // Important: If infinite for loop is not used then this go routine dies after it reaches to the end of function. An infinite for loop makes the go routine to loop again, wait for input from channel.
		i := <-ch
		fmt.Println("Got", i)

		time.Sleep(1 * time.Second)
	}
}

func main() {
	ch := make(chan int, 10) // Note: Now channel capacity is 10 by default 1.

	go listen(ch)

	for i := 0; i < 100; i++ {
		fmt.Println("Sending", i)
		ch <- i
		fmt.Println("Sent", i)
	}

	fmt.Println("Done")
	close(ch)
}
