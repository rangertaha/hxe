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
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/client"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rangertaha/hxe/internal/modules/services"
	"github.com/rangertaha/hxe/internal/rdb"
	"github.com/rs/zerolog"
)

var (
	ErrNotInitialized = Error{Code: http.StatusInternalServerError, Message: "Client not initialized"}
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Service struct {
	Client *client.Client
	log    zerolog.Logger
}

func NewService(c *client.Client) *Service {
	return &Service{
		Client: c,
		log:    log.With().Logger(),
	}
}

func ServiceRoutes(e *echo.Group, c *client.Client) {
	// Services
	service := NewService(c)
	svc := e.Group("/service")
	svc.GET("", service.List)
	svc.GET("/:id", service.Get)
	svc.POST("", service.Create)
	svc.PUT("/:id", service.Update)
	svc.DELETE("/:id", service.Delete)
	// svc.OPTIONS("", service.Schema)

	// Runtime handlers
	svc.POST("/:id/start", service.Start)
	svc.POST("/:id/stop", service.Stop)
	svc.POST("/:id/restart", service.Restart)
	svc.POST("/:id/status", service.Status)
	svc.POST("/:id/reload", service.Reload)
	svc.POST("/:id/enable", service.Enable)
	svc.POST("/:id/disable", service.Disable)

	// Stream handlers
	svc.POST("/:id/shell", service.Shell)
	svc.POST("/:id/log", service.Log)
}

// CRUD HANDLERS
func (s *Service) List(c echo.Context) error {

	req := &services.Request{
		Service: &rdb.Service{},
	}

	data, _ := json.Marshal(req)

	bytes, err := s.Client.Services.Send("service.list", data)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, bytes)
}

func (s *Service) Get(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	record, err := s.Client.Services.Get(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, record)
}

func (s *Service) Create(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var prog rdb.Service
	if err := c.Bind(&prog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	record, err := s.Client.Services.Create(&prog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, record)
}

func (s *Service) Update(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var svc rdb.Service
	if err := c.Bind(&svc); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	service, err := s.Client.Services.Update(&svc)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (s *Service) Delete(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

// RUNTIME HANDLERS
func (s *Service) Start(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Start(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Stop(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Stop(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Restart(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Restart(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Status(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Status(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Reload(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Reload(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Enable(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Enable(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Disable(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Disable(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Shell(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Shell(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}

func (s *Service) Log(c echo.Context) error {
	if s.Client == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := s.Client.Services.Log(uint(id), "stdout")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, svc)
}
