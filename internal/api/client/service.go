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

package client

import (
	"github.com/nats-io/nats.go"
	models "github.com/rangertaha/hxe/internal/api/models"
)

type Service struct {
	nc *nats.Conn
}

func NewService(nc *nats.Conn) (e *Service) {
	return &Service{
		nc: nc,
	}
}

// Load a service
func (e *Service) Load(id uint) (err error) {
	return
}

// List all services
func (e *Service) List() (records []*models.Service, err error) {
	return
}

// Get a service by ID
func (e *Service) Get(id uint) (svc *models.Service, err error) {
	return
}

// Create a new service
func (e *Service) Create(svc *models.Service) (s *models.Service, err error) {
	return
}

// Update a service
func (e *Service) Update(svc *models.Service) (s *models.Service, err error) {
	return
}

// Delete a service
func (e *Service) Delete(id uint) (s *models.Service, err error) {
	return
}

// Start a service
func (e *Service) Start(id uint) (s *models.Service, err error) {
	return
}

// Stop a service
func (e *Service) Stop(id uint) (s *models.Service, err error) {
	return
}

// Restart a service
func (e *Service) Restart(id uint) (s *models.Service, err error) {
	return
}

// Status returns service status
func (e *Service) Status(id uint) (s *models.Service, err error) {
	return
}

// Reload reloads a service
func (e *Service) Reload(id uint) (s *models.Service, err error) {
	return
}

// Enable enables a service
func (e *Service) Enable(id uint) (s *models.Service, err error) {
	return
}

// Disable disables a service
func (e *Service) Disable(id uint) (s *models.Service, err error) {
	return
}

// Shell opens a shell for a service
func (e *Service) Shell(id uint) (s *models.Service, err error) {
	return
}

// Log returns service logs
func (e *Service) Log(id uint) (s *models.Service, err error) {
	return
}
