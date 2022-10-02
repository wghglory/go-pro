// package store provides types and methods
// commonly required for online sales
package store

// Product describes a product for sale
type Product struct {
	Name, Category string
	price          float64 // lowercase means field is available only within the package
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return p.price
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}
