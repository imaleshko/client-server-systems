package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Device interface {
	getClass() string
	getModel() string
	getNumber() int
	getPrice() float64
	PrintInfo()
}
type Engine struct {
	model      string
	number     int
	price      float64
	horsepower int
}

func (e Engine) getClass() string  { return "Двигун" }
func (e Engine) getModel() string  { return e.model }
func (e Engine) getNumber() int    { return e.number }
func (e Engine) getPrice() float64 { return e.price }
func (e Engine) PrintInfo() {
	fmt.Printf("Двигун моделі %s: номер %d, ціна: %.2f, потужність: %d к.с.\n", e.model, e.number, e.price, e.horsepower)
}

type Pump struct {
	model    string
	number   int
	price    float64
	capacity float64
}

func (p Pump) getClass() string  { return "Насос" }
func (p Pump) getModel() string  { return p.model }
func (p Pump) getNumber() int    { return p.number }
func (p Pump) getPrice() float64 { return p.price }
func (p Pump) PrintInfo() {
	fmt.Printf("Насос моделі %s: номер %d, ціна: %.2f, продуктивність: %.1f л/хв\n", p.model, p.number, p.price, p.capacity)
}

type Turbine struct {
	model  string
	number int
	price  float64
	rpm    int
}

func (t Turbine) getClass() string  { return "Турбіна" }
func (t Turbine) getModel() string  { return t.model }
func (t Turbine) getNumber() int    { return t.number }
func (t Turbine) getPrice() float64 { return t.price }
func (t Turbine) PrintInfo() {
	fmt.Printf("Турбіна моделі %s: номер %d, ціна: %.2f, обертальна швидкість: %d об/хв\n", t.model, t.number, t.price, t.rpm)
}

type Machine struct {
	devices []Device
}

func (m *Machine) addDevice(device Device) error {
	if device.getNumber() <= 0 {
		return errors.New("номер повинен бути додатним числом")
	}
	if device.getPrice() <= 0 {
		return errors.New("ціна повинна бути додатним числом")
	}

	if _, err := m.searchDeviceByNumber(device.getNumber()); err == nil {
		return errors.New("пристрій з таким номером вже існує")
	}

	m.devices = append(m.devices, device)
	return nil
}

func (m *Machine) deleteDevice(number int) error {
	if number <= 0 {
		return errors.New("номер повинен бути додатним числом")
	}

	for i, device := range m.devices {
		if device.getNumber() == number {
			m.devices = append(m.devices[:i], m.devices[i+1:]...)
			return nil
		}
	}

	return errors.New("пристрій не знайдено")
}

func (m *Machine) searchDeviceByModel(class, model string) (Device, error) {
	if class == "" {
		return nil, errors.New("тип продукту не може бути порожнім")
	}
	if model == "" {
		return nil, errors.New("назва продукту не може бути порожньою")
	}

	for _, device := range m.devices {
		if device.getClass() == class && device.getModel() == model {
			return device, nil
		}
	}

	return nil, errors.New("пристрій не знайдено")
}

func (m *Machine) searchDeviceByNumber(number int) (Device, error) {
	if number <= 0 {
		return nil, errors.New("номер повинен бути додатним числом")
	}

	for _, device := range m.devices {
		if device.getNumber() == number {
			return device, nil
		}
	}

	return nil, errors.New("пристрій не знайдено")
}

func (m *Machine) totalPriceOfDevices() float64 {
	totalPrice := 0.0
	for _, device := range m.devices {
		totalPrice += device.getPrice()
	}
	return totalPrice
}

func (m *Machine) printMachine() {
	fmt.Println("Стан системи:")

	if len(m.devices) == 0 {
		fmt.Println("Система порожня")
		return
	}

	fmt.Printf("Сума вартості пристроїв: %.2f\n", m.totalPriceOfDevices())

	for _, p := range m.devices {
		p.PrintInfo()
	}
}

func task14() {
	scanner := bufio.NewScanner(os.Stdin)
	var machine Machine
	for {
		fmt.Println("\n1 - надрукувати стан системи")
		fmt.Println("2 - додати пристрій")
		fmt.Println("3 - видалити пристрій")
		fmt.Println("4 - знайти пристрій за моделлю")
		fmt.Println("5 - знайти пристрій за номером")
		fmt.Println("6 - переглянути вартість всіх пристроїв")
		fmt.Println("7 - вийти")

		switch readLine(scanner, "Виберіть опцію:") {
		case "1":
			machine.printMachine()
		case "2":
			class := readLine(scanner, "Введіть тип пристрою (Двигун, Насос, Турбіна):")
			model := readLine(scanner, "Введіть модель пристрою:")

			numberStr := readLine(scanner, "Введіть номер пристрою:")
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				fmt.Println("Помилка: номер пристрою повинен бути числом")
				continue
			}

			priceStr := readLine(scanner, "Введіть ціну товару:")
			price, err := strconv.ParseFloat(priceStr, 64)
			if err != nil {
				fmt.Println("Помилка: ціна повинна бути числом")
				continue
			}

			var newDevice Device
			switch class {
			case "Двигун":
				hpStr := readLine(scanner, "Введіть потужність двигуна (к.с.):")
				hp, err := strconv.Atoi(hpStr)
				if err != nil || hp <= 0 {
					fmt.Println("Помилка: потужність повинна бути додатним числом")
					continue
				}
				newDevice = Engine{model: model, number: number, price: price, horsepower: hp}
			case "Насос":
				capStr := readLine(scanner, "Введіть продуктивність насоса (л/хв):")
				capacity, err := strconv.ParseFloat(capStr, 64)
				if err != nil || capacity <= 0 {
					fmt.Println("Помилка: продуктивність повинна бути додатним числом")
					continue
				}
				newDevice = Pump{model: model, number: number, price: price, capacity: capacity}
			case "Турбіна":
				rpmStr := readLine(scanner, "Введіть обертальну швидкість турбіни (об/хв):")
				rpm, err := strconv.Atoi(rpmStr)
				if err != nil || rpm <= 0 {
					fmt.Println("Помилка: швидкість повинна бути додатним числом")
					continue
				}
				newDevice = Turbine{model: model, number: number, price: price, rpm: rpm}
			default:
				fmt.Println("Помилка: невідомий тип пристрою. Доступні: Двигун, Насос, Турбіна")
				continue
			}

			err = machine.addDevice(newDevice)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
			}
		case "3":
			numberStr := readLine(scanner, "Введіть номер пристрою:")
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				fmt.Println("Помилка: номер пристрою повинен бути числом")
				continue
			}

			err = machine.deleteDevice(number)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
			}
		case "4":
			class := readLine(scanner, "Введіть тип пристрою:")
			model := readLine(scanner, "Введіть модель пристрою:")

			device, err := machine.searchDeviceByModel(class, model)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
				continue
			}

			fmt.Println("Знайдений пристрій:")
			device.PrintInfo()
		case "5":
			numberStr := readLine(scanner, "Введіть номер пристрою:")
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				fmt.Println("Помилка: номер пристрою повинен бути числом")
				continue
			}

			device, err := machine.searchDeviceByNumber(number)
			if err != nil {
				fmt.Printf("Помилка: %v\n", err)
				continue
			}

			fmt.Println("Знайдений пристрій:")
			device.PrintInfo()
		case "6":
			fmt.Println("Сума вартості пристроїв:", machine.totalPriceOfDevices())
		case "7":
			return
		default:
			fmt.Println("Помилка: виберіть опцію серед наведених")
		}
	}
}
