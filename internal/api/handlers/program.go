package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/api/services"
	"github.com/rangertaha/hxe/internal/models"
)

type Program struct {
	Svc *services.Program
}

func NewProgram() *Program {
	return &Program{
		Svc: services.NewProgram(),
	}
}

// CRUD HANDLERS
func (h *Program) List(c echo.Context) error {
	programs, err := h.Svc.List()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, programs)
}

func (h *Program) Get(c echo.Context) error {
	program, err := h.Svc.Get(c.Param("id"))
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
	program, err := h.Svc.Create(prog)
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
	program, err := h.Svc.Update(c.Param("id"), prog)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Delete(c echo.Context) error {
	err := h.Svc.Delete(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Program deleted"})
}

func (h *Program) Schema(c echo.Context) error {
	schema, err := h.Svc.Schema()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, schema)
}

// RUNTIME HANDLERS
func (h *Program) Start(c echo.Context) error {
	program, err := h.Svc.Start(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Stop(c echo.Context) error {
	program, err := h.Svc.Stop(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Restart(c echo.Context) error {
	program, err := h.Svc.Restart(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Status(c echo.Context) error {
	program, err := h.Svc.Status(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, program)
}

func (h *Program) Reload(c echo.Context) error {
	err := h.Svc.Reload(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Program reloaded"})
}

func (h *Program) Enable(c echo.Context) error {
	err := h.Svc.Enable(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Program enabled"})
}

func (h *Program) Disable(c echo.Context) error {
	err := h.Svc.Disable(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Program disabled"})
}

// STREAM HANDLERS
func (h *Program) Shell(c echo.Context) error {
	err := h.Svc.Shell(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Shell started"})
}

func (h *Program) Tail(c echo.Context) error {
	err := h.Svc.Tail(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Tail started"})
}

func (h *Program) Run(c echo.Context) error {
	err := h.Svc.Run(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "Program run started"})
}
