package client

import (
	"fmt"

	"github.com/rangertaha/hxe/internal/models"
)

// Service operations
func (c *Client) List() ([]models.Program, error) {
	var services []models.Program
	err := c.Get("/api/services", &services)
	return services, err
}

func (c *Client) Get(id uint) (*models.Program, error) {
	var service models.Program
	err := c.Get(fmt.Sprintf("/api/services/%d", id), &service)
	return &service, err
}

func (c *Client) Create(service *models.Program) (*models.Program, error) {
	var created models.Program
	err := c.Post("/api/services", service, &created)
	return &created, err
}

func (c *Client) Update(id uint, service *models.Program) (*models.Program, error) {
	var updated models.Program
	err := c.Put(fmt.Sprintf("/api/services/%d", id), service, &updated)
	return &updated, err
}

func (c *Client) Delete(id uint) error {
	return c.Delete(fmt.Sprintf("/api/services/%d", id))
}

func (c *Client) Stop(id uint) error {
	return c.Post(fmt.Sprintf("/api/services/%d/stop", id), nil, nil)
}

func (c *Client) Start(id uint) error {
	return c.Post(fmt.Sprintf("/api/services/%d/start", id), nil, nil)
}

func (c *Client) Restart(id uint) error {
	return c.Post(fmt.Sprintf("/api/services/%d/restart", id), nil, nil)
}
