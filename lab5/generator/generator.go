package generator

import (
	"errors"
	"math"
)

const precision float64 = 1e10

func GenerateSliceInt() []int {
	slice := make([]int, 0, 25)

	for i := 23; i <= 47; i++ {
		slice = append(slice, i)
	}

	return slice
}

func GenerateSliceFloat(step float64) (slice []float64, err error) {
	if step <= 0 {
		return slice, errors.New("крок повинен бути більшим за 0")
	}

	if step >= 1 {
		return slice, errors.New("крок повинен бути меншим за 1")
	}

	for i := 0; ; i++ {
		num := float64(i) * step
		num = math.Round(num*precision) / precision

		if num > 1 {
			break
		}

		slice = append(slice, num)
	}

	return slice, nil
}
