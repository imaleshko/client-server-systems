package main

import (
	"errors"
	"fmt"
	"strconv"
)

func wrapError(err error, context string) error {
	if err == nil {
		return nil
	}

	return fmt.Errorf("%s: %w", context, err)
}

func task27() {
	task27Error := errors.New("помилка у task27")

	err := task27Error
	for i := 1; i <= 5; i++ {
		err = wrapError(err, "загортання №"+strconv.Itoa(i))
	}

	fmt.Println("Загорнута помилка:")
	fmt.Println(err)

	fmt.Println("\nРозгортання помилки:")
	for current := err; current != nil; current = errors.Unwrap(current) {
		fmt.Println(current)
	}

	fmt.Println("\nПеревірка наявності загорнутої помилки:")
	fmt.Println(errors.Is(err, task27Error))
}
