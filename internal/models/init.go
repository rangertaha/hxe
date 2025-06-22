package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Schema struct {
	Properties []Property `json:"properties"`
}

type Property struct {
	Type        string   `json:"type"`
	Name        string   `json:"name"`
	Label       string   `json:"label"`
	Description string   `json:"desc"`
	Default     string   `json:"default"`
	Required    bool     `json:"required"`
	Options     []string `json:"options"`
}

// AutoMigrate auto migrates the database
func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&Program{},
	)
}
