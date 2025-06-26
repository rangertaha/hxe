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

package handlers

import (
	"fmt"

	"github.com/nats-io/nats.go/micro"
)

type Task struct {
	micro.Group
}

func NewTask(group micro.Group) *Task {
	return &Task{group}
}

func (t *Task) Init() (err error) {
	t.Group.AddEndpoint("load", micro.HandlerFunc(t.Load))
	t.Group.AddEndpoint("list", micro.HandlerFunc(t.List))
	t.Group.AddEndpoint("get", micro.HandlerFunc(t.Get))
	t.Group.AddEndpoint("create", micro.HandlerFunc(t.Create))
	t.Group.AddEndpoint("update", micro.HandlerFunc(t.Update))
	t.Group.AddEndpoint("delete", micro.HandlerFunc(t.Delete))
	t.Group.AddEndpoint("start", micro.HandlerFunc(t.Start))
	t.Group.AddEndpoint("stop", micro.HandlerFunc(t.Stop))
	t.Group.AddEndpoint("restart", micro.HandlerFunc(t.Restart))
	t.Group.AddEndpoint("status", micro.HandlerFunc(t.Status))
	t.Group.AddEndpoint("log", micro.HandlerFunc(t.Log))
	t.Group.AddEndpoint("shell", micro.HandlerFunc(t.Shell))

	return
}

// Runtime CRUD operations in a database

// Load a service
func (t *Task) Load(msg micro.Request) {
	fmt.Println("load:service")
}

// List all services
func (t *Task) List(msg micro.Request) {
	fmt.Println("list:services")
}

// Get a service by ID
func (t *Task) Get(msg micro.Request) {
	fmt.Println("get:service")
}

// Create a new service
func (t *Task) Create(msg micro.Request) {
	fmt.Println("create:service")
}

// Update a service
func (t *Task) Update(msg micro.Request) {
	fmt.Println("update:service")
}

// Delete a service
func (t *Task) Delete(msg micro.Request) {
	fmt.Println("delete:service")
}

// Runtime operations sending commands to a service executor plugins

// Start a service
func (t *Task) Start(msg micro.Request) {
	fmt.Println("start:service")
}

// Stop a service
func (t *Task) Stop(msg micro.Request) {
	fmt.Println("stop:service")
}

// Restart a service
func (t *Task) Restart(msg micro.Request) {
	fmt.Println("restart:service")
}

// Status of a service
func (t *Task) Status(msg micro.Request) {
	fmt.Println("status:service")
}

// Log a service
func (t *Task) Log(msg micro.Request) {
	fmt.Println("log:service")
}

func (t *Task) Shell(msg micro.Request) {
	fmt.Println("shell:service")
}
