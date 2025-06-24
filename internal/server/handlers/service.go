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

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/server/services"
	"github.com/rangertaha/hxe/internal/models"
)

type Service struct {
	Svc *services.Service
}

func NewService(b internal.Broker) *Service {
	return &Service{
		Svc: services.NewService(b),
	}
}

// CRUD HANDLERS
func (h *Service) List(c echo.Context) error {
	services, err := h.Svc.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, services)
}

func (h *Service) Get(c echo.Context) error {
	service, err := h.Svc.Get(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Create(c echo.Context) error {
	var prog models.Service
	if err := c.Bind(&prog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	service, err := h.Svc.Create(prog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Update(c echo.Context) error {
	var prog models.Service
	if err := c.Bind(&prog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	service, err := h.Svc.Update(c.Param("id"), prog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Delete(c echo.Context) error {
	service, err := h.Svc.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Schema(c echo.Context) error {
	schema, err := h.Svc.Schema()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schema)
}

// RUNTIME HANDLERS
func (h *Service) Start(c echo.Context) error {
	service, err := h.Svc.Start(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Stop(c echo.Context) error {
	service, err := h.Svc.Stop(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Restart(c echo.Context) error {
	service, err := h.Svc.Restart(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Status(c echo.Context) error {
	service, err := h.Svc.Status(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Reload(c echo.Context) error {
	service, err := h.Svc.Reload(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Enable(c echo.Context) error {
	service, err := h.Svc.Enable(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Disable(c echo.Context) error {
	service, err := h.Svc.Disable(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

// STREAM HANDLERS
func (h *Service) Shell(c echo.Context) error {
	service, err := h.Svc.Shell(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}

func (h *Service) Log(c echo.Context) error {
	service, err := h.Svc.Log(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, service)
}
