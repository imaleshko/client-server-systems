package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func task57() {
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

	fmt.Println((math.Tan(n) - 2*n) / math.Sqrt(10+0.6*n))
}
