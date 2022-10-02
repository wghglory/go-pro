package store

const defaultTaxRate = 0.2
const minThreshold = 10 // less than 10, won't charge tax

// var categoryMaxPrices = map[string]float64{
// 	"Watersports": 250 * (1 + defaultTaxRate),
// 	"Soccer":      150 * (1 + defaultTaxRate),
// 	"Chess":       50 * (1 + defaultTaxRate),
// }

var categoryMaxPrices = map[string]float64{
	"Watersports": 250,
	"Soccer":      150,
	"Chess":       50,
}

// The init function is called automatically and provides an opportunity to prepare the package for use
// a single file can define multiple init functions, all of which will be executed.
func init() {
	for category, price := range categoryMaxPrices {
		categoryMaxPrices[category] = price * (1 + defaultTaxRate)
	}
}

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

func (taxRate *taxRate) calcTax(product *Product) (price float64) {
	if product.price > taxRate.threshold {
		price = product.price * (1 + taxRate.rate)
	} else {
		price = product.price
	}

	if max, ok := categoryMaxPrices[product.Category]; ok && price > max {
		price = max
	}

	return
}
