package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const HUNGER = 3

var philosophers = []string{"Plato", "Socrates", "Aristotle", "Pascal", "Locke"}
var wg sync.WaitGroup
var sleepTime = time.Second * 1
var eatTime = time.Second * 2
var thinkTime = time.Second * 1
var orderFinished []string
var orderMutex sync.Mutex

func diningProblem(philosopher string, leftFork, rightFork *sync.Mutex) {
	defer wg.Done()

	fmt.Println(philosopher, "is seated.")
	time.Sleep(sleepTime)

	for i := HUNGER; i > 0; i-- {
		fmt.Println(philosopher, "is hungry.")
		time.Sleep(sleepTime)

		leftFork.Lock()
		fmt.Printf("\t%s picked up the fork to his left.\n", philosopher)

		rightFork.Lock()
		fmt.Printf("\t%s picked up the fork to his right.\n", philosopher)

		fmt.Println(philosopher, "has both forks and is eating.")
		time.Sleep(eatTime)

		fmt.Println(philosopher, "is thinking!")
		time.Sleep(thinkTime)

		rightFork.Unlock()
		fmt.Printf("\t%s put down the fork on his right\n", philosopher)
		leftFork.Unlock()
		fmt.Printf("\t%s put down the fork on his left\n", philosopher)
		time.Sleep(sleepTime)
	}

	fmt.Println(philosopher, "is full.")
	time.Sleep(sleepTime)

	fmt.Println(philosopher, "has left the table.")

	orderMutex.Lock()
	orderFinished = append(orderFinished, philosopher)
	orderMutex.Unlock()
}

func main() {
	fmt.Println("The DInging Philosophers")

	wg.Add(len(philosophers))
	forkLeft := &sync.Mutex{}

	for i := 0; i < len(philosophers); i++ {
		forkRight := &sync.Mutex{}
		go diningProblem(philosophers[i], forkLeft, forkRight)
		forkLeft = forkRight
	}
	wg.Wait()

	fmt.Println("The table is empty.")

	fmt.Printf("Order finished: %s\n", strings.Join(orderFinished, ","))
}
