package main

import (
	"time"

	"github.com/fatih/color"
)

type BarberShop struct {
	ShopCapacity    int
	HaircutDuration time.Duration
	NumberOfBurbers int
	BurbersDoneChan chan bool
	ClientsChan     chan string
	Open            bool
}

func (shop *BarberShop) addburber(barber string) {
	shop.NumberOfBurbers++

	go func() {
		isSleeping := false
		color.Yellow("%s goes to the waiting room to check for clients.", barber)

		for {
			if len(shop.ClientsChan) == 0 {
				color.Cyan("Nothing to do, so %s takes a nap.", barber)
				isSleeping = true
			}

			client, shopOpen := <-shop.ClientsChan // Note: The 2nd param says if the channel is closed/empty. Warning: Using shop here will be a race and we must avoid it.

			if shopOpen {
				if isSleeping {
					color.Yellow("%s wakes %s up.", client, barber)
					isSleeping = false
				}

				shop.cutHair(barber, client)
			} else {
				shop.sendBurberHome(barber)
				return // Note: This will close the current go routine
			}
		}
	}()
}

func (shop *BarberShop) cutHair(barber, client string) {
	color.Green("%s is cutting %s's hair.", barber, client)
	time.Sleep(shop.HaircutDuration)
	color.Green("%s is finished cutting %s's hair.", barber, client)
}

func (shop *BarberShop) sendBurberHome(barber string) {
	color.Cyan("%s is going home.", barber)
	shop.BurbersDoneChan <- true
}

func (shop *BarberShop) closeShopForDay() {
	color.Cyan("Closing the shop for the day.")
	close(shop.ClientsChan)
	shop.Open = false

	for i := 0; i < shop.NumberOfBurbers; i++ { // Note: Here waiting for done signals from each barber before closing the BurbersDoneChan channel
		<-shop.BurbersDoneChan
	}

	close(shop.BurbersDoneChan)

	color.Green("The shop is closed for the day and everyone has gone home.")
}
