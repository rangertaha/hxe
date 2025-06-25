/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 Rangertaha <rangertaha@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */

package db

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type App struct {
	ID      uint           `json:"id" gorm:"primaryKey"`
	Created int64          `json:"created" gorm:"autoCreateTime"`
	Updated int64          `json:"updated" gorm:"autoUpdateTime"`
	Deleted gorm.DeletedAt `json:"deleted" gorm:"index"`

	// Relations
	Services []Service `json:"services" gorm:"foreignKey:AppID"`
}

type Schema struct {
	Name       string     `json:"name"`
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
		&App{},
		&Service{},
	)
}
