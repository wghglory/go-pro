package main

import (
	// "fmt"
	currencyFmt "packages/fmt"

	// . "packages/fmt"
	_ "packages/data"
	"packages/store"
	"packages/store/cart"

	"github.com/fatih/color"
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

	color.Green("Customer Name: " + cart.CustomerName)
	color.Cyan("Product Name: " + product.Name)
	color.Cyan("Category: " + product.Category)
	color.Cyan("Price: " + currencyFmt.ToCurrency(product.Price()))
	// fmt.Println("Price:", ToCurrency(product.Price()))
}
