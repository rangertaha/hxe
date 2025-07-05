package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/db"
	"github.com/rangertaha/hxe/internal/models"
	"gorm.io/gorm"
)

// Services handles service-related API endpoints
type Services struct {
	DB *gorm.DB
}

// NewServices creates a new services handler
func NewServices() *Services {
	return &Services{
		DB: db.GetDB(),
	}
}

// SetupRoutes configures the service routes
func (s *Services) SetupRoutes(e *echo.Group) {
	services := e.Group("/services")
	services.GET("", s.List)
	services.GET("/:id", s.Get)
	services.POST("", s.Create)
	services.PUT("/:id", s.Update)
	services.DELETE("/:id", s.Delete)
}

// List returns all services
func (s *Services) List(c echo.Context) error {
	var services []models.Service
	if err := s.DB.Find(&services).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, services)
}

// Get returns a single service by ID
func (s *Services) Get(c echo.Context) error {
	var service models.Service
	if err := s.DB.First(&service, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Service not found",
		})
	}
	return c.JSON(http.StatusOK, service)
}

// Create creates a new service
func (s *Services) Create(c echo.Context) error {
	var service models.Service
	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := s.DB.Create(&service).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, service)
}

// Update updates an existing service
func (s *Services) Update(c echo.Context) error {
	var service models.Service
	if err := s.DB.First(&service, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Service not found",
		})
	}

	if err := c.Bind(&service); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := s.DB.Save(&service).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, service)
}

// Delete deletes a service
func (s *Services) Delete(c echo.Context) error {
	var service models.Service
	if err := s.DB.First(&service, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Service not found",
		})
	}

	if err := s.DB.Delete(&service).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Service deleted successfully",
	})
}
