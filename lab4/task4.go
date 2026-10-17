package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task4() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть довільну кількість чисел: ")
	scanner.Scan()

	fmt.Println(makeSlice(scanner.Text()))
}

func makeSlice(input string) []float64 {
	inputField := strings.Fields(input)

	nums := make([]float64, 0, len(inputField))

	for _, str := range inputField {
		num, err := strconv.ParseFloat(str, 64)
		if err != nil {
			fmt.Println("Нечисловий елемент пропущено")
			continue
		}
		nums = append(nums, num)
	}

	return nums
}
