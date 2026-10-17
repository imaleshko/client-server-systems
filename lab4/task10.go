package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task10() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть число: ")
	scanner.Scan()

	num, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	if err != nil {
		fmt.Println("Помилка: введіть одне коректне число")
		return
	}

	fmt.Println(toThirdPower(num))
}

func toThirdPower(num float64) float64 {
	return num * num * num
}
