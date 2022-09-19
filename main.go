package main

import (
	"fmt"
	"sort"
)

// "math/rand"

func main() {
	// fmt.Println("Value:", rand.Int())

	// -----------------------------const-----------------------------------
	// const price float32 = 275.00
	// const tax float32 = 27.50
	// const quantity = 2

	// const price, tax float32 = 275.00, 27.50
	// const quantity, inStock = 2, true
	// fmt.Println("Total:", 2*quantity*(price+tax))
	// fmt.Println("In stock:", inStock)

	// -----------------------------var-----------------------------------
	// var price float32 = 275.00
	// var tax float32 = 27.50
	// fmt.Println(price + tax)
	// price = 300
	// fmt.Println(price + tax)

	// var price = 275.00
	// var price2 = price
	// fmt.Println(price)
	// fmt.Println(price2)

	// var price float32
	// fmt.Println(price) // 0
	// price = 275.00
	// fmt.Println(price)

	// var price, tax = 275.00, 27.50
	// fmt.Println(price + tax)

	// var price, tax float64
	// price = 275.00
	// tax = 27.50
	// fmt.Println(price + tax)

	// price, tax, inStock := 275.00, 27.50, true
	// fmt.Println(price + tax)
	// fmt.Println(inStock)

	// ------------------------------pointers----------------------------------
	// first := 100
	// second := first
	// first++
	// // 2 memory addresses, 2 copies
	// fmt.Println("First:", first)   // 101
	// fmt.Println("Second:", second) // 100

	/*first := 100
	var second *int = &first // first's address
	first++
	fmt.Println("First:", first)                   // 101
	fmt.Println("Second (First address):", second) // 0x14000018110
	fmt.Println("Second value:", *second)          // 101
	*second++
	fmt.Println("First:", first)    //102
	fmt.Println("Second:", *second) //102
	var pointer *int = second
	*pointer++
	fmt.Println("First:", first)    //103
	fmt.Println("Second:", *second) //103
	var nilPointer *int
	fmt.Println(nilPointer) // <nil>
	nilPointer = &first
	fmt.Println(nilPointer) // 0x14000018110
	*/

	// names := [3]string{"Alice", "Charlie", "Bob"}
	// secondName := names[1]  // Charlie, in a new address
	// fmt.Println(secondName) // Charlie
	// sort.Strings(names[:])  // names array: Alice, Bob, Charlie
	// fmt.Println(secondName) // Charlie!!! secondName is in a new address, which doesn't have any relationship with names array

	names := [3]string{"Alice", "Charlie", "Bob"}
	secondPosition := &names[1]  // secondPosition is name array pointer, pointing to index 1
	fmt.Println(*secondPosition) // Charlie
	sort.Strings(names[:])       // names array: Alice, Bob, Charlie
	fmt.Println(*secondPosition) // Bob!!! pointing to index 1, which is Bob now
}
