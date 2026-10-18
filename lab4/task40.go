package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task40() {
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	var num int

	for {
		fmt.Println("Введіть ціле число: ")
		scanner.Scan()
		num, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Помилка: введіть одне число")
			continue
		}

		break
	}

	fmt.Println(reverseNumber(num, 0))
}

func reverseNumber(n int, reversed int) int {
	if n == 0 {
		return reversed
	}

	return reverseNumber(n/10, reversed*10+n%10)
}
