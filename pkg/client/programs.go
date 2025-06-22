package client

import (
	"fmt"

	"github.com/rangertaha/hxe/internal/models"
)

type ProgramClient struct {
	GenericClient
}

// Program operations
func (c *Client) ListPrograms() ([]models.Program, error) {
	var programs []models.Program
	err := c.Get("/api/program", &programs)
	return programs, err
}

func (c *Client) GetProgram(id uint) (*models.Program, error) {
	var program models.Program
	err := c.Get(fmt.Sprintf("/api/program/%d", id), &program)
	return &program, err
}

func (c *Client) CreateProgram(program *models.Program) (*models.Program, error) {
	var created models.Program
	err := c.Post("/api/program", program, &created)
	return &created, err
}

func (c *Client) UpdateProgram(id uint, program *models.Program) (*models.Program, error) {
	var updated models.Program
	err := c.Put(fmt.Sprintf("/api/program/%d", id), program, &updated)
	return &updated, err
}

func (c *Client) DeleteProgram(id uint) error {
	return c.Delete(fmt.Sprintf("/api/program/%d", id))
}

func (c *Client) StopProgram(id uint) error {
	return c.Post(fmt.Sprintf("/api/program/%d/stop", id), nil, nil)
}

func (c *Client) StartProgram(id uint) error {
	return c.Post(fmt.Sprintf("/api/program/%d/start", id), nil, nil)
}

func (c *Client) RestartProgram(id uint) error {
	return c.Post(fmt.Sprintf("/api/program/%d/restart", id), nil, nil)
}
