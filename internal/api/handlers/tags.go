package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/db"
	"github.com/rangertaha/hxe/internal/models"
	"gorm.io/gorm"
)

// Tags handles tag-related API endpoints
type Tags struct {
	DB *gorm.DB
}

// NewTags creates a new tags handler
func NewTags() *Tags {
	return &Tags{
		DB: db.GetDB(),
	}
}

// SetupRoutes configures the tag routes
func (t *Tags) SetupRoutes(e *echo.Group) {
	tags := e.Group("/tags")
	tags.GET("", t.List)
	tags.GET("/:id", t.Get)
	tags.POST("", t.Create)
	tags.PUT("/:id", t.Update)
	tags.DELETE("/:id", t.Delete)
}

// List returns all tags
func (t *Tags) List(ctx echo.Context) error {
	var tags []models.Tag
	if err := t.DB.Find(&tags).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, tags)
}

// Get returns a single tag by ID
func (t *Tags) Get(ctx echo.Context) error {
	var tag models.Tag
	if err := t.DB.First(&tag, ctx.Param("id")).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Tag not found",
		})
	}
	return ctx.JSON(http.StatusOK, tag)
}

// Create creates a new tag
func (t *Tags) Create(ctx echo.Context) error {
	var tag models.Tag
	if err := ctx.Bind(&tag); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := t.DB.Create(&tag).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, tag)
}

// Update updates an existing tag
func (t *Tags) Update(ctx echo.Context) error {
	var tag models.Tag
	if err := t.DB.First(&tag, ctx.Param("id")).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Tag not found",
		})
	}

	if err := ctx.Bind(&tag); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"error": err.Error(),
		})
	}

	if err := t.DB.Save(&tag).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, tag)
}

// Delete deletes a tag
func (t *Tags) Delete(ctx echo.Context) error {
	var tag models.Tag
	if err := t.DB.First(&tag, ctx.Param("id")).Error; err != nil {
		return ctx.JSON(http.StatusNotFound, map[string]string{
			"error": "Tag not found",
		})
	}

	if err := t.DB.Delete(&tag).Error; err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Tag deleted successfully",
	})
} 