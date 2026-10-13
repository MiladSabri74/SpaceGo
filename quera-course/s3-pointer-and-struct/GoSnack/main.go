package main

type VendingMachine struct {
	invetory map[string]int
	revenue  int
	isOn     bool
}

func NewVendingMachine() *VendingMachine {
	vm := &VendingMachine{
		invetory: make(map[string]int),
		revenue:  0,
		isOn:     false,
	}
	return vm
}

func (vm *VendingMachine) TogglePower() {
	vm.isOn = !vm.isOn
}
func (vm *VendingMachine) Restock(product string, count int) {
	if count <= 0 {
		return
	}
	vm.invetory[product] += count
}

func (vm *VendingMachine) Sell(product string, price int) string {
	if vm.isOn == false {
		return "Machine is OFF"
	}
	if vm.invetory[product] == 0 {
		return "Out of stock"
	}
	vm.invetory[product] -= 1
	vm.revenue += price
	return "Success"
}

func (vm *VendingMachine) TotalRevenue() int {
	return vm.revenue
}

func (vm *VendingMachine) StockCount(product string) int {
	return vm.invetory[product]
}
