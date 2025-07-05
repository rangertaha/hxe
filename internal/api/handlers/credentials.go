package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/db"
	"github.com/rangertaha/hxe/internal/models"
	"gorm.io/gorm"
)

// Credentials handles credential-related API endpoints
type Credentials struct {
	DB *gorm.DB
}

// NewCredentials creates a new credentials handler
func NewCredentials() *Credentials {
	return &Credentials{
		DB: db.GetDB(),
	}
}

// SetupRoutes configures the credential routes
func (c *Credentials) SetupRoutes(e *echo.Group) {
	credentials := e.Group("/credentials")
	credentials.GET("", c.List)
	credentials.GET("/:id", c.Get)
	credentials.POST("", c.Create)
	credentials.PUT("/:id", c.Update)
	credentials.DELETE("/:id", c.Delete)
}

// List returns all credentials
func (c *Credentials) List(ctx echo.Context) error {
	var credentials []models.Credential
	if err := c.DB.Preload("Tags").Preload("Group").Find(&credentials).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, credentials)
}

// Get returns a single credential by ID
func (c *Credentials) Get(ctx echo.Context) error {
	var credential models.Credential
	if err := c.DB.Preload("Tags").Preload("Group").First(&credential, ctx.Param("id")).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Credential not found",
		})
	}
	return ctx.JSON(http.StatusOK, credential)
}

// Create creates a new credential
func (c *Credentials) Create(ctx echo.Context) error {
	var credential models.Credential
	if err := ctx.Bind(&credential); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := c.DB.Create(&credential).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Reload with associations
	c.DB.Preload("Tags").Preload("Group").First(&credential, credential.ID)

	return ctx.JSON(http.StatusCreated, credential)
}

// Update updates an existing credential
func (c *Credentials) Update(ctx echo.Context) error {
	var credential models.Credential
	if err := c.DB.First(&credential, ctx.Param("id")).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Credential not found",
		})
	}

	if err := ctx.Bind(&credential); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := c.DB.Save(&credential).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	// Reload with associations
	c.DB.Preload("Tags").Preload("Group").First(&credential, credential.ID)

	return ctx.JSON(http.StatusOK, credential)
}

// Delete deletes a credential
func (c *Credentials) Delete(ctx echo.Context) error {
	var credential models.Credential
	if err := c.DB.First(&credential, ctx.Param("id")).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Credential not found",
		})
	}

	if err := c.DB.Delete(&credential).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Credential deleted successfully",
	})
}
