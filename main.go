package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Supplier struct {
	name, city string
}

//	func printDetails(product *Product) {
//		fmt.Println("name", product.name, "category", product.category, "price", product.price)
//	}
func (product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.calcTax(0.2, 100))
}

func (product *Product) calcTax(rate, threshold float64) float64 {
	if product.price > threshold {
		return product.price + (product.price * rate)
	}
	return product.price
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
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

	suppliers := []*Supplier{{"Acme Co", "New York City"}, {"BoatCo", "Chicago"}}
	for _, s := range suppliers {
		s.printDetails()
	}

	// kayak := Product{"Kayak", "Watersports", 275}  OR
	kayak := &Product{"Kayak", "Watersports", 275}
	kayak.printDetails()
}
