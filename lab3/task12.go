package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func task12() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть значення:")
	scanner.Scan()

	_, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("Other")
		return
	}

	fmt.Println("Number")
}
