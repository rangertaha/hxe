package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/db"
	"gorm.io/gorm"
)

// GenericHandler provides common CRUD operations for any GORM model
type GenericHandler struct {
	DB *gorm.DB
}

// NewGenericHandler creates a new generic handler
func NewGenericHandler() *GenericHandler {
	return &GenericHandler{
		DB: db.GetDB(),
	}
}

// List handles GET requests to list all records
func (h *GenericHandler) List(model interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := h.DB.Find(model).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusOK, model)
	}
}

// Get handles GET requests to retrieve a single record by ID
func (h *GenericHandler) Get(model interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid ID",
			})
		}

		if err := h.DB.First(model, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "Record not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, model)
	}
}

// Create handles POST requests to create a new record
func (h *GenericHandler) Create(model interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := c.Bind(model); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		if err := h.DB.Create(model).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, model)
	}
}

// Update handles PUT requests to update an existing record
func (h *GenericHandler) Update(model interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid ID",
			})
		}

		// First find the existing record
		if err := h.DB.First(model, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "Record not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		if err := c.Bind(model); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": err.Error(),
			})
		}

		if err := h.DB.Save(model).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, model)
	}
}

// Delete handles DELETE requests to remove a record
func (h *GenericHandler) Delete(model interface{}) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.ParseUint(c.Param("id"), 10, 32)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid ID",
			})
		}

		// First find the record to ensure it exists
		if err := h.DB.First(model, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, map[string]string{
					"error": "Record not found",
				})
			}
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		// Then delete it
		if err := h.DB.Delete(model).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]string{
			"message": "Record deleted successfully",
		})
	}
}
