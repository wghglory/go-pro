package main

import (
	"fmt"
	"time"
	// "time"
)

// The for loop will continue to receive values until the channel is closed. (You can use a for...range loop on a channel that isn’t closed, in which case the loop will never exit.)
func receiveDispatches(channel <-chan DispatchNotification) {
	for details := range channel {
		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	}
	fmt.Println("Channel has been closed")
}

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

	// ---------- receive unknown number of channel messages, need check channel open/close ---------
	dispatchChannel := make(chan DispatchNotification, 100)
	// var sendOnlyChannel chan<- DispatchNotification = dispatchChannel
	// var receiveOnlyChannel <-chan DispatchNotification = dispatchChannel

	// go DispatchOrders(sendOnlyChannel)
	go DispatchOrders(chan<- DispatchNotification(dispatchChannel)) // explicit conversion
	// for {
	// 	if details, open := <-dispatchChannel; open {
	// 		fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
	// 	} else {
	// 		fmt.Println("Channel has been closed")
	// 		break
	// 	}
	// }
	// receiveDispatches(receiveOnlyChannel)
	// receiveDispatches((<-chan DispatchNotification)(dispatchChannel)) // explicit conversion

	for {
		select {
		case details, ok := <-dispatchChannel:
			if ok {
				fmt.Println("Dispatch to", details.Customer, ":", details.Quantity, "x", details.Product.Name)
			} else {
				fmt.Println("Channel has been closed")
				goto alldone
			}
		default:
			fmt.Println("-- No message ready to be received")
			time.Sleep(time.Millisecond * 500)
		}
	}
alldone:
	fmt.Println("All values received")

}
