package main

import "fmt"

func main() {
	db := GetDB("gopher.db")
	if db != nil {
		fmt.Println("Database connection successful!")
	}

	err := AutoMigrateModels(db)
	if err == nil {
		fmt.Println("Models migrated successfully!")
	}
}
