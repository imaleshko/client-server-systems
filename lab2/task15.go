package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func task15() {
	scanner := bufio.NewScanner(os.Stdin)
	var a float64
	var b int

	for {
		fmt.Print("Введіть число A: ")
		scanner.Scan()
		aStr := scanner.Text()
		var err error
		a, err = strconv.ParseFloat(aStr, 64)
		if err != nil {
			fmt.Println("Введіть число!")
			continue
		}
		break
	}

	for {
		fmt.Print("Введіть число B: ")
		scanner.Scan()
		bStr := scanner.Text()
		var err error
		b, err = strconv.Atoi(bStr)
		if err != nil {
			fmt.Println("Введіть ціле число!")
			continue
		}
		if b <= 0 {
			fmt.Println("Введіть додатне число!")
			continue
		}
		break
	}

	var fieldStr []string
	var nums []float64

	for {
		fmt.Print("Введіть числа: ")
		scanner.Scan()
		str := scanner.Text()
		fieldStr = strings.Fields(str)

		nums = make([]float64, 0)

		for _, str := range fieldStr {
			num, err := strconv.ParseFloat(str, 64)
			if err != nil {
				fmt.Print("Нечисловий елемент пропущено\n")
				continue
			}
			nums = append(nums, num)
		}

		if b > len(nums)+1 {
			fmt.Println("Введіть більший масив")
			continue
		}
		break
	}

	nums = append(nums, 0)

	for i := len(nums) - 1; i > b-1; i-- {
		nums[i] = nums[i-1]
	}

	nums[b-1] = a

	fmt.Print(nums)
}
