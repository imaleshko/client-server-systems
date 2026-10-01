package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task17() {
	var temperature float64
	var system string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Введіть температуру: ")
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))

		if len(input) < 2 {
			fmt.Println("Введіть коректну температуру")
			continue
		}

		system = string(input[len(input)-1])

		if system != "c" && system != "f" {
			fmt.Println("Введіть коректну систему вимірювання")
			continue
		}

		numStr := input[:len(input)-1]
		var err error
		temperature, err = strconv.ParseFloat(numStr, 64)
		if err != nil {
			fmt.Println("Введіть коректну температуру")
			continue
		}

		break
	}

	if system == "c" {
		fmt.Printf("Температура в градусах Фаренгейта: %.2f\n", temperature*1.8+32)
	} else {
		fmt.Printf("Температура в градусах Цельсія: %.2f\n", (temperature-32)*5/9)
	}
}
