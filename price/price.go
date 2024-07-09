package price

type TaxIncludeJobPrice struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}
