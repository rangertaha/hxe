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
	"github.com/rangertaha/hxe/internal/modules/services/models"
)

type Runner struct {
	services map[string]*models.Service
}

func NewRunner() *Runner {
	return &Runner{
		services: make(map[string]*models.Service),
	}
}

func (m *Runner) Init() (err error) {

	return
}

// Runtime CRUD operations in a database

// // Load a service
// func (m *Runner) Load(svc *models.Service) {
// 	svc.Status = models.ServiceStatus_LOADING
// 	m.services[svc.Name] = svc
// }

// // List all services
// func (m *Runner) List(svc *models.Service) {

// }

// // Start a service
// func (m *Runner) Start(svc *models.Service) {
// 	svc.Status = models.ServiceStatus_STARTING
// 	m.services[svc.Name] = svc
// }

// // Stop a service
// func (m *Runner) Stop(svc *models.Service) {
// 	svc.Status = models.ServiceStatus_STOPPING
// 	m.services[svc.Name] = svc
// }

// // Restart a service
// func (m *Runner) Restart(svc *models.Service) {
// 	svc.Status = models.ServiceStatus_RESTARTING
// 	m.services[svc.Name] = svc
// }

// // Status of a service
// func (m *Runner) Status(svc *models.Service) {
// 	svc.Status = models.ServiceStatus_RUNNING
// 	m.services[svc.Name] = svc
// }

// // Log a service
// func (m *Runner) Log(svc *models.Service) {
// 	fmt.Println("log:service")

// }

// func (m *Runner) Shell(svc *models.Service) {
// 	fmt.Println("shell:service")
// }
