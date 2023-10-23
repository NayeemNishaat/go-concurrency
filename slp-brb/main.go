package main

import (
	"hash/maphash"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

var seatingCapacity = 10
var arrivalRate = 100
var cutDuration = 1000 * time.Microsecond
var timeOpen = 10 * time.Second
var r = rand.New(rand.NewSource(int64(new(maphash.Hash).Sum64())))

func main() {
	color.Yellow("The Sleeping Burber")

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

	shopClosing := make(chan bool)
	closed := make(chan bool)
	go func() {
		<-time.After(timeOpen) // Note: Keep the channel open for the specified time
		shopClosing <- true
		shop.closeShopForDay()
		closed <- true
	}()

	// time.Sleep(time.Second * 5)
}
