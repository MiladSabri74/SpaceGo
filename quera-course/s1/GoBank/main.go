package main

import "fmt"

func main() {
	var age int = 20
	var income int = 500000000
	var hasBadCredit bool = false

	if age < 18 {
		fmt.Println("Loan rejected: Underage")
	} else if hasBadCredit {
		fmt.Println("Loan rejected: Bad credit history")
	} else if income < 10000000 {
		fmt.Println("Loan rejected: Low income")
	} else {
		fmt.Println("Loan approved")
	}
}
