package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func task45() {
	scanner := bufio.NewScanner(os.Stdin)
	var val1 float64
	var val2 float64

	for {
		fmt.Print("Введіть перше число: ")
		scanner.Scan()
		var err error
		val1, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Введіть число!")
			continue
		}
		break
	}

	for {
		fmt.Print("Введіть друге число: ")
		scanner.Scan()
		var err error
		val2, err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Введіть число!")
			continue
		}
		break
	}

	fmt.Println(math.Pow(val1, val2))
}
