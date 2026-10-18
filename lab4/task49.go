package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func task49() {
	scanner := bufio.NewScanner(os.Stdin)
	var err error

	var nums []float64
	var num float64
	fmt.Println("Введіть зріз чисел: ")
	scanner.Scan()
	fieldInput := strings.FieldsSeq(scanner.Text())

	for n := range fieldInput {
		num, err = strconv.ParseFloat(n, 64)
		if err != nil {
			fmt.Println("Нечисловий елемент пропущено")
			continue
		}
		nums = append(nums, num)
	}

	getSliceInPower := getSliceInPowerGetter(nums)

	var power float64
	for {
		fmt.Println("Введіть степінь: ")
		scanner.Scan()
		power, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("Помилка: введіть одне число")
			continue
		}

		break
	}

	fmt.Println(getSliceInPower(power))
}

func getSliceInPowerGetter(nums []float64) func(power float64) []float64 {
	return func(power float64) []float64 {
		result := make([]float64, 0, len(nums))
		for _, num := range nums {
			result = append(result, math.Pow(num, power))
		}
		return result
	}
}
