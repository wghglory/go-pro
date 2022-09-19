package main

import (
	"fmt"
	"strconv"
)

func main() {
	//-----------------------------if-----------------------------------
	kayakPrice := 275.00
	if kayakPrice > 500 {
		scopedVar := 500
		fmt.Println("Price is greater than", scopedVar)
	} else if kayakPrice < 100 {
		scopedVar := "Price is less than 100"
		fmt.Println(scopedVar)
	} else {
		scopedVar := false
		fmt.Println("Matched: ", scopedVar) // false, executed.
	}

	priceString := "275"
	if price, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price is", price) // 275
	} else {
		fmt.Println("error:", err)
	}

	//-----------------------------loop-----------------------------------
	counter := 0
	for {
		fmt.Println("Counter1:", counter) // 0 1 2 3
		counter++
		if counter > 3 {
			break
		}
	}

	for counter <= 6 {
		fmt.Println("Counter2:", counter) // 4 5 6
		counter++
	}

	// scoped local variable counter
	for counter := 0; counter <= 8; counter++ {
		if counter == 1 {
			continue
		}
		fmt.Println("Counter3:", counter) // 0 2 3 4 5 6 7 8
	}

	//---------------------------------Enumerating Sequences--------------------------------
	product := "Kayak"
	for index, character := range product {
		fmt.Println("Index:", index, "ASCII:", character, "Character:", string(character))
		// Index: 0 ASCII: 75 Character: K
		// Index: 1 ASCII: 97 Character: a
		// Index: 2 ASCII: 121 Character: y
		// Index: 3 ASCII: 97 Character: a
		// Index: 4 ASCII: 107 Character: k
	}

	//-----------------------------Enumerating Built-in Data Structures-----------------------------------
	products := []string{"Kayak", "Lifejacket", "Soccer Ball"}
	for index, element := range products {
		fmt.Println("Index:", index, "Element", element)
		// Index: 0 Element Kayak
		// Index: 1 Element Lifejacket
		// Index: 2 Element Soccer Ball
	}

	//-----------------------------Switch-------------------------------
	for index, character := range product {
		switch character {
		case 'K', 'k':
			// fmt.Println("K or k at position", index)
			if character == 'k' {
				fmt.Println("Lowercase k at position", index)
				// break means if k, stop executing the below statements in this case, Uppercase K at position won't be executed
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		default:
			fmt.Println("Character", string(character), "at position", index)
		}
	}

	for counter := 0; counter < 10; counter++ {
		switch {
		case counter == 0:
			fmt.Println("Zero value")
		case counter < 3:
			fmt.Println(counter, "is < 3")
		case counter >= 3 && counter < 7:
			fmt.Println(counter, "is >= 3 && < 7")
		default:
			fmt.Println(counter, "is >= 7")
		}
	}

	counter2 := 0
	// label
target:
	fmt.Println("Counter", counter2)
	counter2++
	if counter2 < 5 {
		goto target
	}
}
