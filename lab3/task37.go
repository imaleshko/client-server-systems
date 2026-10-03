package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func task37() {
	scanner := bufio.NewScanner(os.Stdin)
	var z int
	var sum float64
	var result float64

	for {
		var err error

		fmt.Print("Введіть число: ")
		scanner.Scan()

		z, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Введіть число")
			continue
		}

		break
	}

	for n := 1; n <= z; n++ {
		nF := float64(n)
		sum += 2 * math.Cos(nF)
	}

	result = (10 + sum) / (5 - math.Sqrt(math.Pow(1, 5)))

	fmt.Printf("Значення виразу: %g\n", result)
}
