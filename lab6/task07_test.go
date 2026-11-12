package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		productName string
		price       float64
		isError     bool
	}{
		{"Успішне додавання", "Яблуко", 10.5, false},
		{"Порожня назва", "", 10.5, true},
		{"Нульова ціна", "Банан", 0, true},
		{"Від'ємна ціна", "Груша", -5.0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			order := &Order{}
			err := order.add(tc.productName, tc.price)

			if (err != nil) != tc.isError {
				t.Errorf("add() error = %v, isError %v", err, tc.isError)
			}

			if !tc.isError && len(order.Products) != 1 {
				t.Errorf("Очікувалось 1 товар в замовленні, отримано %d", len(order.Products))
			}
		})
	}
}

func TestTotalPriceOfProductInOrder(t *testing.T) {
	order := &Order{}
	err := order.add("Кава", 50)
	if err != nil {
		t.Errorf("totalPriceOfProductInOrder() error = %v)", err)
	}
	err = order.add("Чай", 30)
	if err != nil {
		t.Errorf("totalPriceOfProductInOrder() error = %v)", err)
	}
	err = order.add("Кава", 50)
	if err != nil {
		t.Errorf("totalPriceOfProductInOrder() error = %v)", err)
	}

	tests := []struct {
		name        string
		productName string
		wantTotal   float64
		isError     bool
	}{
		{"Кілька товарів одного типу", "Кава", 100, false},
		{"Один товар такого типу", "Чай", 30, false},
		{"Товар не знайдено", "Сік", 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := order.totalPriceOfProductInOrder(tc.productName)

			if (err != nil) != tc.isError {
				t.Errorf("totalPriceOfProductInOrder() error = %v, isError %v", err, tc.isError)
				return
			}
			if got != tc.wantTotal {
				t.Errorf("Очікувалось %.2f, отримано %.2f", tc.wantTotal, got)
			}
		})
	}
}

func TestTotalPriceOfOrder(t *testing.T) {
	order := &Order{}

	if got := order.totalPriceOfOrder(); got != 0 {
		t.Errorf("Очікувалось 0, отримано %.2f", got)
	}

	err := order.add("Телефон", 1000)
	if err != nil {
		t.Errorf("totalPriceOfOrder() error = %v)", err)
	}
	err = order.add("Чохол", 200)
	if err != nil {
		t.Errorf("totalPriceOfOrder() error = %v)", err)
	}

	if got := order.totalPriceOfOrder(); got != 1200 {
		t.Errorf("Очікувалось 1200, отримано %.2f", got)
	}

	err = order.discountOrder(10)
	if err != nil {
		t.Errorf("discountProduct() error = %v)", err)
	}
	expected := 1200.0 * 0.9
	if got := order.totalPriceOfOrder(); got != expected {
		t.Errorf("Очікувалось %.2f, отримано %.2f", expected, got)
	}
}

func TestDiscountProduct(t *testing.T) {
	tests := []struct {
		name          string
		productName   string
		percentage    float64
		isError       bool
		expectedPrice float64
	}{
		{"Успішна знижка 50%", "Ноутбук", 50, false, 500},
		{"Знижка менше 1%", "Ноутбук", 0, true, 1000},
		{"Знижка більше 100%", "Ноутбук", 150, true, 1000},
		{"Товар не знайдено", "Мишка", 10, true, 1000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			order := &Order{}
			err := order.add("Ноутбук", 1000)
			if err != nil {
				t.Errorf("discountProduct() error = %v)", err)
			}

			err = order.discountProduct(tc.productName, tc.percentage)

			if (err != nil) != tc.isError {
				t.Errorf("discountProduct() error = %v, isError %v", err, tc.isError)
			}

			if order.Products[0].Price != tc.expectedPrice {
				t.Errorf("Очікувана ціна %.2f, отримано %.2f", tc.expectedPrice, order.Products[0].Price)
			}
		})
	}
}

func TestDiscountOrder(t *testing.T) {
	order := &Order{}

	err := order.discountOrder(15)
	if err != nil {
		t.Errorf("discountOrder() error = %v)", err)
	}
	if order.Discount != 15 {
		t.Errorf("Очікувалась знижка 15, отримано %.2f", order.Discount)
	}

	err = order.discountOrder(0)
	if err == nil {
		t.Error("Очікувалась помилка для знижки < 1%, але її немає")
	}

	err = order.discountOrder(101)
	if err == nil {
		t.Error("Очікувалась помилка для знижки > 100%, але її немає")
	}
}
