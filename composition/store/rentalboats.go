package store

type RentalBoat struct {
	*Boat
	IncludeCrew bool
}

func NewRentalBoat(name string, price float64, capacity int, motorized, crewed bool) *RentalBoat {
	return &RentalBoat{
		NewBoat(name, price, capacity, motorized),
		crewed,
	}
}
