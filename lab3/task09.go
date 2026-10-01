package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func task9() {
	scanner := bufio.NewScanner(os.Stdin)

	x := readCoordinate09(scanner, "Введіть x: ")
	y := readCoordinate09(scanner, "Введіть y: ")

	isXPositive := x > 0
	isYPositive := y > 0

	if x == 0 && y == 0 {
		fmt.Printf("Точка {%g, %g} є центром координат\n", x, y)
	} else if x == 0 {
		if isYPositive {
			fmt.Printf("Точка {%g, %g} знаходиться на перетині 1 та 2 чверті\n", x, y)
		} else {
			fmt.Printf("Точка {%g, %g} знаходиться на перетині 3 та 4 чверті\n", x, y)
		}
	} else if y == 0 {
		if isXPositive {
			fmt.Printf("Точка {%g, %g} знаходиться на перетині 1 та 4 чверті\n", x, y)
		} else {
			fmt.Printf("Точка {%g, %g} знаходиться на перетині 2 та 3 чверті\n", x, y)
		}
	} else {
		if isXPositive && isYPositive {
			fmt.Printf("Точка {%g, %g} знаходиться в 1 чверті\n", x, y)
		} else if !isXPositive && isYPositive {
			fmt.Printf("Точка {%g, %g} знаходиться в 2 чверті\n", x, y)
		} else if !isXPositive {
			fmt.Printf("Точка {%g, %g} знаходиться в 3 чверті\n", x, y)
		} else {
			fmt.Printf("Точка {%g, %g} знаходиться в 4 чверті\n", x, y)
		}
	}
}

func readCoordinate09(scanner *bufio.Scanner, prompt string) float64 {
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
