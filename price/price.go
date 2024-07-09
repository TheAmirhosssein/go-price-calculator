package price

import "fmt"

type TaxIncludeJobPrice struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job *TaxIncludeJobPrice) Process() {
	for _, price := range job.InputPrices {
		job.TaxIncludedPrice[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
}

func NewTaxIncludeJobPrice(taxRate float64) TaxIncludeJobPrice {
	return TaxIncludeJobPrice{
		TaxRate:          taxRate,
		InputPrices:      []float64{20, 10},
		TaxIncludedPrice: make(map[string]float64),
	}
}
