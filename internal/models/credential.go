package models

import (
	"time"

	"gorm.io/gorm"
)

// Credential represents a credential resource
type Credential struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" gorm:"size:255;not null"`
	Type        string         `json:"type" gorm:"size:50;not null"` // password, key, token, etc.
	Username    string         `json:"username" gorm:"size:255"`
	Password    string         `json:"password" gorm:"size:500"` // encrypted
	PrivateKey  string         `json:"private_key" gorm:"type:text"`
	PublicKey   string         `json:"public_key" gorm:"type:text"`
	Token       string         `json:"token" gorm:"size:1000"`
	Description string         `json:"description" gorm:"type:text"`
	Tags        []Tag          `json:"tags" gorm:"many2many:credential_tags;"`
	GroupID     *uint          `json:"group_id"`
	Group       *Group         `json:"group" gorm:"foreignKey:GroupID"`
	IsActive    bool           `json:"is_active" gorm:"default:true"`
	ExpiresAt   *time.Time     `json:"expires_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// TableName specifies the table name for Credential
func (Credential) TableName() string {
	return "credentials"
}

// BeforeCreate is a GORM hook that runs before creating a record
func (c *Credential) BeforeCreate(tx *gorm.DB) error {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	if c.UpdatedAt.IsZero() {
		c.UpdatedAt = time.Now()
	}
	return nil
}

// BeforeUpdate is a GORM hook that runs before updating a record
func (c *Credential) BeforeUpdate(tx *gorm.DB) error {
	c.UpdatedAt = time.Now()
	return nil
}
