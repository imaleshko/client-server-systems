package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	name  string
	price float64
}

type Order struct {
	products []Product
	discount float64
}

func (o *Order) add(name string, price float64) error {
	if name == "" {
		return errors.New("назва продукту не може бути порожньою")
	}
	if price <= 0 {
		return errors.New("ціна повинна бути додатним числом")
	}

	product := Product{
		name:  name,
		price: price,
	}
	o.products = append(o.products, product)

	return nil
}

func (o *Order) totalPriceOfProductInOrder(productName string) (float64, error) {
	var totalPrice float64
	isFound := false
	for _, p := range o.products {
		if p.name == productName {
			totalPrice += p.price
			isFound = true
		}
	}

	if !isFound {
		return 0, errors.New("продукт не знайдено в замовленні")
	}

	return totalPrice, nil
}

func (o *Order) totalPriceOfOrder() float64 {
	var totalPrice float64
	for _, p := range o.products {
		totalPrice += p.price
	}

	return totalPrice * (1 - o.discount/100)
}

func (o *Order) discountProduct(productName string, percentage float64) error {
	if percentage < 1 || percentage > 100 {
		return errors.New("відсоток знижки має бути від 1 до 100")
	}

	isFound := false
	for i := range o.products {
		if o.products[i].name == productName {
			discountPrice := o.products[i].price * percentage / 100
			o.products[i].price -= discountPrice
			isFound = true
		}
	}

	if !isFound {
		return errors.New("продукт не знайдено в замовленні")
	}

	return nil
}

func (o *Order) discountOrder(percentage float64) error {
	if percentage < 1 || percentage > 100 {
		return errors.New("відсоток знижки має бути від 1 до 100")
	}

	o.discount = percentage

	return nil
}

func (o *Order) print() {
	fmt.Println("\nДрук замовлення:")

	if len(o.products) == 0 {
		fmt.Println("Замовлення пусте")
		return
	}

	fmt.Printf("Сума замовлення: %.2f\n", o.totalPriceOfOrder())

	for i, p := range o.products {
		fmt.Printf("Продукт №%d: %s - %.2f\n", i+1, p.name, p.price)
	}
}

func readLine(scanner *bufio.Scanner, prompt string) string {
	fmt.Println(prompt)
	scanner.Scan()
	return strings.TrimSpace(scanner.Text())
}

func task07() {
	scanner := bufio.NewScanner(os.Stdin)
	var order Order
	for {
		fmt.Println("\n1 - надрукувати замовлення")
		fmt.Println("2 - додати продукт до замовлення")
		fmt.Println("3 - зробити знижку на тип товарів у замовленні")
		fmt.Println("4 - зробити знижку на все замовлення")
		fmt.Println("5 - переглянути вартість товарів певного типу")
		fmt.Println("6 - вийти")

		switch readLine(scanner, "Виберіть опцію:") {
		case "1":
			order.print()
		case "2":
			name := readLine(scanner, "Введіть назву товару:")

			priceStr := readLine(scanner, "Введіть ціну товару:")
			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil {
				fmt.Println("Помилка: ціна повинна бути числом")
				continue
			}

			err = order.add(name, price)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
			}
		case "3":
			name := readLine(scanner, "Введіть назву товару:")

			percentageStr := readLine(scanner, "Введіть відсоток знижки:")
			percentage, err := strconv.ParseFloat(percentageStr, 64)
			if err != nil {
				fmt.Println("Помилка: відсоток знижки повинен бути числом")
				continue
			}

			err = order.discountProduct(name, percentage)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
			}
		case "4":
			percentageStr := readLine(scanner, "Введіть відсоток знижки:")
			percentage, err := strconv.ParseFloat(percentageStr, 64)
			if err != nil {
				fmt.Println("Помилка: відсоток знижки повинен бути числом")
				continue
			}

			err = order.discountOrder(percentage)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
			}
		case "5":
			name := readLine(scanner, "Введіть назву товару:")

			total, err := order.totalPriceOfProductInOrder(name)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
				continue
			}

			fmt.Printf("Загальна вартість товарів типу '%s': %.2f\n", name, total)
		case "6":
			return
		default:
			fmt.Println("Помилка: виберіть опцію серед наведених")
		}
	}
}
