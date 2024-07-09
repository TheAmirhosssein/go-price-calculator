package utils

import (
	"strconv"
)

func StringToFloatSlice(stringSlice []string) ([]float64, error) {
	convertedFloats := make([]float64, len(stringSlice))
	for _, value := range stringSlice {
		convertedFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		convertedFloats = append(convertedFloats, convertedFloat)
	}
	return convertedFloats, nil
}
