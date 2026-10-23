package main

import (
	"bufio"
	"fmt"
	"generator"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\n1 - переглянути масив цілих чисел від 23 до 47")
		fmt.Println("2 - переглянути масив чисел з плаваючою комою від 0 до 1 із заданим кроком")
		fmt.Println("3 - вийти")
		fmt.Println("Виберіть опцію: ")
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			fmt.Println(generator.GenerateSliceInt())
		case "2":
			fmt.Println("Введіть крок")
			scanner.Scan()
			step, err := strconv.ParseFloat(scanner.Text(), 64)
			if err != nil {
				fmt.Println("Помилка: введено нечислове значення")
				continue
			}

			slice, err := generator.GenerateSliceFloat(step)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
				continue
			}
			fmt.Println(slice)
		case "3":
			return
		default:
			fmt.Println("Виберіть варіант")
			continue
		}
	}
}
