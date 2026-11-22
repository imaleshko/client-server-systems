package main

import "fmt"

type PermissionError struct {
	User   string
	Action string
}

func (e *PermissionError) Error() string {
	return "дозвіл відхилено для " + e.User + " при спробі " + e.Action
}

func checkPermission(user, action string) error {
	if user != "admin" && action != "read" {
		return &PermissionError{User: user, Action: action}
	}

	return nil
}

func task19() {
	fmt.Println("user + read:")
	err := checkPermission("user", "read")
	if err != nil {
		fmt.Printf("Помилка: %v\n", err)
	} else {
		fmt.Println("Доступ дозволено")
	}

	fmt.Println("user + write:")
	err = checkPermission("user", "write")
	if err != nil {
		fmt.Printf("Помилка: %v\n", err)
	} else {
		fmt.Println("Доступ дозволено")
	}

	fmt.Println("user + execute:")
	err = checkPermission("user", "execute")
	if err != nil {
		fmt.Printf("Помилка: %v\n", err)
	} else {
		fmt.Println("Доступ дозволено")
	}

	fmt.Println("admin + write:")
	err = checkPermission("admin", "write")
	if err != nil {
		fmt.Printf("Помилка: %v\n", err)
	} else {
		fmt.Println("Доступ дозволено")
	}
}
