package price

import (
	"fmt"

	"github.com/TheAmirhosssein/go-price-calculator/filemanger"
	"github.com/TheAmirhosssein/go-price-calculator/utils"
)

type TaxIncludeJobPrice struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job TaxIncludeJobPrice) Process() {
	job.loadData()
	for _, price := range job.InputPrices {
		job.TaxIncludedPrice[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	err := filemanger.WriteJson(job.TaxIncludedPrice, fmt.Sprintf("%.2f", job.TaxRate))
	if err != nil {
		fmt.Println(err)
	}
}

func (job *TaxIncludeJobPrice) loadData() {

	line, err := filemanger.ReadFile("price.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	price, err := utils.StringToFloatSlice(line)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = price

}

func NewTaxIncludeJobPrice(taxRate float64) TaxIncludeJobPrice {
	return TaxIncludeJobPrice{
		TaxRate:          taxRate,
		TaxIncludedPrice: make(map[string]float64),
	}
}
