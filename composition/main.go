package main

import (
	"composition/store"
	"fmt"
)

func main() {
	kayak := store.NewProduct("kayak", "Watersports", 275)
	lifejacket := &store.Product{Name: "lifejacket", Category: "Watersports"}

	for _, p := range []*store.Product{kayak, lifejacket} {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}
}
