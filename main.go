package main

import (
	"github.com/TheAmirhosssein/go-price-calculator/filemanger"
	"github.com/TheAmirhosssein/go-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 5, 6}

	for _, taxRate := range taxRates {
		fm := filemanger.New("price.txt", "result.json")
		jobTax := price.NewTaxIncludedPriceJob(fm, taxRate)
		jobTax.Process()
	}
}
