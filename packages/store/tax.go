package store

const defaultTaxRate = 0.2
const minThreshold = 10 // less than 10, won't charge tax

type taxRate struct {
	rate, threshold float64
}

func newTaxRate(rate float64, threshold float64) *taxRate {
	if rate == 0 {
		rate = defaultTaxRate
	}
	if threshold < minThreshold {
		threshold = minThreshold
	}

	return &taxRate{rate, threshold}
}

func (taxRate *taxRate) calcTax(product *Product) float64 {
	if product.price > taxRate.threshold {
		return product.price * (1 + taxRate.rate)
	}

	return product.price
}
