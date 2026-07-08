package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	units := make(map[string]int)
	units["quarter_of_a_dozen"] = 3
	units["half_of_a_dozen"] = 6
	units["dozen"] = 12
	units["small_gross"] = 120
	units["gross"] = 144
	units["great_gross"] = 1728
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	bill := make(map[string]int)
	return bill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	amount, exist := units[unit]
	if exist {
		bill[item] += amount
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	amount, exist_unit := units[unit]
	bill_amount, exist_item := bill[item]

	if exist_item && exist_unit {
		if bill_amount-amount < 0 {
			return false
		}
		bill[item] -= amount
		if bill[item] == 0 {
			delete(bill, item)
		}
		return true
	}
	return false

}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	amount, exist := bill[item]
	return amount, exist
}
