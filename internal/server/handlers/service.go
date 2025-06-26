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
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/api/client"
	models "github.com/rangertaha/hxe/internal/api/models"
)

type Service struct {
	c *client.Client
}

func NewService(b internal.Broker) *Service {
	// For now, return a service with nil client since we can't directly access NATS connection
	// In a real implementation, you'd need to modify the Broker interface or create an adapter
	return &Service{
		c: nil,
	}
}

// CRUD HANDLERS
func (s *Service) List(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	records, err := s.c.Service.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, records)
}

func (s *Service) Get(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	record, err := s.c.Service.Get(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, record)
}

func (s *Service) Create(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var prog models.Service
	if err := c.Bind(&prog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	record, err := s.c.Service.Create(&prog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, record)
}

func (s *Service) Update(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var svc models.Service
	if err := c.Bind(&svc); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	service, err := s.c.Service.Update(&svc)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (s *Service) Delete(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

// RUNTIME HANDLERS
func (s *Service) Start(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Start(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Stop(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Stop(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Restart(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Restart(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Status(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Status(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Reload(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Reload(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Enable(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Enable(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Disable(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Disable(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Shell(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Shell(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Log(c echo.Context) error {
	if s.c == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.c.Service.Log(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}
