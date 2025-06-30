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
)


type Runner struct {
}

func NewRunner() *Runner {
	return &Runner{}
}

func (m *Runner) Init() (err error) {

	return
}

// Runtime CRUD operations in a database

// Load a service
func (m *Runner) Load() {
	fmt.Println("load:service")
}

// List all services
func (m *Runner) List() {

}

// Start a service
func (m *Runner) Start() {
	fmt.Println("start:service")

}

// Stop a service
func (m *Runner) Stop() {
	fmt.Println("stop:service")

}

// Restart a service
func (m *Runner) Restart() {
	fmt.Println("restart:service")

}

// Status of a service
func (m *Runner) Status() {
	fmt.Println("status:service")

}

// Log a service
func (m *Runner) Log() {
	fmt.Println("log:service")

}

func (m *Runner) Shell() {
	fmt.Println("shell:service")
}
