package models

import (
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	// Get user's application directory
	appDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}

	// Create hxe directory if it doesn't exist
	hxeDir := filepath.Join(appDir, "hxe")
	if err := os.MkdirAll(hxeDir, 0755); err != nil {
		log.Fatal(err)
	}

	// Set database path in hxe directory
	dbPath := filepath.Join(hxeDir, "hxe.db")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Auto migrate the schema
	db.AutoMigrate(
		// Plugins
		&Program{},

	)
	DB = db
	SeedAll(DB)
}
