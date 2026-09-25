package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func task59() {
	scanner := bufio.NewScanner(os.Stdin)
	var intN int
	for {
		fmt.Print("Введіть ціле число: ")
		scanner.Scan()
		var err error
		intN, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Введіть ціле число!")
			continue
		}

		break
	}

	n := float64(intN)

	fmt.Println((10 + 2*math.Cos(n)) / (5 - math.Sqrt(math.Pow(n, 5))))
}
