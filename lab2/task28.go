package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task28() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введіть рядок: ")
	scanner.Scan()
	str := scanner.Text()

	strField := strings.Fields(str)

	var a int
	var b int
	for {
		fmt.Print("Введіть число A: ")
		scanner.Scan()
		aStr := scanner.Text()
		var err error
		a, err = strconv.Atoi(aStr)
		if err != nil {
			fmt.Println("Введіть ціле число!")
			continue
		}
		if a < 0 {
			fmt.Println("Введіть невідʼємне число!")
			continue
		}
		if a >= len(strField) {
			fmt.Println("Введіть число менше за довжину рядка!")
			continue
		}
		break
	}
	for {
		fmt.Print("Введіть число B: ")
		scanner.Scan()
		bStr := scanner.Text()
		var err error
		b, err = strconv.Atoi(bStr)
		if err != nil {
			fmt.Println("Введіть ціле число!")
			continue
		}
		if b <= 0 {
			fmt.Println("Введіть додатне число!")
			continue
		}
		if b >= len(strField) {
			fmt.Println("Введіть число менше за довжwину рядка!")
			continue
		}
		if b < a {
			fmt.Println("Введіть число більше за A!")
			continue
		}
		break
	}

	strField = append(strField[:a], strField[b+1:]...)

	fmt.Println(strings.Join(strField, " "))
}
