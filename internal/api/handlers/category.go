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
	"github.com/rangertaha/hxe/internal/client"
	"github.com/rangertaha/hxe/internal/modules/services/models"
)

type Category struct {
	Client *client.Client
}

func NewCategory(c *client.Client) *Category {
	return &Category{
		Client: c,
	}
}

func CategoryRoutes(e *echo.Group, c *client.Client) {
	category := NewCategory(c)
	cat := e.Group("/category")
	cat.GET("", category.List)
	cat.GET("/:id", category.Get)
	cat.POST("", category.Create)
	cat.PUT("/:id", category.Update)
	cat.DELETE("/:id", category.Delete)
}

// CRUD HANDLERS
func (cat *Category) List(ctx echo.Context) error {
	if cat.Client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	records, err := cat.Client.Services.List()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, records)
}

func (cat *Category) Get(ctx echo.Context) error {
	if cat.Client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	record, err := cat.Client.Services.Get(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, record)
}

func (cat *Category) Create(ctx echo.Context) error {
	if cat.Client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var category models.Service
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	record, err := cat.Client.Services.Create(&category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, record)
}

func (cat *Category) Update(ctx echo.Context) error {
	if cat.Client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	var category models.Service
	if err := ctx.Bind(&category); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	service, err := cat.Client.Services.Update(&category)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, service)
}

// RUNTIME HANDLERS
func (cat *Category) Delete(ctx echo.Context) error {
	if cat.Client == nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "Client not initialized"})
	}
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID"})
	}
	svc, err := cat.Client.Services.Delete(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return ctx.JSON(http.StatusOK, svc)
}
