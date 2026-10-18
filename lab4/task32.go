package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task32() {
	scanner := bufio.NewScanner(os.Stdin)
	var err error
	var n float64
	var k int

	for {
		fmt.Println("Введіть число n: ")
		scanner.Scan()

		n, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Println("Помилка: введіть одне число")
			continue
		}

		break
	}

	for {
		fmt.Println("Введіть число k: ")
		scanner.Scan()

		k, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || k < 1 {
			fmt.Println("Помилка: введіть одне ціле число >= 1")
			continue
		}

		break
	}

	fmt.Println(nToPowerOfK(n, k))
}

func nToPowerOfK(n float64, k int) float64 {
	if k == 0 {
		return 1
	}

	if k < 0 {
		return 1 / nToPowerOfK(n, -k)
	}

	return n * nToPowerOfK(n, k-1)
}
