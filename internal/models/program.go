package models

import "gorm.io/gorm"

type Program struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	Command     string `json:"command"`
	Args        string `json:"args"`
	WorkingDir  string `json:"workingDir"`
	User        string `json:"user"`
	Group       string `json:"group"`
	Status      string `json:"status"`
	PID         int    `json:"pid"`
	ExitCode    int    `json:"exitCode"`
	StartTime   int64  `json:"startTime"`
	EndTime     int64  `json:"endTime"`
	Autostart   bool   `json:"autostart"`
	Enabled     bool   `json:"enabled"`
	Retries     int    `json:"retries"`
	MaxRetries  int    `json:"maxRetries"`
}
