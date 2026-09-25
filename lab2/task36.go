package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task36() {
	scanner := bufio.NewScanner(os.Stdin)
	var keys []int
	var values []string
	for {
		var keysTemp []int

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

		fmt.Print("Введіть рядок значень: ")
		scanner.Scan()
		str2 := scanner.Text()
		values = strings.Fields(str2)

		if len(keysTemp) != len(values) {
			fmt.Println("Введіть рядки однакової довжини")
			continue
		}

		keys = keysTemp
		break
	}

	fmt.Print("Введіть значення для видалення: ")
	scanner.Scan()
	a := scanner.Text()

	strMap := make(map[int]string)

	for i := 0; i < len(keys); i++ {
		strMap[keys[i]] = values[i]
	}

	for key, value := range strMap {
		if value == a {
			delete(strMap, key)
		}
	}

	fmt.Println(strMap)
}
