package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(connectionString string) error {

	db, err := gorm.Open(
		postgres.Open(connectionString),
		&gorm.Config{},
	)

	if err != nil {
		return err
	}

	DB = db

	return nil
}
