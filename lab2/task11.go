package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func task11() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введіть рядок: ")
	scanner.Scan()
	str := scanner.Text()

	strField := strings.Fields(str)

	fmt.Printf("Отримано масив із %d елементів %v", len(strField), strField)
}
