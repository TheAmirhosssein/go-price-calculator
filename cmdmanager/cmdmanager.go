package cmdmanager

import "fmt"

type CMDManager struct {
}

func (c *CMDManager) ReadLines() ([]string, error) {
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

func (c *CMDManager) WriteResult(data map[string]float64, key string) error {
	fmt.Println(data)
	return nil
}

func New() *CMDManager {
	return &CMDManager{}
}
