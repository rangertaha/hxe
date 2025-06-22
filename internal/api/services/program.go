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

	// Publish program created event
	if err := p.broker.Publish(internal.PROGRAM_CREATED_TOPIC, []byte(program.ID)); err != nil {
		return nil, fmt.Errorf("failed to publish program created event: %w", err)
	}

	return &program, nil
}

// Update updates a program
func (p *Program) Update(id string, updates models.Program) (*models.Program, error) {
	programID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid program ID: %w", err)
	}

	updates.ID = uint(programID)
	if err := p.db.Save(&updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update program: %w", err)
	}
	return &updates, nil
}

// Delete deletes a program
func (p *Program) Delete(id string) error {
	programID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return fmt.Errorf("invalid program ID: %w", err)
	}

	if err := p.db.Delete(&models.Program{}, programID).Error; err != nil {
		return fmt.Errorf("failed to delete program: %w", err)
	}
	return nil
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
	program.Status = models.ProgramStarted


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
	program.Status = "stopped"
	program.PID = 0
	program.EndTime = 1234567890 // Mock end time
	program.ExitCode = 0

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

	// In a real implementation, you might check the actual process status
	// For now, we just return the stored status
	_ = program // Use the variable to avoid warning
	return program, nil
}

// Reload reloads the configuration for a program
func (p *Program) Reload(id string) error {
	program, err := p.Get(id)
	if err != nil {
		return err
	}

	// In a real implementation, you would reload the program's configuration
	// For now, we just log that reload was requested
	_ = program // Use the variable to avoid warning
	return nil
}

// Enable enables a program to start automatically
func (p *Program) Enable(id string) error {
	program, err := p.Get(id)
	if err != nil {
		return err
	}

	program.Enabled = true
	if err := p.db.Save(program).Error; err != nil {
		return fmt.Errorf("failed to enable program: %w", err)
	}

	return nil
}

// Disable disables a program from starting automatically
func (p *Program) Disable(id string) error {
	program, err := p.Get(id)
	if err != nil {
		return err
	}

	program.Enabled = false
	if err := p.db.Save(program).Error; err != nil {
		return fmt.Errorf("failed to disable program: %w", err)
	}

	return nil
}

// Tail follows the logs of a program
func (p *Program) Tail(id string) error {
	program, err := p.Get(id)
	if err != nil {
		return err
	}

	// In a real implementation, you would start tailing the program's log file
	// For now, we just return success
	_ = program // Use the variable to avoid warning
	return nil
}

// Shell opens an interactive shell for a program
func (p *Program) Shell(id string) error {
	program, err := p.Get(id)
	if err != nil {
		return err
	}

	// In a real implementation, you would open an interactive shell
	// For now, we just return success
	_ = program // Use the variable to avoid warning
	return nil
}

// Run executes a command
func (p *Program) Run(cmd string) error {
	// In a real implementation, you would execute the command
	// For now, we just return success
	return nil
}
