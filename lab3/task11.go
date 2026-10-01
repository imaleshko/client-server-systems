package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func task11() {
	scanner := bufio.NewScanner(os.Stdin)

	x := readCoordinate11(scanner, "Введіть x: ")
	y := readCoordinate11(scanner, "Введіть y: ")

	if x*x+y*y <= 25 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func readCoordinate11(scanner *bufio.Scanner, prompt string) float64 {
	for {
		fmt.Println(prompt)
		scanner.Scan()
		input, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Значення має бути числом")
			continue
		}
		return input
	}
}
