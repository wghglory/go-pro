package main

import "fmt"

func printPrice(product string, price, taxRate float64) {
	taxAmount := taxRate * price
	fmt.Println(product, "Price:", price, "Tax:", taxAmount)
}

// Variadic Parameters
func printSuppliers(product string, suppliers ...string) {
	if len(suppliers) == 0 {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func swapValues(first, second int) (int, int) {
	return second, first
}

func calcTax(price float64) (float64, bool) {
	if price > 100 {
		return price * 0.2, true
	}
	return 0, false
}

func calcTotalPrice(products map[string]float64) (count int, total float64) {
	// Function started
	// Function about to return
	// Second defer call
	// First defer call
	// Total: 323.95

	// defer keyword is to call functions that release resources, such as closing open files or HTTP connections. The defer keyword lets you group the statements that create, use, and release the resource together.
	// Immediately before the function returns, Go will perform the calls scheduled with the defer keyword in the order in which they were defined
	fmt.Println("Function started")
	defer fmt.Println("First defer call")
	count = len(products)
	for _, price := range products {
		total += price
	}
	defer fmt.Println("Second defer call")
	fmt.Println("Function about to return")
	return
}

func main() {
	// printPrice("Kayak", 275, 0.2)
	// printPrice("Lifejacket", 48.95, 0.2)
	// printPrice("Soccer Ball", 19.50, 0.15)

	// printSuppliers("Kayak", "Acme Kayaks", "Bob's Boats", "Crazy Canoes")
	// printSuppliers("Soccer Ball")

	names := []string{"Acme Kayaks", "Bob's Boats", "Crazy Canoes"}
	printSuppliers("Kayak", names...)

	products := map[string]float64{"Kayak": 275, "Lifejacket": 48.95}
	for product, price := range products {
		if taxAmount, taxDue := calcTax(price); taxDue {
			fmt.Println("Product: ", product, "Tax:", taxAmount)
		} else {
			fmt.Println("Product: ", product, "No tax due")
		}
	}

	_, total := calcTotalPrice(products)
	fmt.Println("Total:", total)
}
