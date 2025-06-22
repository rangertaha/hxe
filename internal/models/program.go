package models

import (
	"github.com/rangertaha/hxe/internal/log"

	"gorm.io/gorm"
)

type ProgramStatus string

const (
	ProgramStopped ProgramStatus = "stopped"
	ProgramRunning ProgramStatus = "running"
	ProgramStarted ProgramStatus = "started"
	ProgramFailed  ProgramStatus = "failed"
)

type Program struct {
	gorm.Model
	ID          uint          `gorm:"primaryKey" json:"id"`
	Name        string        `json:"name"`
	Description string        `json:"desc"`
	Command     string        `json:"command"`
	Args        string        `json:"args"`
	Directory   string        `json:"cwd"`
	User        string        `json:"user"`
	Group       string        `json:"group"`
	Status      ProgramStatus `json:"status"`
	PID         int           `json:"pid"`
	ExitCode    int           `json:"exitCode"`
	StartTime   int64         `json:"startTime"`
	EndTime     int64         `json:"endTime"`
	Autostart   bool          `json:"autostart"`
	Enabled     bool          `json:"enabled"`
	Retries     int           `json:"retries"`
	MaxRetries  int           `json:"maxRetries"`
}

func (program *Program) AfterSave(tx *gorm.DB) (err error) {
	// Implement your post-save logic here
	log.Info().Str("program", program.Name).Str("status", string(program.Status)).Str("command", program.Command).Msgf("Program successfully saved!")
	return nil
}
