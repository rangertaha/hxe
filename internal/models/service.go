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

package models

type ServiceStatus string

const (
	ServiceStop    ServiceStatus = "stop"
	ServiceStart   ServiceStatus = "start"
	ServiceReload  ServiceStatus = "reload"
	ServiceRestart ServiceStatus = "restart"
)

type Group struct {
	Base
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`

	Services []Service `json:"programs" gorm:"foreignKey:GID"`
}

type Action struct {
	Name  string     `json:"name"`
	Desc  string     `json:"desc"`
	Props []Property `json:"props"`
}

type Service struct {
	Base
	AID         uint          `json:"aid" gorm:"column:aid"`
	GID         uint          `json:"gid" gorm:"column:gid"`
	SID         uint          `json:"sid" gorm:"column:sid"`
	Name        string        `json:"name" gorm:"column:name"`
	Description string        `json:"desc" gorm:"column:desc"`
	Command     string        `json:"command" gorm:"column:command"`
	Args        string        `json:"args" gorm:"column:args"`
	Directory   string        `json:"cwd"`
	User        string        `json:"user"`
	Group       string        `json:"group"`
	Status      ServiceStatus `json:"status"`
	PID         int           `json:"pid" gorm:"column:pid"`
	ExitCode    int           `json:"exitCode" gorm:"column:exit"`
	StartTime   int64         `json:"startTime" gorm:"column:start_time"`
	EndTime     int64         `json:"endTime" gorm:"column:end_time"`
	Autostart   bool          `json:"autostart"`
	Enabled     bool          `json:"enabled"`
	Retries     int           `json:"retries"`
	MaxRetries  int           `json:"maxRetries"`

	// Metrics map[string]float64 `json:"metrics" gorm:"serialize:json"`
	// Actions []Action           `json:"actions" gorm:"serialize:json"`
}

// func (p *Service) AfterSave(tx *gorm.DB) (err error) {
// 	// Implement your post-save logic here
// 	log.Info().Str("program", p.Name).Str("status", string(p.Status)).Str("command", p.Command).Msgf("Service successfully saved!")
// 	return nil
// }
