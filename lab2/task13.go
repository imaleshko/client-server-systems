package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task13() {
	scanner := bufio.NewScanner(os.Stdin)
	var a int

	for {
		fmt.Print("Введіть число A: ")
		scanner.Scan()
		aStr := scanner.Text()
		var err error
		a, err = strconv.Atoi(aStr)
		if err != nil {
			fmt.Println("Введіть ціле число!")
			continue
		}
		break
	}

	fmt.Print("Введіть числа: ")
	scanner.Scan()
	str := scanner.Text()

	fieldStr := strings.Fields(str)

	nums := make([]int, 0)

	for _, str := range fieldStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Print("Нечисловий елемент пропущено\n")
			continue
		}
		nums = append(nums, num)
	}

	for i := range nums {
		nums[i] += a
	}

	nums = append(nums, a)

	fmt.Println(nums)
}
