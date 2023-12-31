package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type Income struct {
	Source string
	Amount int
}

func main() {
	var bankBalance int
	var balance sync.Mutex

	fmt.Printf("Account balance: $%d.00", bankBalance)
	fmt.Println()

	incomes := []Income{
		{Source: "Main Job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part Time Job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}

	wg.Add(len(incomes))

	for i, income := range incomes {
		go func(i int, income Income) {
			defer wg.Done()

			// balance.Lock() // Note: For running serially
			for week := 0; week < 52; week++ {
				balance.Lock()
				temp := bankBalance
				temp += income.Amount
				bankBalance = temp
				balance.Unlock()

				fmt.Printf("On week - %d, you earned $%d.00. from %s. Current Balance: %d\n", week, income.Amount, income.Source, bankBalance)
			}
			// balance.Unlock()
		}(i, income)
	}

	wg.Wait()

	fmt.Printf("Final Bal: $%d.00\n", bankBalance)
}
