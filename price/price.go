package price

import (
	"fmt"

	"github.com/TheAmirhosssein/go-price-calculator/iomanager"
	"github.com/TheAmirhosssein/go-price-calculator/utils"
)

type TaxIncludeJobPrice struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
	IoManager        iomanager.IOManager
}

func (job TaxIncludeJobPrice) Process() {
	job.loadData()
	for _, price := range job.InputPrices {
		job.TaxIncludedPrice[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
	err := job.IoManager.WriteResult(job.TaxIncludedPrice, fmt.Sprintf("%.2f", job.TaxRate))
	if err != nil {
		fmt.Println(err)
	}
}

func (job *TaxIncludeJobPrice) loadData() {

	line, err := job.IoManager.ReadLines()
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

func NewTaxIncludedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxIncludeJobPrice {
	return &TaxIncludeJobPrice{
		IoManager:        iom,
		InputPrices:      []float64{10, 20, 30},
		TaxRate:          taxRate,
		TaxIncludedPrice: make(map[string]float64),
	}
}
