package main

import (
	"fmt"

	"lab7/errors"
)

func task52() {
	baseErr := errors.New("запис не знайдено")
	wrappedErr := errors.Wrap("отримання користувача", baseErr)
	deepErr := errors.Wrap("обробка запиту", wrappedErr)

	fmt.Println("Помилка:", deepErr)

	otherErr := errors.New("час очікування вичерпано")
	fmt.Println("Is(deepErr, baseErr):", errors.Is(deepErr, baseErr))
	fmt.Println("Is(deepErr, otherErr):", errors.Is(deepErr, otherErr))

	permBase := &PermissionError{User: "user", Action: "write"}
	wrappedPerm := errors.Wrap("панель адміністратора", permBase)

	var permErr *PermissionError
	if errors.As(wrappedPerm, &permErr) {
		fmt.Printf("Знайдено PermissionError: користувач = %s, дія = %s\n", permErr.User, permErr.Action)
	}
}
