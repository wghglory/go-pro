package main

import "fmt"

type Product struct {
	name, category string
	price          float64
	*Supplier
}

type Supplier struct {
	name, city string
}

// constructor
func newProduct(name, category string, price float64, supplier *Supplier) *Product {
	return &Product{name, category, price - 10, supplier}
}

// deep copy
func copyProduct(product *Product) Product {
	p := *product
	s := *product.Supplier
	p.Supplier = &s
	return p
}

func calcTax(product *Product) *Product {
	if product.price > 100 {
		product.price += product.price * 0.2
	}
	return product
}

func main() {
	type Item struct {
		name     string
		category string
		price    float64
		*Supplier
	}

	type StockLevel struct {
		Product // If a field is defined without a name, it is an embedded field, and it is accessed using the name of its type.
		count   int
	}

	kayak := Product{
		name:     "kayak",
		category: "Watersports",
		price:    275,
	}

	acme := &Supplier{"Acme co", "New York"}

	stockItem := StockLevel{
		Product: Product{"Kayak", "Watersports", 275, acme},
		count:   100,
	}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	fmt.Println(stockItem.Product.name)

	// Struct values are comparable if all their fields can be compared
	// Structs cannot be compared if the struct type defines fields with incomparable types, such as slices,
	p1 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	p2 := Product{name: "Kayak", category: "Watersports", price: 275.00}
	fmt.Println("p1 == p2:", p1 == p2)

	// Item struct can be converted to Product cuz they define the same fields in the same order!
	prod := Product{name: "Kayak", category: "Watersports", price: 275.00, Supplier: acme}
	item := Item{category: "Watersports", name: "Kayak", price: 275.00, Supplier: acme}
	fmt.Println("prod == item:", prod == Product(item))

	// struct with pointer
	p3 := Product{name: "Kayak", category: "Watersports", price: 275}
	p4 := &p3
	p3.name = "Original Kayak"
	fmt.Println("p3:", p3.name)    // p3: Original Kayak
	fmt.Println("p4:", (*p4).name) // p4: Original Kayak

	calcTax(&kayak)
	kayak2 := &Product{name: "Kayak", category: "Watersports", price: 275.00}
	kayak3 := calcTax(kayak2)
	fmt.Println(kayak3.price)

	products := [2]*Product{newProduct("Kayak", "Watersports", 275, acme), newProduct("Hat", "Skiing", 42.50, acme)}
	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price:", p.price, "Supplier:", p.Supplier.name, p.Supplier.city)
	}

	p5 := newProduct("Kayak", "Watersports", 275, acme)
	p6 := *p5
	p5.name = "Original Kayak"
	p5.Supplier.name = "BoatCo"
	for _, p := range []Product{*p5, p6} {
		fmt.Println("Name:", p.name, ", Supplier:", p.Supplier.name)
	}
	// shallow copy: copy the pointer assigned to supplier
	// Name: Original Kayak , Supplier: BoatCo
	// Name: Kayak , Supplier: BoatCo

	// deep copy
	p7 := copyProduct(p5)
	p5.name = "won't affect p7"
	fmt.Println("Name:", p7.name, ", Supplier:", p7.Supplier.name)
}
