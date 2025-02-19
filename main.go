package main

import (
	"fmt"
	"sync"

	"github.com/AGX18/CryptoMasters/api"
)

func main() {

	currencies := []string{"BTC", "BCH", "ETH"}

	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1);
		go func () {
			getCurrencyData(currency)
			wg.Done()
		}()

	}

	wg.Wait()
}


func getCurrencyData(currency string) {
	rate, err := api.GetRate(currency)
	if err == nil {
		fmt.Printf("The rate for %v is %.2f\n", rate.Currency, rate.Price)
	}
}