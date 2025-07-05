package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/db"
	"github.com/rangertaha/hxe/internal/models"
	"gorm.io/gorm"
)

// Groups handles group-related API endpoints
type Groups struct {
	DB *gorm.DB
}

// NewGroups creates a new groups handler
func NewGroups() *Groups {
	return &Groups{
		DB: db.GetDB(),
	}
}

// SetupRoutes configures the group routes
func (g *Groups) SetupRoutes(e *echo.Group) {
	groups := e.Group("/groups")
	groups.GET("", g.List)
	groups.GET("/:id", g.Get)
	groups.POST("", g.Create)
	groups.PUT("/:id", g.Update)
	groups.DELETE("/:id", g.Delete)
}

// List returns all groups
func (g *Groups) List(c echo.Context) error {
	var groups []models.Group
	if err := g.DB.Preload("Children").Preload("Parent").Find(&groups).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, groups)
}

// Get returns a single group by ID
func (g *Groups) Get(c echo.Context) error {
	var group models.Group
	if err := g.DB.Preload("Children").Preload("Parent").First(&group, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Group not found",
		})
	}
	return c.JSON(http.StatusOK, group)
}

// Create creates a new group
func (g *Groups) Create(c echo.Context) error {
	var group models.Group
	if err := c.Bind(&group); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := g.DB.Create(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, group)
}

// Update updates an existing group
func (g *Groups) Update(c echo.Context) error {
	var group models.Group
	if err := g.DB.First(&group, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Group not found",
		})
	}

	if err := c.Bind(&group); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := g.DB.Save(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, group)
}

// Delete deletes a group
func (g *Groups) Delete(c echo.Context) error {
	var group models.Group
	if err := g.DB.First(&group, c.Param("id")).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Group not found",
		})
	}

	if err := g.DB.Delete(&group).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Group deleted successfully",
	})
}
