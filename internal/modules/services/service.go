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
	"log"
	"strings"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/modules/services/handlers"
	"github.com/rangertaha/hxe/internal/modules/services/models"
	"github.com/rangertaha/hxe/internal/rdb"
)

type Service struct {
	Runner  *Runner
	Service micro.Service
	Client  *Client
}

func NewService(nc *nats.Conn) (s *Service, err error) {
	s = &Service{
		Runner: NewRunner(),
	}
	config := micro.Config{
		Name:        "Services",
		Version:     internal.VERSION,
		Description: "Process manager",
	}

	s.Client = NewClient(nc)

	s.Service, err = micro.AddService(nc, config)
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Init(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return
}

func (e *Service) Init() (err error) {
	svc := e.Service.AddGroup("service")

	// Custom error handler for better error responses
	errorHandler := func(err error) (string, string, map[string]string) {
		// You can implement custom error mapping here
		if strings.Contains(err.Error(), "database") {
			return "503", "Service temporarily unavailable", map[string]string{"retry_after": "30"}
		}
		if strings.Contains(err.Error(), "not found") {
			return "404", "Resource not found", nil
		}
		return "500", err.Error(), nil
	}

	// Use the wrapper functions for cleaner handler registration
	svc.AddEndpoint("load", handlers.WrapSimple(e.Load))
	svc.AddEndpoint("list", handlers.WrapWithErrorHandler(e.List, errorHandler))
	svc.AddEndpoint("get", handlers.WrapWithErrorHandler(e.Get, errorHandler))
	svc.AddEndpoint("create", handlers.Wrap(e.Create))
	svc.AddEndpoint("update", handlers.Wrap(e.Update))
	svc.AddEndpoint("delete", handlers.Wrap(e.Delete))
	svc.AddEndpoint("start", handlers.WrapNoResponse(e.Start))
	svc.AddEndpoint("stop", handlers.WrapNoResponse(e.Stop))
	svc.AddEndpoint("restart", handlers.WrapNoResponse(e.Restart))
	svc.AddEndpoint("status", handlers.Wrap(e.Status))
	svc.AddEndpoint("log", handlers.Wrap(e.Log))
	svc.AddEndpoint("shell", handlers.Wrap(e.Shell))

	return
}

// Now your handlers become much simpler - they just focus on business logic

// Load a service (no request needed)
func (s *Service) Load() (*models.Services, error) {
	// Your load logic here
	return &models.Services{}, nil
}

// List all services
func (s *Service) List(req *models.Service) (*models.Services, error) {
	// Get services from database
	services := []rdb.Service{}
	if err := rdb.DB.Find(&services).Error; err != nil {
		return nil, fmt.Errorf("database error: %w", err)
	}

	for _, service := range services {
		log.Println(service)
	}

	return models.ToProtoServices(services), nil
}

// Get a service by ID
func (s *Service) Get(req *models.Service) (*models.Service, error) {
	// Your get logic here
	return req, nil
}

// Create a new service
func (s *Service) Create(req *models.Service) (*models.Service, error) {
	// Your create logic here
	return req, nil
}

// Update a service
func (s *Service) Update(req *models.Service) (*models.Service, error) {
	// Your update logic here
	return req, nil
}

// Delete a service
func (s *Service) Delete(req *models.Service) (*models.Service, error) {
	// Your delete logic here
	return req, nil
}

// Start a service (no response needed)
func (s *Service) Start(req *models.Service) error {
	s.Runner.Start()
	return nil
}

// Stop a service (no response needed)
func (s *Service) Stop(req *models.Service) error {
	s.Runner.Stop()
	return nil
}

// Restart a service (no response needed)
func (s *Service) Restart(req *models.Service) error {
	s.Runner.Restart()
	return nil
}

// Status of a service
func (s *Service) Status(req *models.Service) (*models.Service, error) {
	s.Runner.Status()
	// Return status response
	return req, nil
}

// Log a service
func (s *Service) Log(req *models.Service) (*models.Service, error) {
	// Your log logic here
	return req, nil
}

// Shell a service
func (s *Service) Shell(req *models.Service) (*models.Service, error) {
	// Your shell logic here
	return req, nil
}
