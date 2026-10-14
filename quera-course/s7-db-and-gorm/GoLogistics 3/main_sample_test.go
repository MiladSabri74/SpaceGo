package main

import (
	"testing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSampleRelationships(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test DB: %v", err)
	}
	AutoMigrateModels(testDB)

	c := &Courier{Name: "Gopher Express"}
	testDB.Create(c)

	p := &Package{TrackingCode: "TRK_SAMPLE"}
	testDB.Create(p)

	err = AssignPackageToCourier(testDB, "TRK_SAMPLE", c.ID)
	if err != nil {
		t.Fatalf("AssignPackageToCourier returned error: %v", err)
	}
}
