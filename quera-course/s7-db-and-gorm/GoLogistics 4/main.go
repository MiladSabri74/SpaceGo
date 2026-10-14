package main

import "fmt"

func main() {
	db := GetDB("gopher.db")
	AutoMigrateModels(db)

	c1 := &Courier{Name: "Peyk 1"}
	c2 := &Courier{Name: "Peyk 2"}
	db.Create(c1)
	db.Create(c2)

	p := &Package{TrackingCode: "URGENT1", CourierID: &c1.ID}
	db.Create(p)

	err := TransferPackage(db, "URGENT1", c2.ID)
	if err == nil {
		fmt.Println("Package transferred successfully.")
	}

	err = TransferPackage(db, "URGENT1", 999)
	if err != nil {
		fmt.Println("Transfer failed:", err)
	}
}
