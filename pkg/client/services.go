package client

import (
	"fmt"

	"github.com/rangertaha/hxe/internal/db"
)

// Service operations
func (c *Client) ListServices() ([]db.Service, error) {
	var services []db.Service
	err := c.Get("/api/services", &services)
	return services, err
}

func (c *Client) GetService(id uint) (*db.Service, error) {
	var service db.Service
	err := c.Get(fmt.Sprintf("/api/services/%d", id), &service)
	return &service, err
}

func (c *Client) CreateService(service *db.Service) (*db.Service, error) {
	var created db.Service
	err := c.Post("/api/services", service, &created)
	return &created, err
}

func (c *Client) UpdateService(id uint, service *db.Service) (*db.Service, error) {
	var updated db.Service
	err := c.Put(fmt.Sprintf("/api/services/%d", id), service, &updated)
	return &updated, err
}

func (c *Client) DeleteService(id uint) error {
	return c.Delete(fmt.Sprintf("/api/services/%d", id))
}

func (c *Client) StopService(id uint) error {
	return c.Post(fmt.Sprintf("/api/services/%d/stop", id), nil, nil)
}

func (c *Client) StartService(id uint) error {
	return c.Post(fmt.Sprintf("/api/services/%d/start", id), nil, nil)
}

func (c *Client) RestartService(id uint) error {
	return c.Post(fmt.Sprintf("/api/services/%d/restart", id), nil, nil)
}

