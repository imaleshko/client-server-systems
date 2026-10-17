package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func task22() {
	scanner := bufio.NewScanner(os.Stdin)
	map1 := map[string]int{}
	var map2 map[string]int

	for {
		fmt.Println("Введіть координати x та y одної або двох точок: ")
		scanner.Scan()

		input := strings.Fields(scanner.Text())

		if len(input) != 2 && len(input) != 4 {
			fmt.Println("Введіть 2 або 4 числа")
			continue
		}

		var err error

		map1["x"], err = strconv.Atoi(input[0])
		if err != nil {
			fmt.Println("Введіть цілі числа")
			continue
		}
		map1["y"], err = strconv.Atoi(input[1])
		if err != nil {
			fmt.Println("Введіть цілі числа")
			continue

		}

		map2 = nil
		if len(input) == 4 {
			map2 = map[string]int{}
			map2["x"], err = strconv.Atoi(input[2])
			if err != nil {
				fmt.Println("Введіть цілі числа")
				continue
			}
			map2["y"], err = strconv.Atoi(input[3])
			if err != nil {
				fmt.Println("Введіть цілі числа")
				continue
			}
		}

		break
	}

	if map2 == nil {
		fmt.Println(calculateDistance(map1))
	} else {
		fmt.Println(calculateDistance(map1, map2))
	}
}

func calculateDistance(point1 map[string]int, point2 ...map[string]int) float64 {
	p2 := map[string]int{"x": -7, "y": 3}

	if len(point2) > 0 {
		p2 = point2[0]
	}

	return math.Hypot(float64(p2["x"]-point1["x"]), float64(p2["y"]-point1["y"]))
}
