package main

import (
	"fmt"
	// "time"
)

func CalcStoreTotal(data ProductData) {
	var storeTotal float64
	var channel chan float64 = make(chan float64) // channel carries float64 values

	for category, group := range data {
		// storeTotal += group.SubTotalPrice(category)
		go group.SubTotalPrice(category, channel)
	}
	// one category per channel
	for i := 0; i < len(data); i++ {
		// receive a result using a channel
		// Receiving from a channel is a blocking operation, meaning that execution will not continue until a value has been received, which means I no longer have to prevent the program from terminating
		storeTotal += <-channel
		fmt.Println("-- channel read complete")
	}
	fmt.Println("Total:", ToCurrency(storeTotal))
}

func (group ProductGroup) SubTotalPrice(category string, resultChannel chan float64) {
	var total float64
	for _, p := range group {
		// fmt.Println(category, "product:", p.Name)
		total += p.Price
		// time.Sleep(time.Millisecond * 100)
	}
	fmt.Println(category, "channel sending", ToCurrency(total))
	// send value thru resultChannel
	resultChannel <- total
	fmt.Println(category, "channel send complete")
}
