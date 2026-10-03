package main

import "fmt"

func task29() {
	var sum int

	for i := -54; i < 15; i++ {
		if i%4 == 0 {
			sum += i
		}
	}

	fmt.Printf("Сума: %d\n", sum)
}
