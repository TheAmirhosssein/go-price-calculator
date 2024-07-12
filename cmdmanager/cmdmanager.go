package cmdmanager

import "fmt"

type stringFloatMap map[string]float64

type CMDManger struct {
}

func (cmd CMDManger) ReadFile() ([]string, error) {
	var prices []string

	for {
		var price string
		fmt.Print("enter the price please and 'e' to exit: ")
		fmt.Scan(&price)
		fmt.Print("\n")
		if price == "e" {
			break
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (cmd CMDManger) WriteJson(data stringFloatMap, key string) error {
	fmt.Println(data)
	return nil
}

func New() CMDManger {
	return CMDManger{}
}
