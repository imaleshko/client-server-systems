package main

import (
	"bufio"
	"fmt"
	"os"
)

func task36() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть рядок: ")
	scanner.Scan()

	fmt.Println(cutString([]rune(scanner.Text())))
}

func cutString(runes []rune) string {
	if len(runes) < 2 {
		return ""
	}

	if runes[0] != '(' {
		return cutString(runes[1:])
	}

	if runes[len(runes)-1] != ')' {
		return cutString(runes[:len(runes)-1])
	}

	return string(runes[1 : len(runes)-1])
}
