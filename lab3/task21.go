package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task21() {
	scanner := bufio.NewScanner(os.Stdin)
	product := 1

	fmt.Print("Введіть числа: ")
	scanner.Scan()
	str := scanner.Text()

	fieldStr := strings.FieldsSeq(str)

	for str := range fieldStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Print("У введеному рядку є нечисловий елемент\n")
			continue
		}
		product *= num
	}

	fmt.Printf("Добуток чисел: %d\n", product)
}
