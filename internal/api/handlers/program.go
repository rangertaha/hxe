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
	"github.com/rangertaha/hxe/internal/api/services"
	"github.com/rangertaha/hxe/internal/models"
)

type Program struct {
	Prog *services.Program
}

func NewProgram(b internal.Broker) *Program {
	return &Program{
		Prog: services.NewProgram(b),
	}
}

// CRUD HANDLERS
func (h *Program) List(c echo.Context) error {
	programs, err := h.Prog.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, programs)
}

func (h *Program) Get(c echo.Context) error {
	program, err := h.Prog.Get(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Create(c echo.Context) error {
	var prog models.Program
	if err := c.Bind(&prog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	program, err := h.Prog.Create(prog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Update(c echo.Context) error {
	var prog models.Program
	if err := c.Bind(&prog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	program, err := h.Prog.Update(c.Param("id"), prog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Delete(c echo.Context) error {
	program, err := h.Prog.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Schema(c echo.Context) error {
	schema, err := h.Prog.Schema()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schema)
}

// RUNTIME HANDLERS
func (h *Program) Start(c echo.Context) error {
	program, err := h.Prog.Start(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Stop(c echo.Context) error {
	program, err := h.Prog.Stop(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Restart(c echo.Context) error {
	program, err := h.Prog.Restart(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Status(c echo.Context) error {
	program, err := h.Prog.Status(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Reload(c echo.Context) error {
	program, err := h.Prog.Reload(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Enable(c echo.Context) error {
	program, err := h.Prog.Enable(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Disable(c echo.Context) error {
	program, err := h.Prog.Disable(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

// STREAM HANDLERS
func (h *Program) Shell(c echo.Context) error {
	program, err := h.Prog.Shell(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Tail(c echo.Context) error {
	program, err := h.Prog.Tail(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}
