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

// import (
// 	models "github.com/rangertaha/hxe/internal/api/models"
// )

// type Service struct {
// 	bc *BaseClient
// }

// func NewService(bc *BaseClient) *Service {
// 	return &Service{bc: bc}
// }

// // Load a service
// func (s *Service) Load(id uint) (resp *models.Service, err error){
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.get", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // List all services
// func (s *Service) List() (resp *models.Services, err error) {
// 	req := &models.Service{}
// 	err = s.bc.Send("service.list", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// func (s *Service) Get(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.get", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp, nil
// }

// // Create a new service
// func (s *Service) Create(req *models.Service) (resp *models.Service, err error) {
// 	err = s.bc.Send("service.create", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Update a service
// func (s *Service) Update(req *models.Service) (resp *models.Service, err error) {
// 	err = s.bc.Send("service.update", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Delete a service
// func (s *Service) Delete(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.delete", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Start a service
// func (s *Service) Start(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.start", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Stop a service
// func (s *Service) Stop(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.stop", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Restart a service
// func (s *Service) Restart(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.restart", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Status returns service status
// func (s *Service) Status(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.status", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Reload reloads a service
// func (s *Service) Reload(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.reload", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Enable enables a service
// func (s *Service) Enable(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.enable", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Disable disables a service
// func (s *Service) Disable(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.disable", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Shell opens a shell for a service
// func (s *Service) Shell(id uint) (resp *models.Service, err error) {
// 	req := &models.Service{Id: uint32(id)}
// 	err = s.bc.Send("service.shell", req, resp)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }

// // Log returns service logs
// func (s *Service) Log(id uint, stream string) (resp chan []byte, err error) {
// 	resp, err = s.bc.Stream("service.log", id, stream)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return
// }
