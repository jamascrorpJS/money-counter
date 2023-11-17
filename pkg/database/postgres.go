package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewClient() *gorm.DB {

	fmt.Println("x")
	dsn := "host=localhost port=5432 user=postgres dbname=database password=postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		print(err)
	}
	return db
}
