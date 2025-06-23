/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 rangertaha@gmail.com
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

package models

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

type Base struct {
	ID      uint           `gorm:"primaryKey" json:"id"`
	Created int64          `gorm:"autoCreateTime" json:"created"`
	Updated int64          `gorm:"autoUpdateTime" json:"updated"`
	Deleted gorm.DeletedAt `gorm:"index" json:"deleted"`
}

type App struct {
	Base
	ID       string    `gorm:"primaryKey" json:"id"`
	Programs []Program `json:"programs" gorm:"foreignKey:AID"`
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
		&Program{},
	)
}

func ProgramSchema() *Schema {

	return &Schema{
		Properties: []Property{
			{
				Name:        "name",
				Label:       "Name",
				Description: "The name of the program",
				Type:        "string",
				Default:     "",
				Required:    true,
				Options:     []string{},
			},
			{
				Name:        "desc",
				Label:       "Description",
				Description: "The description of the program",
				Type:        "string",
				Default:     "",
				Required:    true,
				Options:     []string{},
			},
			{
				Name:        "cmd",
				Label:       "Command",
				Description: "The command to run the program",
				Type:        "string",
				Default:     "",
				Required:    true,
				Options:     []string{},
			},
			{
				Name:        "args",
				Label:       "Arguments",
				Description: "The arguments to pass to the program",
				Type:        "string",
				Default:     "",
				Required:    false,
				Options:     []string{},
			},
			{
				Name:        "cwd",
				Label:       "Working Directory",
				Description: "The working directory of the program",
				Type:        "string",
				Default:     "",
				Required:    true,
				Options:     []string{},
			},
			{
				Name:        "user",
				Label:       "User",
				Description: "The user to run the program as",
				Type:        "string",
				Default:     "",
				Required:    false,
				Options:     []string{},
			},
			{
				Name:        "group",
				Label:       "Group",
				Description: "The group to run the program as",
				Type:        "string",
				Default:     "",
				Required:    false,
				Options:     []string{},
			},
		},
	}
}
