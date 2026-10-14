package main

import "gorm.io/gorm"

// TODO: Courier struct

type Courier struct {
	gorm.Model
	Name string
}

// TODO: Package Struct
type Package struct {
	gorm.Model
	TrackingCode string `gorm:"uniqueIndex"`
	Description  string
	Weight       float64
	Status       string
}
