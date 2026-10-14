package main

import "fmt"

func main() {
	db := GetDB("gopher.db")
	AutoMigrateModels(db)

	c := &Courier{Name: "Ali"}
	db.Create(c)

	p1 := &Package{TrackingCode: "PKG1"}
	p2 := &Package{TrackingCode: "PKG2"}
	db.Create(p1)
	db.Create(p2)

	AssignPackageToCourier(db, "PKG1", c.ID)
	AssignPackageToCourier(db, "PKG2", c.ID)
	fmt.Println("Packages assigned.")

	courier, _ := GetCourierWithPackages(db, c.ID)
	fmt.Printf("Courier %s has %d packages.\n", courier.Name, len(courier.Packages))
}
