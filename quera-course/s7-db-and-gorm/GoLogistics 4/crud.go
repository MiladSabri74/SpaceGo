package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func CreatePackage(db *gorm.DB, p *Package) error {
	// TODO: Implement Create
	if p == nil {
		return errors.New("package cannot be nil")
	}
	return db.Create(p).Error
}

func GetPackageByTracking(db *gorm.DB, code string) (*Package, error) {
	// TODO: Implement Read
	if code == "" {
		return nil, errors.New("Tracking code can not be empty")
	}
	var p Package
	result := db.Where("tracking_code = ?", code).First(&p)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("package with tracking code %s not found", code)
		}
		return nil, result.Error
	}

	return &p, nil
}

func UpdatePackageStatus(db *gorm.DB, trackingCode string, newStatus string) error {
	// TODO: Implement Update
	if trackingCode == "" {
		return errors.New("Tracking code can not be empty")
	}

	result := db.Model(&Package{}).
		Where("tracking_code = ?", trackingCode).
		Update("status", newStatus)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func DeletePackage(db *gorm.DB, trackingCode string) error {
	// TODO: Implement Delete
	if trackingCode == "" {
		return errors.New("Tracking code can not be empty")
	}

	result := db.Where("tracking_code = ?", trackingCode).
		Delete(&Package{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
