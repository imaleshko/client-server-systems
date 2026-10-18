package main

import (
	"bufio"
	"fmt"
	"os"
)

func task47() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть рядок: ")
	scanner.Scan()
	str := scanner.Text()

	getCharacter := counterGetter(str)

	var char rune
	for {
		fmt.Println("Введіть один символ: ")
		scanner.Scan()
		inputRunes := []rune(scanner.Text())

		if len(inputRunes) != 1 {
			fmt.Println("Помилка: введіть рівно один символ")
			continue
		}

		char = inputRunes[0]
		break
	}

	fmt.Println(getCharacter(char))
}

func counterGetter(str string) func(rune) int {
	runes := []rune(str)

	return func(r rune) int {
		counter := 0
		for _, v := range runes {
			if v == r {
				counter++
			}
		}
		return counter
	}
}
