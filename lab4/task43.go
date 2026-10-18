package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task43() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть рядок: ")
	scanner.Scan()
	str := scanner.Text()

	getCharacter := getCharacterGetter(str)

	var num int
	var err error
	for {
		fmt.Println("Введіть індекс символа: ")
		scanner.Scan()
		num, err = strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil {
			fmt.Println("Помилка: введіть одне число")
			continue
		}

		break
	}

	fmt.Println(getCharacter(num))
}

func getCharacterGetter(str string) func(int) string {
	runes := []rune(str)

	return func(index int) string {
		if index < 0 || index >= len(runes) {
			return ""
		}
		return string(runes[index])
	}
}
