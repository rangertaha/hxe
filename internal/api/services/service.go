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

package services

import (
	"fmt"
	"strconv"

	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/models/db"
	pb "github.com/rangertaha/hxe/internal/models/pb"
	"gorm.io/gorm"
)

type Service struct {
	db     *gorm.DB
	broker internal.Broker
}

func NewService(b internal.Broker) *Service {
	return &Service{
		db:     db.DB,
		broker: b,
	}
}

// List returns all services
func (p *Service) List() ([]pb.Service, error) {
	var svcs []db.Service
	if err := p.db.Find(&svcs).Error; err != nil {
		return nil, fmt.Errorf("failed to list services: %w", err)
	}
	return models.ToPb(svcs...), nil
}

// Get returns a service by ID
func (p *Service) Get(id string) (*pb.Service, error) {
	var service pb.Service
	serviceID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("invalid service ID: %w", err)
	}

	if err := p.pb.First(&service, serviceID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("service not found")
		}
		return nil, fmt.Errorf("failed to get service: %w", err)
	}

	return &service, nil
}

// Create creates a new service
func (p *Service) Create(service *pb.Service) (*pb.Service, error) {
	if err := p.db.Create(&service).Error; err != nil {
		return nil, fmt.Errorf("failed to create service: %w", err)
	}
	return &service, nil
}

// Update updates a service
func (p *Service) Update(id string, updates *pb.Service) (*pb.Service, error) {
	if err := p.db.Save(&updates).Error; err != nil {
		return nil, fmt.Errorf("failed to update service: %w", err)
	}
	return &updates, nil
}

// Delete deletes a service
func (p *Service) Delete(id string) (*pb.Service, error) {
	deleted, err := p.Stop(id)
	if err != nil {
		return nil, err
	}

	if err := p.db.Delete(&deleted, id).Error; err != nil {
		return nil, fmt.Errorf("failed to delete service: %w", err)
	}

	return deleted, nil
}

// Schema returns the schema of a service
func (p *Service) Schema() (*pb.Schema, error) {
	return pb.ServiceSchema(), nil
}

// Start starts a service
func (p *Service) Start(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// Update status to started
	service.Status = pb.ServiceStatus_STARTED

	if err := p.db.Save(service).Error; err != nil {
		return nil, fmt.Errorf("failed to start service: %w", err)
	}

	return service, nil
}

// Stop stops a service
func (p *Service) Stop(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// Update status to stopped
	service.Status = pb.ServiceStatus_STOPPED

	if err := p.db.Save(service).Error; err != nil {
		return nil, fmt.Errorf("failed to stop service: %w", err)
	}

	return service, nil
}

// Restart restarts a service
func (p *Service) Restart(id string) (*pb.Service, error) {
	// Stop the service first
	if _, err := p.Stop(id); err != nil {
		return nil, err
	}

	// Start the service
	return p.Start(id)
}

// Status returns the status of a service
func (p *Service) Status(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	return service, nil
}

// Reload reloads the configuration for a service
func (p *Service) Reload(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	service.Status = pb.ServiceStatus_RELOADED

	if err := p.pb.Save(service).Error; err != nil {
		return nil, fmt.Errorf("failed to reload service: %w", err)
	}
	return service, nil
}

// Enable enables a service to start automatically
func (p *Service) Enable(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	service.Enabled = true
	if err := p.pb.Save(service).Error; err != nil {
		return nil, fmt.Errorf("failed to enable service: %w", err)
	}

	return service, nil
}

// Disable disables a service from starting automatically
func (p *Service) Disable(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	service.Enabled = false
	service.Status = pb.ServiceStop
	if err := p.pb.Save(service).Error; err != nil {
		return nil, fmt.Errorf("failed to disable service: %w", err)
	}

	return service, nil
}

// Tail follows the logs of a service
func (p *Service) Log(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// In a real implementation, you would start tailing the service's log file
	// For now, we just return success
	return service, nil
}

// Shell opens an interactive shell for a service
func (p *Service) Shell(id string) (*pb.Service, error) {
	service, err := p.Get(id)
	if err != nil {
		return nil, err
	}

	// In a real implementation, you would open an interactive shell
	// For now, we just return success
	return service, nil
}
