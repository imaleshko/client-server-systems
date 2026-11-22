package main

import "fmt"

func panicIfFalse(b bool) {
	if !b {
		panic("false у panicIfFalse")
	}
}

func demoRecover(b bool) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Спрацював recover:", r)
		} else {
			fmt.Println("Нічого не сталося")
		}
	}()

	panicIfFalse(b)
}

func task37() {
	fmt.Println("Якщо false:")
	demoRecover(false)
	fmt.Println("Якщо true:")
	demoRecover(true)
}
