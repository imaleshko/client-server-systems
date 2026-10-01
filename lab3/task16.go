package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func task16() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть номер місяця: ")
	scanner.Scan()

	month, err := strconv.Atoi(scanner.Text())
	if err != nil || month < 1 || month > 12 {
		fmt.Println("Введіть значення від 1 до 12")
		return
	}

	switch month {
	case 2:
		fmt.Println("28")
	case 4, 6, 9, 11:
		fmt.Println("30")
	default:
		fmt.Println("31")
	}
}
