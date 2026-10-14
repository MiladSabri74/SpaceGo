package main

import (
	"gorm.io/gorm"
)

type Courier struct {
	// TODO: Add Packages relationship
	gorm.Model
	Name     string
	Packages []Package
}

// TODO: Package Struct
type Package struct {
	// TODO: Add CourierID foreign key
	gorm.Model
	TrackingCode string `gorm:"uniqueIndex"`
	CourierID    *uint
	Description  string
	Weight       float64
	Status       string
}
