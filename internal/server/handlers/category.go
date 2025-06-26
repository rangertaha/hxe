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

type Category struct {
	client *client.Client
}

func NewCategory(b internal.Broker) *Category {
	// For now, return a service with nil client since we can't directly access NATS connection
	// In a real implementation, you'd need to modify the Broker interface or create an adapter
	return &Category{
		client: nil,
	}
}

// CRUD HANDLERS
func (cat *Category) List(ctx echo.Context) error {
	if cat.client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	records, err := cat.client.Category.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, records)
}

func (cat *Category) Get(ctx echo.Context) error {
	if cat.client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	record, err := cat.client.Category.Get(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, record)
}

func (cat *Category) Create(ctx echo.Context) error {
	if cat.client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var category models.Category
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	record, err := cat.client.Category.Create(&category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, record)
}

func (cat *Category) Update(ctx echo.Context) error {
	if cat.client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var category models.Category
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	service, err := cat.client.Category.Update(&category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, service)
}

// RUNTIME HANDLERS
func (cat *Category) Delete(ctx echo.Context) error {
	if cat.client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := cat.client.Category.Restart(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, svc)
}
