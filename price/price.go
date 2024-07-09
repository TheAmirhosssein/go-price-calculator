package price

type TaxIncludeJobPrice struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func NewTaxIncludeJobPrice(taxRate float64) *TaxIncludeJobPrice {
	return &TaxIncludeJobPrice{
		TaxRate:          taxRate,
		InputPrices:      []float64{20, 10, 30},
		TaxIncludedPrice: make(map[string]float64),
	}
}
