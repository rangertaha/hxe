package db

import (
	"log"
	"os"
	"path/filepath"

	"github.com/rangertaha/hxe/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init initializes the database connection and runs migrations
func Init() error {
	// Get user config directory
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		return err
	}

	// Create hxe directory if it doesn't exist
	hxeDir := filepath.Join(userConfigDir, "hxe")
	if err := os.MkdirAll(hxeDir, 0755); err != nil {
		return err
	}

	// Database file path
	dbPath := filepath.Join(hxeDir, "hxe.db")

	// Open database connection
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	// Auto migrate models
	if err := db.AutoMigrate(
		&models.Group{},
		&models.Service{},
		&models.Tag{},
		&models.Field{},
		&models.Variable{},
	); err != nil {
		return err
	}

	DB = db
	log.Printf("Database initialized at: %s", dbPath)
	return nil
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
