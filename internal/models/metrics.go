/*
Copyright © 2025 Rangertaha <rangertaha@gmail.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/

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
