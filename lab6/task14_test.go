package main

import (
	"testing"
)

func TestAddDevice(t *testing.T) {
	machine := &Machine{}

	err := machine.addDevice(Engine{model: "V8", number: 100, price: 5000, horsepower: 400})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
	}

	tests := []struct {
		name    string
		device  Device
		isError bool
	}{
		{"Успішне додавання насоса", Pump{model: "WaterX", number: 101, price: 1500, capacity: 120.5}, false},
		{"Успішне додавання турбіни", Turbine{model: "T-Max", number: 102, price: 8000, rpm: 3000}, false},
		{"Некоректний номер (0)", Engine{model: "V6", number: 0, price: 3000, horsepower: 300}, true},
		{"Від'ємна ціна", Pump{model: "WaterY", number: 103, price: -500, capacity: 50}, true},
		{"Дублікат інвентарного номера", Turbine{model: "T-Mini", number: 100, price: 4000, rpm: 1500}, true},
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

	err := machine.addDevice(Engine{model: "V8", number: 101, price: 1000, horsepower: 300})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
	}

	err = machine.addDevice(Pump{model: "P1", number: 102, price: 500, capacity: 100})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
	}

	tests := []struct {
		name    string
		number  int
		isError bool
	}{
		{"Успішне видалення існуючого пристрою", 101, false},
		{"Видалення неіснуючого пристрою", 999, true},
		{"Некоректний номер (від'ємний)", -5, true},
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

	err := machine.addDevice(Turbine{model: "T-Max", number: 200, price: 5000, rpm: 4000})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
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

	err := machine.addDevice(Pump{model: "P-100", number: 55, price: 300, capacity: 50})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
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

	if got := machine.totalPriceOfDevices(); got != 0 {
		t.Errorf("Очікувалась вартість 0 для порожньої системи, отримано %.2f", got)
	}

	err := machine.addDevice(Engine{model: "D1", number: 1, price: 1000.50, horsepower: 200})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
	}

	err = machine.addDevice(Pump{model: "P1", number: 2, price: 500.25, capacity: 60})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
	}

	err = machine.addDevice(Turbine{model: "T1", number: 3, price: 2000.00, rpm: 1500})
	if err != nil {
		t.Errorf("addDevice() error = %v", err)
	}

	expected := 1000.50 + 500.25 + 2000.00
	if got := machine.totalPriceOfDevices(); got != expected {
		t.Errorf("Очікувалась вартість %.2f, отримано %.2f", expected, got)
	}

	err = machine.deleteDevice(2)
	if err != nil {
		t.Errorf("deleteDevice() error = %v", err)
	}

	expectedAfterDelete := 1000.50 + 2000.00
	if got := machine.totalPriceOfDevices(); got != expectedAfterDelete {
		t.Errorf("Очікувалась вартість після видалення %.2f, отримано %.2f", expectedAfterDelete, got)
	}
}
