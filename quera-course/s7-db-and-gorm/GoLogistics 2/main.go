package main

import "fmt"

func main() {
	db := GetDB("gopher.db")
	AutoMigrateModels(db)

	p := &Package{
		TrackingCode: "IR-1002",
		Description:  "Laptop",
		Weight:       2.5,
		Status:       "Pending",
	}

	CreatePackage(db, p)
	fmt.Println("Package created.")

	fetched, _ := GetPackageByTracking(db, "IR-1002")
	fmt.Printf("Fetched Status: %s\n", fetched.Status)

	UpdatePackageStatus(db, "IR-1002", "Delivered")
	fetched2, _ := GetPackageByTracking(db, "IR-1002")
	fmt.Printf("Updated Status: %s\n", fetched2.Status)

	err := DeletePackage(db, "IR-1002")
	if err == nil {
		fmt.Println("Package deleted successfully.")
	}
}
