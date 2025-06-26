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

package handlers

import (
	"fmt"

	"github.com/nats-io/nats.go/micro"
)

type Service struct {
	micro.Group
}

func NewService(group micro.Group) *Service {
	return &Service{group}
}

func (e *Service) Init() (err error) {
	e.AddEndpoint("load", micro.HandlerFunc(e.Load))
	e.AddEndpoint("list", micro.HandlerFunc(e.List))
	e.AddEndpoint("list", micro.HandlerFunc(e.Get))
	e.AddEndpoint("create", micro.HandlerFunc(e.Create))
	e.AddEndpoint("update", micro.HandlerFunc(e.Update))
	e.AddEndpoint("delete", micro.HandlerFunc(e.Delete))
	e.AddEndpoint("start", micro.HandlerFunc(e.Start))
	e.AddEndpoint("stop", micro.HandlerFunc(e.Stop))
	e.AddEndpoint("restart", micro.HandlerFunc(e.Restart))
	e.AddEndpoint("status", micro.HandlerFunc(e.Status))
	e.AddEndpoint("log", micro.HandlerFunc(e.Log))
	e.AddEndpoint("shell", micro.HandlerFunc(e.Shell))

	return
}

// Runtime CRUD operations in a database

// Load a service
func (s *Service) Load(msg micro.Request) {
	fmt.Println("load:service")
}

// List all services
func (s *Service) List(msg micro.Request) {
	fmt.Println("list:services")
}

// Get a service by ID
func (s *Service) Get(msg micro.Request) {
	fmt.Println("get:service")
}

// Create a new service
func (s *Service) Create(msg micro.Request) {
	fmt.Println("create:service")
}

// Update a service
func (s *Service) Update(msg micro.Request) {
	fmt.Println("update:service")
}

// Delete a service
func (s *Service) Delete(msg micro.Request) {
	fmt.Println("delete:service")
}

// Runtime operations sending commands to a service executor plugins

// Start a service
func (s *Service) Start(msg micro.Request) {
	fmt.Println("start:service")
}

// Stop a service
func (s *Service) Stop(msg micro.Request) {
	fmt.Println("stop:service")
}

// Restart a service
func (s *Service) Restart(msg micro.Request) {
	fmt.Println("restart:service")
}

// Status of a service
func (s *Service) Status(msg micro.Request) {
	fmt.Println("status:service")
}

// Log a service
func (s *Service) Log(msg micro.Request) {
	fmt.Println("log:service")
}

func (s *Service) Shell(msg micro.Request) {
	fmt.Println("shell:service")
}
