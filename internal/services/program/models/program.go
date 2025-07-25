package models

import (
	"gorm.io/gorm"
)

type Program struct {
	ID      uint           `json:"id" gorm:"primaryKey"`
	Created int64          `json:"created" gorm:"autoCreateTime"`
	Updated int64          `json:"updated" gorm:"autoUpdateTime"`
	Deleted gorm.DeletedAt `json:"deleted" gorm:"index"`

	// Basic Info
	Name string `json:"name" gorm:"column:name"`
	Desc string `json:"desc" gorm:"column:description"`

	// Runtime configurations
	Dir   string   `json:"dir" gorm:"column:dir"`
	Path  string   `json:"path" gorm:"column:path"`
	User  string   `json:"user" gorm:"column:user"`
	Group string   `json:"group" gorm:"column:group"`
	Args  []string `json:"args" gorm:"column:args;serializer:json"`
	Env   []string `json:"env" gorm:"column:env;serializer:json"`

	PreExec  string `json:"preExec" gorm:"column:preExec"`
	Exec     string `json:"exec" gorm:"column:cmdExec"`
	PostExec string `json:"postExec" gorm:"column:postExec"`

	Autostart bool `json:"autostart"`
	Retries   int  `json:"retries"`
	Enabled   bool `json:"enabled"`
}
