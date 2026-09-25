package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"strconv"
	"strings"
)

func task38() {
	scanner := bufio.NewScanner(os.Stdin)

	keys1, values1 := readKeysAndValues(scanner)
	keys2, values2 := readKeysAndValues(scanner)

	mp1 := make(map[int]int)

	for i := range keys1 {
		mp1[keys1[i]] = values1[i]
	}

	mp1Len := len(mp1)

	mp2 := make(map[int]int)

	for i := range keys2 {
		mp2[keys2[i]] = values2[i]
	}

	maps.Copy(mp1, mp2)

	fmt.Println(mp1)
	fmt.Printf("У першій map %d елементів, у другій - %d, в обʼєднаній - %d", mp1Len, len(mp2), len(mp1))
}

func readKeysAndValues(scanner *bufio.Scanner) ([]int, []int) {
	var keys []int
	var values []int
	for {
		var keysTemp []int
		var valuesTemp []int

		fmt.Print("Введіть рядок ключів (цілі числа): ")
		scanner.Scan()
		str1 := scanner.Text()
		str1Field := strings.FieldsSeq(str1)

		for str := range str1Field {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Print("Нечисловий елемент пропущено\n")
				continue
			}
			keysTemp = append(keysTemp, num)
		}

		fmt.Print("Введіть рядок значень (цілі числа): ")
		scanner.Scan()
		str2 := scanner.Text()
		str2Field := strings.FieldsSeq(str2)

		for str := range str2Field {
			num, err := strconv.Atoi(str)
			if err != nil {
				fmt.Print("Нечисловий елемент пропущено\n")
				continue
			}
			valuesTemp = append(valuesTemp, num)
		}

		if len(keysTemp) != len(valuesTemp) {
			fmt.Println("Введіть рядки однакової довжини")
			continue
		}

		keys = keysTemp
		values = valuesTemp
		break
	}

	return keys, values
}
