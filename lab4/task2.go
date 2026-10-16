package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task2() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Введіть зріз цілих чисел: ")
	scanner.Scan()
	input := strings.Fields(scanner.Text())

	nums := make([]int, 0, len(input))

	for _, str := range input {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Println("Нечисловий елемент пропущено")
			continue
		}
		nums = append(nums, num)
	}

	fmt.Println(productOfSlice(nums))
}

func productOfSlice(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	result := 1

	for _, num := range nums {
		result *= num
	}

	return result
}
