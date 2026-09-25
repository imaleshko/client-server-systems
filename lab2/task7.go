package main

import (
	"bufio"
	"fmt"
	"os"
)

func task7() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введіть перший рядок: ")
	scanner.Scan()
	str1 := scanner.Text()

	fmt.Print("Введіть другий рядок: ")
	scanner.Scan()
	str2 := scanner.Text()

	runes1 := []rune(str1)

	middle := len(runes1) / 2

	result := string(runes1[:middle]) + str2 + string(runes1[middle:])
	fmt.Println(result)
}
