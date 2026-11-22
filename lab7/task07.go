package main

import (
	"fmt"
	"os"
	"strconv"
)

func task07() {
	inputs := os.Args

	if len(inputs) < 2 {
		fmt.Println("Помилка: введіть ціле число")
		return
	}

	num, err := strconv.Atoi(inputs[1])
	if err != nil {
		fmt.Println("Помилка: введіть ціле число")
		return
	}

	fmt.Println(num)
}
