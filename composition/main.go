package main

import (
	"composition/store"
	"fmt"
)

func main() {
	// kayak := store.NewProduct("kayak", "Watersports", 275)
	// lifejacket := &store.Product{Name: "lifejacket", Category: "Watersports"}
	// for _, p := range []*store.Product{kayak, lifejacket} {
	// 	fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	// }

	// boats := []*store.Boat{
	// 	store.NewBoat("Kayak", 275, 1, false),
	// 	store.NewBoat("Canoe", 400, 3, false),
	// 	store.NewBoat("Tender", 650.25, 2, true),
	// }
	// boats[0].Name = "Green Kayak"
	// for _, b := range boats {
	// 	fmt.Println("Conventional:", b.Product.Name, "Direct (field promotion):", b.Name, "Price (promotion):", b.Price(0.2))
	// }

	rentals := []*store.RentalBoat{
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}
	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
	}
}
