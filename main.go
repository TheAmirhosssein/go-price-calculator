package main

import (
	"fmt"

	"github.com/TheAmirhosssein/go-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.2}

	for _, taxRate := range taxRates {
		jobTax := price.NewTaxIncludeJobPrice(taxRate)
		jobTax.Process()
		fmt.Println(jobTax.TaxIncludedPrice)
	}
}
