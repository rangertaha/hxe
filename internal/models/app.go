package models

import (
	"gorm.io/gorm"
)

// type Tag struct {
// 	gorm.Model
// 	Icon   string   `json:"icon"`
// 	Name   string   `json:"name"`
// 	Label  string   `json:"label"`
// 	Tip    string   `json:"tip"`
// 	Query  string   `json:"query"`
// 	Values []string `json:"values"`
// }

type User struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	Password    string `json:"password"`
	Role        string `json:"role"`
	Status      string `json:"status"`

	Projects []*Project `json:"projects" gorm:"many2many:user_projects;"`
	Sessions []Session  `json:"sessions" gorm:"many2many:user_sessions;"`
}

type Session struct {
	gorm.Model
	IP       string `json:"ip"`
	Agent    string `json:"agent"`
	UserID   uint   `json:"user_id"`
	User     User   `json:"user"`
	Active   bool   `json:"active"`
	Verified bool   `json:"verified"`
}

type Project struct {
	gorm.Model
	ID          uint              `gorm:"primaryKey" json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Tags        map[string]string `json:"tags" gorm:"serializer:json"`
}
