// package store provides types and methods
// commonly required for online sales
package store

var standardTax = newTaxRate(0.25, 20)

// Product describes a product for sale
type Product struct {
	Name, Category string
	price          float64 // lowercase means field is available only within the package
}

func NewProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func (p *Product) Price() float64 {
	return standardTax.calcTax(p)
}

func (p *Product) SetPrice(price float64) {
	p.price = price
}
