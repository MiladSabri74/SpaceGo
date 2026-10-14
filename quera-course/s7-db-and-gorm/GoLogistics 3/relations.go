package main

import (
	"errors"

	"gorm.io/gorm"
)

func AssignPackageToCourier(db *gorm.DB, packageTrackingCode string, courierID uint) error {
	// TODO: Implement assignment
	result := db.Model(&Package{}).
		Where("tracking_code = ?", packageTrackingCode).
		Update("courier_id", &courierID)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("Package not found or can not updated by courierID")
	}
	return nil
}

func GetCourierWithPackages(db *gorm.DB, courierID uint) (*Courier, error) {
	// TODO: Implement Read with Preload
	var courier Courier
	result := db.Preload("Packages").
		Where("id = ?", courierID).First(&courier)

	if result.Error != nil {
		return nil, result.Error
	}
	return &courier, nil
}
