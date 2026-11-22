package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		productName string
		price       decimal.Decimal
		isError     bool
	}{
		{"Успішне додавання", "Яблуко", decimal.RequireFromString("10.5"), false},
		{"Порожня назва", "", decimal.RequireFromString("10.5"), true},
		{"Нульова ціна", "Банан", decimal.Zero, true},
		{"Відʼємна ціна", "Груша", decimal.RequireFromString("-5.0"), true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			order := &Order{}
			err := order.add(tc.productName, tc.price)

			if (err != nil) != tc.isError {
				t.Errorf("add() error = %v, isError %v", err, tc.isError)
			}

			if !tc.isError && len(order.products) != 1 {
				t.Errorf("Очікувався 1 товар в замовленні, отримано %d", len(order.products))
			}
		})
	}
}

func TestTotalPriceOfProductInOrder(t *testing.T) {
	order := &Order{}
	if err := order.add("Кава", decimal.NewFromInt(50)); err != nil {
		t.Fatalf("add() error = %v", err)
	}
	if err := order.add("Чай", decimal.NewFromInt(30)); err != nil {
		t.Fatalf("add() error = %v", err)
	}
	if err := order.add("Кава", decimal.NewFromInt(50)); err != nil {
		t.Fatalf("add() error = %v", err)
	}

	tests := []struct {
		name        string
		productName string
		wantTotal   decimal.Decimal
		isError     bool
	}{
		{"Кілька товарів одного типу", "Кава", decimal.NewFromInt(100), false},
		{"Один товар такого типу", "Чай", decimal.NewFromInt(30), false},
		{"Товар не знайдено", "Сік", decimal.Zero, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := order.totalPriceOfProductInOrder(tc.productName)

			if (err != nil) != tc.isError {
				t.Errorf("totalPriceOfProductInOrder() error = %v, isError %v", err, tc.isError)
				return
			}
			if !got.Equal(tc.wantTotal) {
				t.Errorf("Очікувалось %s, отримано %s", tc.wantTotal.StringFixed(2), got.StringFixed(2))
			}
		})
	}
}

func TestTotalPriceOfOrder(t *testing.T) {
	order := &Order{}

	if got := order.totalPriceOfOrder(); !got.IsZero() {
		t.Errorf("Очікувалось 0, отримано %s", got.StringFixed(2))
	}

	if err := order.add("Телефон", decimal.NewFromInt(1000)); err != nil {
		t.Fatalf("add() error = %v", err)
	}
	if err := order.add("Чохол", decimal.NewFromInt(200)); err != nil {
		t.Fatalf("add() error = %v", err)
	}

	wantWithoutDiscount := decimal.NewFromInt(1200)
	if got := order.totalPriceOfOrder(); !got.Equal(wantWithoutDiscount) {
		t.Errorf("Очікувалось %s, отримано %s", wantWithoutDiscount.StringFixed(2), got.StringFixed(2))
	}

	if err := order.discountOrder(decimal.NewFromInt(10)); err != nil {
		t.Fatalf("discountOrder() error = %v", err)
	}

	wantWithDiscount := decimal.NewFromInt(1080)
	if got := order.totalPriceOfOrder(); !got.Equal(wantWithDiscount) {
		t.Errorf("Очікувалось %s, отримано %s", wantWithDiscount.StringFixed(2), got.StringFixed(2))
	}
}

func TestDiscountProduct(t *testing.T) {
	tests := []struct {
		name          string
		productName   string
		percentage    decimal.Decimal
		isError       bool
		expectedPrice decimal.Decimal
	}{
		{"Успішна знижка 50%", "Ноутбук", decimal.NewFromInt(50), false, decimal.NewFromInt(500)},
		{"Знижка менше 1%", "Ноутбук", decimal.Zero, true, decimal.NewFromInt(1000)},
		{"Знижка більше 100%", "Ноутбук", decimal.NewFromInt(150), true, decimal.NewFromInt(1000)},
		{"Товар не знайдено", "Мишка", decimal.NewFromInt(10), true, decimal.NewFromInt(1000)},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			order := &Order{}
			if err := order.add("Ноутбук", decimal.NewFromInt(1000)); err != nil {
				t.Fatalf("add() error = %v", err)
			}

			err := order.discountProduct(tc.productName, tc.percentage)

			if (err != nil) != tc.isError {
				t.Errorf("discountProduct() error = %v, isError %v", err, tc.isError)
			}

			if !order.products[0].price.Equal(tc.expectedPrice) {
				t.Errorf("Очікувана ціна %s, отримано %s", tc.expectedPrice.StringFixed(2), order.products[0].price.StringFixed(2))
			}
		})
	}
}

func TestDiscountOrder(t *testing.T) {
	order := &Order{}

	wantDiscount := decimal.NewFromInt(15)
	if err := order.discountOrder(wantDiscount); err != nil {
		t.Fatalf("discountOrder() error = %v", err)
	}
	if !order.discount.Equal(wantDiscount) {
		t.Errorf("Очікувалась знижка %s, отримано %s", wantDiscount, order.discount)
	}

	if err := order.discountOrder(decimal.Zero); err == nil {
		t.Error("Очікувалась помилка для знижки < 1%, але її немає")
	}

	if err := order.discountOrder(decimal.NewFromInt(101)); err == nil {
		t.Error("Очікувалась помилка для знижки > 100%, але її немає")
	}
}
