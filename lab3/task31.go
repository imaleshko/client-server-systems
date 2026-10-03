package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func task31() {
	scanner := bufio.NewScanner(os.Stdin)
	var z int
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
		result += (math.Sqrt(math.Pow(nF+2.5*nF, 3))) / 4
	}

	fmt.Printf("Значення виразу: %g\n", result)
}
