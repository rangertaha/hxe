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
	"log"

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
	Storage *models.Storage
}

func NewService(nc *nats.Conn) (s *Service, err error) {
	s = &Service{
		Storage: models.NewStorage(rdb.DB),
		Runner:  NewRunner(),
	}
	config := micro.Config{
		Name:        "Services",
		Version:     internal.VERSION,
		Description: "Process manager",
	}

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
	svc.AddEndpoint("load", handlers.Proto(e.Load))
	svc.AddEndpoint("list", handlers.Proto(e.List))
	svc.AddEndpoint("get", handlers.Proto(e.Get))
	svc.AddEndpoint("create", handlers.Proto(e.Create))
	svc.AddEndpoint("update", handlers.Proto(e.Update))
	svc.AddEndpoint("delete", handlers.Proto(e.Delete))
	svc.AddEndpoint("start", handlers.Proto(e.Start))
	svc.AddEndpoint("stop", handlers.Proto(e.Stop))
	svc.AddEndpoint("restart", handlers.Proto(e.Restart))
	svc.AddEndpoint("status", handlers.Proto(e.Status))
	svc.AddEndpoint("log", handlers.Proto(e.Log))
	svc.AddEndpoint("shell", handlers.Proto(e.Shell))

	return
}

// Load a service (no request needed)
func (s *Service) Load(req *models.Request) (res *models.Response) {
	s.Runner.Load(req.Service)
	return
}

// List all services
func (s *Service) List(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Services, res.Status = s.Storage.List(req.Service)
	return
}

// Get a service by ID
func (s *Service) Get(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Service, res.Status = s.Storage.Get(req.Service)
	return
}

// Create a new service
func (s *Service) Create(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Service, res.Status = s.Storage.Create(req.Service)
	return
}

// Update a service
func (s *Service) Update(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Service, res.Status = s.Storage.Update(req.Service)
	return
}

// Delete a service
func (s *Service) Delete(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Service, res.Status = s.Storage.Delete(req.Service)
	return
}

// Start a service
func (s *Service) Start(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Service, res.Status = s.Storage.Get(req.Service)
	s.Runner.Start(res.Service)
	res.Service.Save()
	return res
}

// Stop a service
func (s *Service) Stop(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Service, res.Status = s.Storage.Get(req.Service)
	s.Runner.Stop(res.Service)
	return
}

// Restart a service
func (s *Service) Restart(req *models.Request) (res *models.Response) {
	res = &models.Response{}
	res.Service, res.Status = s.Storage.Get(req.Service)
	s.Runner.Restart(res.Service)
	res.Service.Save()
	return
}

// Status of a service
func (s *Service) Status(req *models.Request) (res *models.Response) {
	return s.Get(req)
}

// Log a service
func (s *Service) Log(req *models.Request) (res *models.Response) {

	return
}

// Shell a service
func (s *Service) Shell(req *models.Request) (res *models.Response) {

	return
}
