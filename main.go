package main

import (
	"github.com/TheAmirhosssein/go-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 5, 6}

	for _, taxRate := range taxRates {
		jobTax := price.NewTaxIncludeJobPrice(taxRate)
		jobTax.Process()
	}
}
