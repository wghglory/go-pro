package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Value:", rand.Int())

	// const price float32 = 275.00
	// const tax float32 = 27.50
	// const quantity = 2

	const price, tax float32 = 275.00, 27.50
	const quantity, inStock = 2, true
	fmt.Println("Total:", 2*quantity*(price+tax))
	fmt.Println("In stock:", inStock)
}
