package main

import "fmt"

type Supplier struct {
	name, city string
}

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Account struct {
	accountNumber int
	expenses      []Expense
}

type Person struct {
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

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price)
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
	case Person:
		fmt.Println("Person:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Person Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default:", value)
	}
}
func processItems(items ...interface{}) {
	for _, item := range items {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
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

	expenses := []Expense{
		&Product{"Kayak", "Watersports", 275},
		Service{"Boat Cover", 12, 89.50},
	}
	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(expenses))

	account := Account{accountNumber: 12345, expenses: []Expense{&Product{"Kayak", "Watersports", 275}, Service{"Boat Cover", 12, 89.50}}}
	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))

	expenses2 := []Expense{Service{"Boat Cover", 12, 89.50}, Service{"Paddle Protect", 12, 8}}
	for _, expense := range expenses2 {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee*float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}

	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee*float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}

	// empty interface
	data := []interface{}{
		expenses[0],
		Product{"Lifejacket", "Watersports", 48.95},
		Service{"Boat Cover", 12, 89.50},
		Person{"Alice", "London"},
		&Person{"Bob", "New York"},
		"This is a string",
		100,
		true,
	}

	for _, item := range data {
		processItem(item)
	}

	processItems(data...)
}
