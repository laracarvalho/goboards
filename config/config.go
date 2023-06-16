package config

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error

	db, err = StartDB()

	if err != nil {
		return fmt.Errorf("Error initializing database: %v", err)
	}

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger := Start(p)
	return logger
}