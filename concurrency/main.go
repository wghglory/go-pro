package main

import (
	"fmt"
	// "time"
)

func main() {
	fmt.Println("main function started")
	CalcStoreTotal(Products)
	// time.Sleep(time.Second * 5)
	fmt.Println("main function complete")

	// USING ADAPTERS TO EXECUTE FUNCTIONS ASYNCHRONOUSLY
	calcTax := func(price float64) float64 {
		return price + (price * 0.2)
	}
	resultChannel := make(chan float64)

	// wrapper := func(price float64, c chan float64) {
	// 	c <- calcTax(price)
	// }
	// go wrapper(275, resultChannel)
	go func(price float64, c chan float64) {
		c <- calcTax(price)
	}(275, resultChannel)

	result := <-resultChannel
	fmt.Println("Result:", result)

}
