package main

import (
	"fmt"
	"hash/maphash"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 500
var cutDuration = 1000 * time.Millisecond
var timeOpen = 10 * time.Second
var r = rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))

func main() {
	color.Cyan("The Sleeping Burber")

	clientChan := make(chan string, seatingCapacity)
	doneChan := make(chan bool)

	shop := BarberShop{
		ShopCapacity:    seatingCapacity,
		HaircutDuration: cutDuration,
		NumberOfBurbers: 0,
		BurbersDoneChan: doneChan,
		ClientsChan:     clientChan,
		Open:            true,
	}
	color.Green("Shop is open.")

	shop.addburber("Frank")
	shop.addburber("Susi")

	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen) // Note: Keep the current go routine alive open for the specified time
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	i := 0

	go func() {
		for {
			randomMilleSeconds := r.Int() % (2 * arrivalRate)

			select {
			case <-shopClosing:
				return
			case <-time.After(time.Millisecond * time.Duration(randomMilleSeconds)):
				shop.addClient(fmt.Sprintf("Client-%d", i))
				i++
			}
		}
	}()

	// time.Sleep(time.Second * 5)
	<-closed
}
