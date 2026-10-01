package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task12() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть значення: ")
	scanner.Scan()

	input := strings.TrimSpace(scanner.Text())

	_, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Other")
		return
	}

	fmt.Println("Number")
}
