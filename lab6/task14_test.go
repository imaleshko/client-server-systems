package main

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestAddDevice(t *testing.T) {
	machine := &Machine{}

	err := machine.addDevice(Engine{model: "V8", number: 100, price: decimal.NewFromInt(5000), horsepower: 400})
	if err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	tests := []struct {
		name    string
		device  Device
		isError bool
	}{
		{"Успішне додавання насоса", Pump{model: "WaterX", number: 101, price: decimal.NewFromInt(1500), capacity: 120.5}, false},
		{"Успішне додавання турбіни", Turbine{model: "T-Max", number: 102, price: decimal.NewFromInt(8000), rpm: 3000}, false},
		{"Некоректний номер (0)", Engine{model: "V6", number: 0, price: decimal.NewFromInt(3000), horsepower: 300}, true},
		{"Відʼємна ціна", Pump{model: "WaterY", number: 103, price: decimal.NewFromInt(-500), capacity: 50}, true},
		{"Дублікат інвентарного номера", Turbine{model: "T-Mini", number: 100, price: decimal.NewFromInt(4000), rpm: 1500}, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := machine.addDevice(tc.device)

			if (err != nil) != tc.isError {
				t.Errorf("addDevice() error = %v, isError = %v", err, tc.isError)
			}
		})
	}
}

func TestDeleteDevice(t *testing.T) {
	machine := &Machine{}

	if err := machine.addDevice(Engine{model: "V8", number: 101, price: decimal.NewFromInt(1000), horsepower: 300}); err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	if err := machine.addDevice(Pump{model: "P1", number: 102, price: decimal.NewFromInt(500), capacity: 100}); err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	tests := []struct {
		name    string
		number  int
		isError bool
	}{
		{"Успішне видалення існуючого пристрою", 101, false},
		{"Видалення неіснуючого пристрою", 999, true},
		{"Некоректний номер (відʼємний)", -5, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := machine.deleteDevice(tc.number)

			if (err != nil) != tc.isError {
				t.Errorf("deleteDevice() error = %v, isError = %v", err, tc.isError)
			}
		})
	}

	if len(machine.devices) != 1 {
		t.Errorf("Очікувалась кількість пристроїв 1, отримано %d", len(machine.devices))
	}
}

func TestSearchDeviceByModel(t *testing.T) {
	machine := &Machine{}

	if err := machine.addDevice(Turbine{model: "T-Max", number: 200, price: decimal.NewFromInt(5000), rpm: 4000}); err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	tests := []struct {
		name    string
		class   string
		model   string
		isError bool
	}{
		{"Успішний пошук", "Турбіна", "T-Max", false},
		{"Пошук неіснуючої моделі", "Турбіна", "T-Mini", true},
		{"Пошук з порожнім класом", "", "T-Max", true},
		{"Пошук з порожньою моделлю", "Турбіна", "", true},
		{"Пошук іншого класу з такою ж моделлю", "Двигун", "T-Max", true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			device, err := machine.searchDeviceByModel(tc.class, tc.model)

			if (err != nil) != tc.isError {
				t.Errorf("searchDeviceByModel() error = %v, isError = %v", err, tc.isError)
				return
			}
			if !tc.isError && (device.getClass() != tc.class || device.getModel() != tc.model) {
				t.Errorf("Знайдено неправильний пристрій: %+v", device)
			}
		})
	}
}

func TestSearchDeviceByNumber(t *testing.T) {
	machine := &Machine{}

	if err := machine.addDevice(Pump{model: "P-100", number: 55, price: decimal.NewFromInt(300), capacity: 50}); err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	tests := []struct {
		name    string
		number  int
		isError bool
	}{
		{"Успішний пошук", 55, false},
		{"Пристрій не знайдено", 99, true},
		{"Некоректний номер", 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			device, err := machine.searchDeviceByNumber(tc.number)

			if (err != nil) != tc.isError {
				t.Errorf("searchDeviceByNumber() error = %v, isError = %v", err, tc.isError)
				return
			}
			if !tc.isError && device.getNumber() != tc.number {
				t.Errorf("Знайдено неправильний пристрій: %+v", device)
			}
		})
	}
}

func TestTotalPriceOfDevices(t *testing.T) {
	machine := &Machine{}

	if got := machine.totalPriceOfDevices(); !got.IsZero() {
		t.Errorf("Очікувалась вартість 0 для порожньої системи, отримано %s", got.StringFixed(2))
	}

	if err := machine.addDevice(Engine{model: "D1", number: 1, price: decimal.RequireFromString("1000.50"), horsepower: 200}); err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	if err := machine.addDevice(Pump{model: "P1", number: 2, price: decimal.RequireFromString("500.25"), capacity: 60}); err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	if err := machine.addDevice(Turbine{model: "T1", number: 3, price: decimal.RequireFromString("2000.00"), rpm: 1500}); err != nil {
		t.Fatalf("addDevice() error = %v", err)
	}

	expected := decimal.RequireFromString("3500.75")
	if got := machine.totalPriceOfDevices(); !got.Equal(expected) {
		t.Errorf("Очікувалась вартість %s, отримано %s", expected.StringFixed(2), got.StringFixed(2))
	}

	if err := machine.deleteDevice(2); err != nil {
		t.Fatalf("deleteDevice() error = %v", err)
	}

	expectedAfterDelete := decimal.RequireFromString("3000.50")
	if got := machine.totalPriceOfDevices(); !got.Equal(expectedAfterDelete) {
		t.Errorf("Очікувалась вартість після видалення %s, отримано %s", expectedAfterDelete.StringFixed(2), got.StringFixed(2))
	}
}
