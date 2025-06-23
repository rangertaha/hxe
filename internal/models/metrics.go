package models

import (
	"github.com/rangertaha/hxe/internal/log"

	"gorm.io/gorm"
)

type Metric struct {
	ID        uint               `gorm:"primaryKey" json:"id"`
	Timestamp int64              `json:"timestamp"`
	Tags      map[string]string  `json:"tags"`
	Fields    map[string]float64 `json:"fields"`
}

func (program *Program) AfterSave(tx *gorm.DB) (err error) {
	// Implement your post-save logic here
	log.Info().Str("program", program.Name).Str("status", string(program.Status)).Str("command", program.Command).Msgf("Program successfully saved!")
	return nil
}
