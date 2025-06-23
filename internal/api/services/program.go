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

package services

import (
	"fmt"
	"strconv"

	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/models"
	"gorm.io/gorm"
)

type Program struct {
	db     *gorm.DB
	broker internal.Broker
}

func NewProgram(b internal.Broker) *Program {
	return &Program{
		db:     models.DB,
		broker: b,
	}
}

// List returns all programs
func (p *Program) List() ([]models.Program, error) {
	var programs []models.Program
	if err := p.db.Find(&programs).Error; err != nil {
		return nil, fmt.Errorf("failed to list programs: %w", err)
	}
	return programs, nil
}

// Get returns a program by ID
func (p *Program) Get(id string) (*models.Program, error) {
	var program models.Program
	programID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid program ID: %w", err)
	}

	if err := p.db.First(&program, programID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("program not found")
		}
		return nil, fmt.Errorf("failed to get program: %w", err)
	}

	return &program, nil
}

// Create creates a new program
func (p *Program) Create(program models.Program) (*models.Program, error) {
	if err := p.db.Create(&program).Error; err != nil {
		return nil, fmt.Errorf("failed to create program: %w", err)
	}
	return &program, nil
}

// Update updates a program
func (p *Program) Update(id string, updates models.Program) (*models.Program, error) {
	if err := p.db.Save(&updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update program: %w", err)
	}
	return &updates, nil
}

// Delete deletes a program
func (p *Program) Delete(id string) (*models.Program, error) {
	deleted, err := p.Stop(id)
	if err != nil {
		return nil, err
	}

	if err := p.db.Delete(&deleted, id).Error; err != nil {
		return nil, fmt.Errorf("failed to delete program: %w", err)
	}

	return deleted, nil
}

// Schema returns the schema of a program
func (p *Program) Schema() (*models.Schema, error) {
	return models.ProgramSchema(), nil
}

// Start starts a program
func (p *Program) Start(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// Update status to started
	program.Status = models.ProgramStart

	if err := p.db.Save(program).Error; err != nil {
		return nil, fmt.Errorf("failed to start program: %w", err)
	}

	return program, nil
}

// Stop stops a program
func (p *Program) Stop(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// Update status to stopped
	program.Status = models.ProgramStop

	if err := p.db.Save(program).Error; err != nil {
		return nil, fmt.Errorf("failed to stop program: %w", err)
	}

	return program, nil
}

// Restart restarts a program
func (p *Program) Restart(id string) (*models.Program, error) {
	// Stop the program first
	if _, err := p.Stop(id); err != nil {
		return nil, err
	}

	// Start the program
	return p.Start(id)
}

// Status returns the status of a program
func (p *Program) Status(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	return program, nil
}

// Reload reloads the configuration for a program
func (p *Program) Reload(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	program.Status = models.ProgramReload

	if err := p.db.Save(program).Error; err != nil {
		return nil, fmt.Errorf("failed to reload program: %w", err)
	}
	return program, nil
}

// Enable enables a program to start automatically
func (p *Program) Enable(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	program.Enabled = true
	if err := p.db.Save(program).Error; err != nil {
		return nil, fmt.Errorf("failed to enable program: %w", err)
	}

	return program, nil
}

// Disable disables a program from starting automatically
func (p *Program) Disable(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	program.Enabled = false
	program.Status = models.ProgramStop
	if err := p.db.Save(program).Error; err != nil {
		return nil, fmt.Errorf("failed to disable program: %w", err)
	}

	return program, nil
}

// Tail follows the logs of a program
func (p *Program) Tail(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// In a real implementation, you would start tailing the program's log file
	// For now, we just return success
	return program, nil
}

// Shell opens an interactive shell for a program
func (p *Program) Shell(id string) (*models.Program, error) {
	program, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// In a real implementation, you would open an interactive shell
	// For now, we just return success
	return program, nil
}
