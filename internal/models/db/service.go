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
	"time"

	"gorm.io/gorm"
)

type ServiceStatus int

const (
	StateReady ServiceStatus = iota
	StateLoading
	StateStarting
	StateRestaring
	StateRunning
	StateStopping
	StateStopped
	StateFailed
	StateSuccess
	StateUnknown
)

func (s ServiceStatus) String() string {
	switch s {
	case StateReady:
		return "Ready"
	case StateLoading:
		return "Loading"
	case StateStarting:
		return "Starting"
	case StateRestaring:
		return "Restarting"
	case StateRunning:
		return "Running"
	case StateStopping:
		return "Stopping"
	case StateStopped:
		return "Stopped"
	case StateFailed:
		return "Failed"
	case StateSuccess:
		return "Success"
	case StateUnknown:
		return "Unknown"
	default:
		return "Unknown"
	}
}

type Category struct {
	ID      uint           `json:"id" gorm:"primaryKey"`
	Created int64          `json:"created" gorm:"autoCreateTime"`
	Updated int64          `json:"updated" gorm:"autoUpdateTime"`
	Deleted gorm.DeletedAt `json:"deleted" gorm:"index"`
	Name    string         `json:"name"`
	Title   string         `json:"title"`
	Desc    string         `json:"desc"`

	// Relations
	Services []Service `json:"services" gorm:"foreignKey:CategoryID"`
}

type Service struct {
	ID      uint           `json:"id" gorm:"primaryKey"`
	Created int64          `json:"created" gorm:"autoCreateTime"`
	Updated int64          `json:"updated" gorm:"autoUpdateTime"`
	Deleted gorm.DeletedAt `json:"deleted" gorm:"index"`

	// Basic Info
	Name        string `json:"name" gorm:"column:name"`
	Description string `json:"desc" gorm:"column:description"`

	// Runtime configurations
	User      int    `json:"user" gorm:"column:user"`
	Group     int    `json:"group" gorm:"column:group"`
	Directory string `json:"cwd" gorm:"column:dir"`
	PreExec   string `json:"preExec" gorm:"column:preExec"`
	CmdExec   string `json:"cmdExec" gorm:"column:cmdExec"`
	PostExec  string `json:"postExec" gorm:"column:postExec"`
	Autostart bool   `json:"autostart"`
	Retries   int    `json:"retries"`
	Enabled   bool   `json:"enabled"`

	// Runtime Status
	PID       int                `json:"pid" gorm:"column:pid"`
	ExitCode  int                `json:"exitCode" gorm:"column:exit"`
	StartTime int64              `json:"startTime" gorm:"column:started"`
	EndTime   int64              `json:"endTime" gorm:"column:ended"`
	Uptime    time.Duration      `json:"uptime" gorm:"column:uptime"`
	Status    ServiceStatus      `json:"status"`
	Metrics   map[string]float64 `json:"metrics" gorm:"serialize:json"`

	// Relations
	CategoryID uint     `json:"-" gorm:"column:cid"`
	Category   Category `json:"category"`
	AppID      uint     `json:"-" gorm:"column:aid"`
	App        App      `json:"app"`
	Actions    []Action `json:"actions" gorm:"serialize:json"`
}

type Action struct {
	Icon    string `json:"icon"`
	Name    string `json:"name"`
	Label   string `json:"label"`
	Tooltip string `json:"tooltip"`
}

// func (p *Service) AfterSave(tx *gorm.DB) (err error) {
// 	// Implement your post-save logic here
// 	log.Info().Str("program", p.Name).Str("status", string(p.Status)).Str("command", p.Command).Msgf("Service successfully saved!")
// 	return nil
// }

func ServiceSchema() *Schema {

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
