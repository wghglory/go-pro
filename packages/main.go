package main

import (
	"fmt"
	currencyFmt "packages/fmt"

	// . "packages/fmt"
	"packages/store"
	"packages/store/cart"
)

func main() {
	// product := store.Product{
	// 	Name:     "Kayak",
	// 	Category: "Watersports",
	// }

	product := store.NewProduct("Kayak", "Watersports", 279)

	cart := cart.Cart{
		CustomerName: "Derek",
		Products:     []store.Product{*product},
	}

	fmt.Println("Customer Name:", cart.CustomerName)
	fmt.Println("Product Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))
	// fmt.Println("Price:", ToCurrency(product.Price()))
}
