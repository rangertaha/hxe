package models

import (
	"gorm.io/gorm"
)

// Group represents a hierarchical group resource
type Group struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `gorm:"not null" json:"title"`
	Icon        string  `json:"icon"`
	Description string  `json:"description"`
	ParentID    *uint   `json:"parent_id"`
	Parent      *Group  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children    []Group `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	gorm.Model
}

// Service represents a service resource
type Service struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	gorm.Model
}

// Tag represents a tag resource
type Tag struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"not null" json:"title"`
	Description string `json:"description"`
	Values      string `json:"values"`
	Query       string `json:"query"`
	Weight      int    `json:"weight"`
	gorm.Model
}

// Field represents a field resource
type Field struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Type        string `gorm:"not null" json:"type"`
	Description string `json:"description"`
	Weight      int    `json:"weight"`
	gorm.Model
}

// Variable represents a variable resource
type Variable struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Key   string `gorm:"uniqueIndex;not null" json:"key"`
	Value string `json:"value"`
	gorm.Model
}
