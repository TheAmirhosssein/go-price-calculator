package main

import (
	"fmt"

	"github.com/TheAmirhosssein/go-price-calculator/cmdmanager"
	"github.com/TheAmirhosssein/go-price-calculator/filemanger"
	"github.com/TheAmirhosssein/go-price-calculator/iomanager"
	"github.com/TheAmirhosssein/go-price-calculator/price"
)

func main() {
	taxRates := []float64{0, 5, 6}

	for _, taxRate := range taxRates {
		var command string
		fmt.Print("enter 'file' for file entry and 'cmd' for cmd: ")
		fmt.Scan(&command)
		var io iomanager.IOManager
		if command == "cmd" {
			io = cmdmanager.New()
		} else if command == "file" {
			io = filemanger.New("price.txt", "result.json")
		} else {
			fmt.Println("are you fucking retarded?")
			return
		}
		jobTax := price.NewTaxIncludedPriceJob(io, taxRate)
		jobTax.Process()
	}
}
