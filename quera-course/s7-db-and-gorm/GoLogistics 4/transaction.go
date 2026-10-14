package main

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func TransferPackage(db *gorm.DB, trackingCode string, newCourierID uint) error {
	// TODO: Implement safe transfer using db.Transaction
	if trackingCode == "" {
		return errors.New("tracking code cannot be empty")
	}

	return db.Transaction(func(tx *gorm.DB) error {
		var p Package
		result := tx.Where("tracking_code = ?", trackingCode).First(&p)

		if result.Error != nil {
			return result.Error
		}
		var c Courier
		result = tx.Where("id = ?", newCourierID).First(&c)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return errors.New("courier not found")
			}
		}

		result = tx.Model(p).Update("courier_id", newCourierID)
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return fmt.Errorf("package was not updated")
		}
		return nil
	})
}
