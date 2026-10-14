package main

import (
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func GetDB(dsn string) *gorm.DB {
	// TODO: Initialize a singleton SQLite database connection using sync.Once.
	once.Do(func() {
		var err error
		db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
	})
	return db
}

func AutoMigrateModels(database *gorm.DB) error {
	// TODO: Auto-migrate the Courier and Package models.
	return database.AutoMigrate(&Courier{}, &Package{})
}
