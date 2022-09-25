package main

import "fmt"

type Product struct {
	name, category string
	price          float64
}

type Supplier struct {
	name, city string
}

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
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

func getProducts() []Product {
	return []Product{{"Kayak", "Watersports", 275}, {"Lifejacket", "Watersports", 48.95}, {"Soccer Ball", "Soccer", 19.50}}
}

func (supplier *Supplier) printDetails() {
	fmt.Println("Supplier:", supplier.name, "City:", supplier.city)
}

func main() {
	// products := []*Product{
	// products := ProductList{
	// 	{"Kayak", "Watersports", 275},
	// 	{"Lifejacket", "Watersports", 48.95},
	// 	{"Soccer Ball", "Soccer", 19.5},
	// }
	products := ProductList(getProducts())
	for _, p := range products {
		p.printDetails()
	}
	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category: ", category, "Total:", total)
	}

	suppliers := []*Supplier{{"Acme Co", "New York City"}, {"BoatCo", "Chicago"}}
	for _, s := range suppliers {
		s.printDetails()
	}

	// kayak := Product{"Kayak", "Watersports", 275}  OR
	kayak := &Product{"Kayak", "Watersports", 275}
	kayak.printDetails()
}
