package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

//	func printDetails(product *Product) {
//		fmt.Println("name", product.name, "category", product.category, "price", product.price)
//	}
func (product *Product) printDetails() {
	fmt.Println("name", product.name, "category", product.category, "price", product.price)
}

func main() {
	products := []*Product{
		{"Kayak", "Watersports", 275},
		{"Lifejacket", "Watersports", 48.95},
		{"Soccer Ball", "Soccer", 19.5},
	}

	for _, p := range products {
		p.printDetails()
	}
}
