package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/shopspring/decimal"
)

type Product struct {
	name  string
	price decimal.Decimal
}

type Order struct {
	products []Product
	discount decimal.Decimal
}

func (o *Order) add(name string, price decimal.Decimal) error {
	if name == "" {
		return errors.New("назва продукту не може бути порожньою")
	}
	if price.Compare(decimal.Zero) <= 0 {
		return errors.New("ціна повинна бути додатним числом")
	}

	product := Product{
		name:  name,
		price: price,
	}
	o.products = append(o.products, product)

	return nil
}

func (o *Order) totalPriceOfProductInOrder(productName string) (decimal.Decimal, error) {
	var totalPrice decimal.Decimal
	isFound := false
	for _, p := range o.products {
		if p.name == productName {
			totalPrice = totalPrice.Add(p.price)
			isFound = true
		}
	}

	if !isFound {
		return decimal.Zero, errors.New("продукт не знайдено в замовленні")
	}

	return totalPrice, nil
}

func (o *Order) totalPriceOfOrder() decimal.Decimal {
	var totalPrice decimal.Decimal
	for _, p := range o.products {
		totalPrice = totalPrice.Add(p.price)
	}

	if !o.discount.IsZero() {
		one := decimal.NewFromInt(1)
		hundred := decimal.NewFromInt(100)

		totalPrice = totalPrice.Mul(one.Sub(o.discount.Div(hundred)))
	}

	return totalPrice
}

func (o *Order) discountProduct(productName string, percentage decimal.Decimal) error {
	one := decimal.NewFromInt(1)
	hundred := decimal.NewFromInt(100)
	if percentage.Compare(one) < 0 || percentage.Compare(hundred) > 0 {
		return errors.New("відсоток знижки має бути від 1 до 100")
	}

	isFound := false
	discount := percentage.Div(hundred)
	for i := range o.products {
		if o.products[i].name == productName {
			discountPrice := o.products[i].price.Mul(discount)
			o.products[i].price = o.products[i].price.Sub(discountPrice)
			isFound = true
		}
	}

	if !isFound {
		return errors.New("продукт не знайдено в замовленні")
	}

	return nil
}

func (o *Order) discountOrder(percentage decimal.Decimal) error {
	one := decimal.NewFromInt(1)
	hundred := decimal.NewFromInt(100)
	if percentage.Compare(one) < 0 || percentage.Compare(hundred) > 0 {
		return errors.New("відсоток знижки має бути від 1 до 100")
	}

	o.discount = percentage

	return nil
}

func (o *Order) print() {
	fmt.Println("Друк замовлення:")

	if len(o.products) == 0 {
		fmt.Println("Замовлення пусте")
		return
	}

	fmt.Printf("Сума замовлення: %s\n", o.totalPriceOfOrder().StringFixed(2))

	for i, p := range o.products {
		fmt.Printf("Продукт №%d: %s - %s\n", i+1, p.name, p.price.StringFixed(2))
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
			price, err := decimal.NewFromString(priceStr)
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
			percentage, err := decimal.NewFromString(percentageStr)
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
			percentage, err := decimal.NewFromString(percentageStr)
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

			fmt.Printf("Загальна вартість товарів типу '%s': %s\n", name, total.StringFixed(2))
		case "6":
			return
		default:
			fmt.Println("Помилка: виберіть опцію серед наведених")
		}
	}
}
