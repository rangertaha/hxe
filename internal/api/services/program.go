package services

import (
	"github.com/rangertaha/hxe/internal/models"
)

type Program struct {
}

func NewProgram() *Program {
	return &Program{}
}

// List returns all programs
func (p *Program) List() ([]models.Program, error) {
	var programs []models.Program

	return programs, nil
}

// Get returns a program by ID
func (p *Program) Get(id string) (*models.Program, error) {
	var program models.Program

	return &program, nil
}

// Create creates a new program
func (p *Program) Create(program models.Program) (*models.Program, error) {

	return &program, nil
}

// Update updates a program
func (p *Program) Update(id string, updates models.Program) (*models.Program, error) {

	return &updates, nil
}

// Delete deletes a program
func (p *Program) Delete(id string) error {

	return nil
}

// Schema returns the schema of a program
func (p *Program) Schema() (*models.Program, error) {

	return nil, nil
}

// Start starts a program
func (p *Program) Start(id string) (*models.Program, error) {

	return nil, nil
}

// Stop stops a program
func (p *Program) Stop(id string) (*models.Program, error) {

	return nil, nil
}

// Restart restarts a program
func (p *Program) Restart(id string) (*models.Program, error) {

	return nil, nil
}

// Action executes a custom action on a program
func (p *Program) Status(id string) (*models.Program, error) {

	return nil, nil
}

func (p *Program) Reload(id string) error {

	return nil
}

func (p *Program) Enable(id string) error {

	return nil
}

func (p *Program) Disable(id string) error {

	return nil
}

func (p *Program) Tail(id string) error {

	return nil
}
func (p *Program) Shell(id string) error {

	return nil
}
func (p *Program) Run(cmd string) error {

	return nil
}
