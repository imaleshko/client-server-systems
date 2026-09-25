package main

import (
	"bufio"
	"fmt"
	"os"
)

func task9() {
	scanner := bufio.NewScanner(os.Stdin)
	var str1, str2 string

	for {
		fmt.Print("Введіть перший рядок: ")
		scanner.Scan()
		str1 = scanner.Text()

		fmt.Print("Введіть другий рядок: ")
		scanner.Scan()
		str2 = scanner.Text()

		if len(str1) == 0 || len(str2) == 0 {
			fmt.Printf("Введіть рядки! \n \n")
			continue
		}

		break
	}

	runes1 := []rune(str1)
	runes2 := []rune(str2)

	result := string([]rune{
		runes1[0], runes1[len(runes1)/2], runes1[len(runes1)-1],
		runes2[0], runes2[len(runes2)/2], runes2[len(runes2)-1],
	})

	fmt.Println(result)
}
