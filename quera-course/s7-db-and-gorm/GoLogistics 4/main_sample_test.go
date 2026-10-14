package main

import (
	"testing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSampleTransfer(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test DB: %v", err)
	}
	AutoMigrateModels(testDB)

	c1 := &Courier{Name: "Courier A"}
	c2 := &Courier{Name: "Courier B"}
	testDB.Create(c1)
	testDB.Create(c2)

	p := &Package{TrackingCode: "SAMPLE_TRK", CourierID: &c1.ID}
	testDB.Create(p)

	err = TransferPackage(testDB, "SAMPLE_TRK", c2.ID)
	if err != nil {
		t.Fatalf("TransferPackage returned error: %v", err)
	}
}
