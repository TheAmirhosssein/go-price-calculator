package price

import (
	"bufio"
	"fmt"
	"os"

	"github.com/TheAmirhosssein/go-price-calculator/utils"
)

type TaxIncludeJobPrice struct {
	TaxRate          float64
	InputPrices      []float64
	TaxIncludedPrice map[string]float64
}

func (job TaxIncludeJobPrice) Process(fileName string) {
	job.loadData(fileName)
	for _, price := range job.InputPrices {
		job.TaxIncludedPrice[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}
}

func (job *TaxIncludeJobPrice) loadData(path string) {
	file, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
		file.Close()
		return
	}

	scan := bufio.NewScanner(file)
	var fileLines []string
	for scan.Scan() {
		fileLines = append(fileLines, scan.Text())
	}
	file.Close()

	if scan.Err() != nil {
		fmt.Println(scan.Err())
		return
	}

	price, err := utils.StringToFloatSlice(fileLines)
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
