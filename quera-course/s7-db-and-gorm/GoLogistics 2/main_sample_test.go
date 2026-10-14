package main

import (
	"testing"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestSampleCreate(t *testing.T) {
	testDB, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test DB: %v", err)
	}
	AutoMigrateModels(testDB)

	p := &Package{
		TrackingCode: "SAMPLE1",
		Status:       "Pending",
	}

	err = CreatePackage(testDB, p)
	if err != nil {
		t.Fatalf("CreatePackage returned error: %v", err)
	}

	var count int64
	testDB.Model(&Package{}).Where("tracking_code = ?", "SAMPLE1").Count(&count)
	if count != 1 {
		t.Errorf("Expected 1 package, found %d", count)
	}
}
