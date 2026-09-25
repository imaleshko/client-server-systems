package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task30() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Введіть числа: ")
	scanner.Scan()
	str := scanner.Text()

	fieldStr := strings.Fields(str)

	var nums []int

	for _, str := range fieldStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Print("Нечисловий елемент пропущено\n")
			continue
		}
		nums = append(nums, num)
	}

	var oddNums []int

	for _, num := range nums {
		if num%2 != 0 {
			oddNums = append(oddNums, num)
		}
	}

	fmt.Println("Непарні числа:", oddNums)
}
