package main

import "fmt"

func main() {
	prices := []float64{20, 30}
	taxRates := []float64{0, 0.07, 0.1, 0.2}
	taxRatesPrices := make(map[float64][]float64)

	for _, taxRate := range taxRates {
		calculatedTaxRate := make([]float64, len(prices))
		for priceIndex, price := range prices {
			calculatedTaxRate[priceIndex] = price * (1 + taxRate)
		}
		taxRatesPrices[taxRate] = calculatedTaxRate
	}

	fmt.Println(taxRatesPrices)
}
