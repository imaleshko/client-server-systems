package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task22() {
	scanner := bufio.NewScanner(os.Stdin)
	var sum int
	var nums []int

	fmt.Print("Введіть числа: ")
	scanner.Scan()

	fieldStr := strings.FieldsSeq(scanner.Text())

	for element := range fieldStr {
		num, err := strconv.Atoi(element)
		if err != nil {
			fmt.Print("У введеному рядку є нечисловий елемент\n")
			continue
		}
		nums = append(nums, num)
	}

	for i := range nums {
		if i%2 != 0 {
			sum += nums[i]
		}
	}

	fmt.Printf("Сума елементів з непарним індексом: %d\n", sum)
}
