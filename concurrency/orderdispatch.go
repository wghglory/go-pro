package main

import (
	"fmt"
	"math/rand"
	"time"
)

type DispatchNotification struct {
	Customer string
	*Product
	Quantity int
}

var Customers = []string{"Alice", "Bob", "Charlie", "Dora"}

// chan<- means channel can be used only to send.
// <-chan means to receive only
func DispatchOrders(channel chan<- DispatchNotification) {
	rand.Seed(time.Now().UTC().UnixNano())
	orderCount := rand.Intn(3) + 2
	fmt.Println("Order count:", orderCount)

	// There is no way to know in advance how many DispatchNotification values the DispatchOrders function will create, which presents a challenge when writing the code that receives from the channel.
	for i := 0; i < orderCount; i++ {
		channel <- DispatchNotification{
			Customer: Customers[rand.Intn(len(Customers)-1)],
			Quantity: rand.Intn(10),
			Product:  ProductList[rand.Intn(len(ProductList)-1)],
		}
	}
	close(channel)
}
