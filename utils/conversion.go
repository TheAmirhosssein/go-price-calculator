package utils

import (
	"strconv"
)

func StringToFloatSlice(stringSlice []string) ([]float64, error) {
	convertedFloats := make([]float64, len(stringSlice))
	for index, value := range stringSlice {
		convertedFloat, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return nil, err
		}
		convertedFloats[index] = convertedFloat
	}
	return convertedFloats, nil
}
