package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task31() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть число: ")
	scanner.Scan()

	num, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil || num < 0 {
		fmt.Println("Помилка: введіть одне ціле число")
		return
	}

	fmt.Println(fib(num))
}

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	return fib(n-1) + fib(n-2)
}
