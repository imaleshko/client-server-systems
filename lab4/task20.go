package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task20() {
	scanner := bufio.NewScanner(os.Stdin)

	num1 := read("Введіть перше число: ", scanner)
	num2 := read("Введіть друге число: ", scanner)

	fmt.Println(euclideanAlgorithm(num1, num2))
}

func read(message string, scanner *bufio.Scanner) int {
	for {
		fmt.Println(message)
		scanner.Scan()

		num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || num <= 0 {
			fmt.Println("Помилка: введіть одне додатне число")
			continue
		}

		return num
	}
}

func euclideanAlgorithm(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}

	return a
}
