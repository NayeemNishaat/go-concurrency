package main

import (
	"fmt"
	"hash/maphash"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

const numberOfPizzas = 10

var pizzasMade, pizzasFailed, total int

type Producer struct {
	data chan PizzaOrder
	quit chan chan error // Note: quite channel's type is error channel
}

type PizzaOrder struct {
	pizzaNumber int
	message     string
	success     bool
}

func (p *Producer) Close() error {
	ch := make(chan error)
	p.quit <- ch
	return <-ch
}

func makePizza(pizzaNum int) *PizzaOrder {
	pizzaNum++

	if pizzaNum <= 10 {
		delay := rand.Intn(5) + 1
		fmt.Printf("Received order #%d.\n", pizzaNum)

		rnd := rand.Intn(12) + 1
		msg := ""
		success := false

		if rnd < 5 {
			pizzasFailed++
		} else {
			pizzasMade++
		}

		total++

		fmt.Printf("Making pizza #%d. It will take %d seconds.\n", pizzaNum, delay)

		time.Sleep(time.Duration(delay) * time.Second)

		if rnd <= 2 {
			msg = fmt.Sprintf("Ran out of ingredients for pizza #%d.\n", pizzaNum)
		} else if rnd <= 4 {
			msg = fmt.Sprintf("Cook unavailable for pizza #%d.\n", pizzaNum)
		} else {
			success = true
			msg = fmt.Sprintf("Pizza order #%d is ready.\n", pizzaNum)
		}

		p := PizzaOrder{
			pizzaNumber: pizzaNum,
			message:     msg,
			success:     success,
		}

		return &p
	}

	return &PizzaOrder{
		pizzaNumber: pizzaNum,
	}
}

func pizzaHut(pizzaMaker *Producer) {
	var i = 0

	for {
		currentPizza := makePizza(i)

		if currentPizza != nil {
			i = currentPizza.pizzaNumber
			select {
			case pizzaMaker.data <- *currentPizza: // Note: If currentPizza send it to pizzaMaker

			case quitChan := <-pizzaMaker.quit: // Note: If pizzaMaker.quit assign it to quitChan
				close(pizzaMaker.data)
				close(quitChan)
				return
			}
		}
	}
}

func main() {
	r := rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))
	fmt.Println(r.Int())

	color.Cyan("The shop is open.")
	color.Red("------------------")

	pizzaJob := &Producer{
		data: make(chan PizzaOrder),
		quit: make(chan chan error),
	}

	go pizzaHut(pizzaJob)

	for i := range pizzaJob.data {
		if i.pizzaNumber <= numberOfPizzas {
			if i.success {
				color.Green(i.message)
				color.Green("Order #%d is out for delivery!", i.pizzaNumber)
			} else {
				color.Red(i.message)
				color.Red("The customer is very mad!")
			}
		} else {
			color.Red("------------------")
			color.Cyan("Done making with pizzas!")

			err := pizzaJob.Close()

			if err != nil {
				color.Red("Error closing channel", err)
			}
		}
	}

	color.Yellow("We made %d pizzas, but failed to make %d with %d attempts in total.", pizzasMade, pizzasFailed, total)

	switch {
	case pizzasFailed > 9:
		color.Red("It was an awful day.")
	case pizzasFailed > 6:
		color.Red("It was not a very good day.")
	case pizzasFailed > 3:
		color.Red("It was an okay day.")
	case pizzasFailed >= 2:
		color.Red("It was a good day.")
	default:
		color.Green("It was a perfect day!")
	}
}
