package config

import (
	"os"

	"github.com/laracarvalho/goboards/schema"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartDB() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./main.db"
	// Check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating...")

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		
		file.Close()
	}

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}
	// Migrate the Schema
	err = db.AutoMigrate(&schema.Listing{})
	if err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}
	// Return the DB
	return db, nil
}