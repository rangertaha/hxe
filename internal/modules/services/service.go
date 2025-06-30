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

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/modules/services/models"
	"github.com/rangertaha/hxe/internal/rdb"
	"google.golang.org/protobuf/proto"
)

type Service struct {
	Runner  *Runner
	Service micro.Service
	Client  *Client
}

func NewService(nc *nats.Conn) (s *Service, err error) {
	s = &Service{
		Runner:  NewRunner(),
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
	svc.AddEndpoint("load", micro.HandlerFunc(e.Load))
	svc.AddEndpoint("list", micro.HandlerFunc(e.List))
	svc.AddEndpoint("get", micro.HandlerFunc(e.Get))
	svc.AddEndpoint("create", micro.HandlerFunc(e.Create))
	svc.AddEndpoint("update", micro.HandlerFunc(e.Update))
	svc.AddEndpoint("delete", micro.HandlerFunc(e.Delete))
	svc.AddEndpoint("start", micro.HandlerFunc(e.Start))
	svc.AddEndpoint("stop", micro.HandlerFunc(e.Stop))
	svc.AddEndpoint("restart", micro.HandlerFunc(e.Restart))
	svc.AddEndpoint("status", micro.HandlerFunc(e.Status))
	svc.AddEndpoint("log", micro.HandlerFunc(e.Log))
	svc.AddEndpoint("shell", micro.HandlerFunc(e.Shell))

	return
}

// Runtime CRUD operations in a database

// Load a service
func (s *Service) Load(msg micro.Request) {
	fmt.Println("load:service")
}

// List all services
func (s *Service) List(msg micro.Request) {
	var req models.Service
	if err := proto.Unmarshal(msg.Data(), &req); err != nil {
		msg.Error("500", "Failed to unmarshal request", nil)
		return
	}

	// Get services from database
	services := []rdb.Service{}
	if err := rdb.DB.Find(&services).Error; err != nil {
		msg.Error("500", "Database error", nil)
		return
	}
	for _, service := range services {
		log.Println(service)
	}

	protoServices := models.ToProtoServices(services)

	// Marshal response
	data, err := proto.Marshal(protoServices)
	if err != nil {
		msg.Error("500", "Failed to marshal response", nil)
		return
	}
	msg.Respond(data)
}

// Get a service by ID
func (s *Service) Get(msg micro.Request) {
	// var req models.Services
	// msg.Unmarshal(req)

	fmt.Println("get:service")
	msg.Respond([]byte("get:service"))
}

// Create a new service
func (s *Service) Create(msg micro.Request) {
	fmt.Println("create:service")
	msg.Respond([]byte("create:service"))
}

// Update a service
func (s *Service) Update(msg micro.Request) {
	fmt.Println("update:service")
	msg.Respond([]byte("update:service"))
}

// Delete a service
func (s *Service) Delete(msg micro.Request) {
	fmt.Println("delete:service")
	msg.Respond([]byte("delete:service"))
}

// Runtime operations sending commands to a service executor plugins

// Start a service
func (s *Service) Start(msg micro.Request) {
	s.Runner.Start()

	fmt.Println("start:service")
	msg.Respond([]byte("start:service"))
}

// Stop a service
func (s *Service) Stop(msg micro.Request) {
	s.Runner.Stop()
	fmt.Println("stop:service")
	msg.Respond([]byte("stop:service"))
}

// Restart a service
func (s *Service) Restart(msg micro.Request) {
	s.Runner.Restart()

	fmt.Println("restart:service")
	msg.Respond([]byte("restart:service"))
}

// Status of a service
func (s *Service) Status(msg micro.Request) {
	s.Runner.Status()

	fmt.Println("status:service")
	msg.Respond([]byte("status:service"))
}

// Log a service
func (s *Service) Log(msg micro.Request) {
	fmt.Println("log:service")
	msg.Respond([]byte("log:service"))
}

func (s *Service) Shell(msg micro.Request) {
	fmt.Println("shell:service")
	msg.Respond([]byte("shell:service"))
}
