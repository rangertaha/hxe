/*
 * HXE - Host-based Process Execution Engine
 * Copyright (C) 2025 Rangertaha <rangertaha@gmail.com>
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

package client

import (
	"fmt"
	"os"
	"os/user"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/rangertaha/hxe/internal/models"
)

type ServiceClient struct {
	client *Client
	schema *models.Schema
}

func NewServiceClient(client *Client) *ServiceClient {
	return &ServiceClient{client: client, schema: models.ServiceSchema()}
}

// Service operations
func (c *ServiceClient) List() ([]*models.Service, error) {
	var services []*models.Service
	err := c.client.Get("/api/service", &services)
	return services, err
}

func (c *ServiceClient) Get(id string) (*models.Service, error) {
	var service *models.Service
	err := c.client.Get(fmt.Sprintf("/api/service/%s", id), &service)
	return service, err
}

func (c *ServiceClient) Status(id string) (*models.Service, error) {
	return c.Get(id)
}

func (c *ServiceClient) MultiStatus(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Status(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Create(service *models.Service) (*models.Service, error) {
	var created models.Service
	err := c.client.Post("/api/service", service, &created)
	return &created, err
}

func (c *ServiceClient) Update(id string, service *models.Service) (*models.Service, error) {
	var updated models.Service
	err := c.client.Put(fmt.Sprintf("/api/service/%s", id), service, &updated)
	return &updated, err
}

func (c *ServiceClient) Delete(id string) (*models.Service, error) {
	var deleted models.Service
	err := c.client.Delete(fmt.Sprintf("/api/service/%s", id), &deleted)
	return &deleted, err
}

func (c *ServiceClient) MultiDelete(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Delete(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Start(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/start", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) MultiStart(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Start(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Stop(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/stop", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) MultiStop(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Stop(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Restart(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/restart", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) MultiRestart(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Restart(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Enable(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/enable", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) MultiEnable(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Enable(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Disable(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/disable", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) MultiDisable(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Disable(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Reload(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/reload", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) MultiReload(ids ...string) (services []*models.Service, err error) {
	for _, id := range ids {
		service, err := c.Reload(id)
		if err != nil {
			return nil, err
		}
		services = append(services, service)
	}
	return services, nil
}

func (c *ServiceClient) Shell(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/shell", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) Tail(id string) (*models.Service, error) {
	var service models.Service
	err := c.client.Post(fmt.Sprintf("/api/service/%s/tail", id), nil, &service)
	return &service, err
}

func (c *ServiceClient) Run(command string) (*models.Service, error) {
	dir, err := os.Getwd()
	if err != nil {
		dir = "/tmp"
	}

	currentUser, err := user.Current()
	if err != nil {
		return nil, fmt.Errorf("failed to get current user: %w", err)
	}

	service := models.Service{
		Command:   command,
		Directory: dir,
		User:      currentUser.Username,
		Group:     currentUser.Gid,
		Status:    models.ServiceStart,
		Enabled:   true,
	}

	return c.Create(&service)
}

func (c *ServiceClient) PrintDetail(service *models.Service) {
	// Create table
	t := table.NewWriter()
	t.SetOutputMirror(nil)

	t.AppendHeader(table.Row{"Field", "Value"})

	// Add rows
	t.AppendRow(table.Row{"ID", service.ID})
	t.AppendRow(table.Row{"Name", service.Name})
	t.AppendRow(table.Row{"Command", service.Command})
	t.AppendRow(table.Row{"Status", service.Status})
	t.AppendRow(table.Row{"Enabled", service.Enabled})

	// Print the table
	fmt.Println(t.Render())
}

func (c *ServiceClient) PrintList(services []*models.Service) {
	// Create table
	t := table.NewWriter()
	t.SetOutputMirror(nil)

	t.AppendHeader(table.Row{"ID", "Name", "Command", "Status", "Enabled"})

	// Add rows
	for _, service := range services {
		t.AppendRow(table.Row{
			service.ID,
			service.Name,
			service.Command,
			service.Status,
			service.Enabled,
		})
	}
	// Print the table
	fmt.Println(t.Render())
}

func (c *ServiceClient) Print(services []*models.Service) {
	if len(services) == 10 {
		fmt.Println("No services found")
		return
	}
	if len(services) == 1 {
		c.PrintDetail(services[0])
		return
	}
	if len(services) > 1 {
		c.PrintList(services)
		return
	}
}
