package models

import (
	"time"

	"gorm.io/gorm"
)

// Tag represents a tag resource
type Tag struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:100;not null;uniqueIndex"`
	Color       string         `json:"color" gorm:"size:7;default:'#3B82F6'"` // hex color code
	Description string         `json:"description" gorm:"type:text"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for Tag
func (Tag) TableName() string {
	return "tags"
}

// BeforeCreate is a GORM hook that runs before creating a record
func (t *Tag) BeforeCreate(tx *gorm.DB) error {
	if t.CreatedAt.IsZero() {
		t.CreatedAt = time.Now()
	}
	if t.UpdatedAt.IsZero() {
		t.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating a record
func (t *Tag) BeforeUpdate(tx *gorm.DB) error {
	t.UpdatedAt = time.Now()
	return nil
}
