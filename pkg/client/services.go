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

// import (
// 	"fmt"

// 	"github.com/jedib0t/go-pretty/table"
// 	"github.com/rangertaha/hxe/internal/services/program/models"
// )

// // ServiceClient handles service-related API requests
// type ProgramClient struct {
// 	client *Client
// }

// // List returns all services
// func (s *ProgramClient) List() ([]*models.Program, error) {
// 	var services []*models.Program
// 	err := s.client.Get("/api/programs", &services)
// 	return services, err
// }

// // Reload reloads a service configuration
// func (s *ProgramClient) Reload(id string) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Post(fmt.Sprintf("/api/services/%s/reload", id), nil, &service)
// 	return &service, err
// }

// // Get returns a single service by ID
// func (s *ProgramClient) Get(id uint) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Get(fmt.Sprintf("/api/programs/%d", id), &service)
// 	return &service, err
// }

// // Create creates a new service
// func (s *ProgramClient) Create(service *models.Program) (*models.Program, error) {
// 	err := s.client.Post("/api/programs", service, &service)
// 	return service, err
// }

// // Update updates an existing service
// func (s *ProgramClient) Update(service *models.Program) (*models.Program, error) {
// 	err := s.client.Put(fmt.Sprintf("/api/programs/%d", service.ID), service, &service)
// 	return service, err
// }

// // Delete deletes a service
// func (s *ProgramClient) Delete(id uint) error {
// 	var response interface{}
// 	return s.client.Delete(fmt.Sprintf("/api/programs/%d", id), &response)
// }

// // Start starts a service
// func (s *ProgramClient) Start(id uint) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/start", id), nil, &service)
// 	return &service, err
// }

// // Stop stops a service
// func (s *ProgramClient) Stop(id uint) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/stop", id), nil, &service)
// 	return &service, err
// }

// // Restart restarts a service
// func (s *ProgramClient) Restart(id uint) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/restart", id), nil, &service)
// 	return &service, err
// }

// // Status gets the status of a service
// func (s *ProgramClient) Status(id uint) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/status", id), nil, &service)
// 	return &service, err
// }

// // Enable enables a service
// func (s *ProgramClient) Enable(id uint) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/enable", id), nil, &service)
// 	return &service, err
// }

// // Disable disables a service
// func (s *ProgramClient) Disable(id uint) (*models.Program, error) {
// 	var service models.Program
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/disable", id), nil, &service)
// 	return &service, err
// }

// // Shell opens a shell session for a service
// func (s *ProgramClient) Shell(id uint) (interface{}, error) {
// 	var response interface{}
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/shell", id), nil, &response)
// 	return response, err
// }

// // Log gets logs for a service
// func (s *ProgramClient) Log(id uint, logType string) (interface{}, error) {
// 	var response interface{}
// 	payload := map[string]string{"type": logType}
// 	err := s.client.Post(fmt.Sprintf("/api/programs/%d/log", id), payload, &response)
// 	return response, err
// }

// // GetByGroup returns services filtered by group ID
// func (s *ProgramClient) GetByGroup(groupID uint) ([]models.Program, error) {
// 	var services []models.Program
// 	err := s.client.Get(fmt.Sprintf("/api/programs?group_id=%d", groupID), &services)
// 	return services, err
// }

// // GetByStatus returns services filtered by status
// func (s *ProgramClient) GetByStatus(status string) ([]models.Program, error) {
// 	var services []models.Program
// 	err := s.client.Get(fmt.Sprintf("/api/programs?status=%s", status), &services)
// 	return services, err
// }

// // GetActive returns only active services
// func (s *ProgramClient) GetActive() ([]models.Program, error) {
// 	var services []models.Program
// 	err := s.client.Get("/api/programs?is_active=true", &services)
// 	return services, err
// }

// func (c *ProgramClient) PrintDetail(res *models.Program) {
// 	// Create table
// 	t := table.NewWriter()
// 	t.SetOutputMirror(nil)

// 	t.AppendHeader(table.Row{"Field", "Value"})

// 	// Add rows
// 	t.AppendRow(table.Row{"ID", res.ID})
// 	t.AppendRow(table.Row{"Name", res.Name})
// 	t.AppendRow(table.Row{"Description", res.Desc})

// 	// Print the table
// 	fmt.Println(t.Render())
// }

// func (c *ProgramClient) PrintList(services []*models.Program) {
// 	// Create table
// 	t := table.NewWriter()
// 	t.SetOutputMirror(nil)

// 	t.AppendHeader(table.Row{"ID", "Name", "Description"})

// 	// Add rows
// 	for _, service := range services {
// 		t.AppendRow(table.Row{
// 			service.ID,
// 			service.Name,
// 			service.Desc,
// 		})
// 	}
// 	// Print the table
// 	fmt.Println(t.Render())
// }

// func (c *ProgramClient) Print(services ...*models.Program) {
// 	if len(services) > 1 {
// 		c.PrintList(services)
// 	} else if len(services) == 1 {
// 		c.PrintDetail(services[0])
// 	}
// }
