package api

import (
	"github.com/labstack/echo/v4"
	"github.com/rangertaha/hxe/internal/api/handlers"
)

func New() *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Initialize CRUD handlers
	api := e.Group("/api")
	proc := handlers.NewProgram()
	api.GET("/program", proc.List)
	api.GET("/program/:id", proc.Get)
	api.POST("/program", proc.Create)
	api.PUT("/program/:id", proc.Update)
	api.DELETE("/program/:id", proc.Delete)
	api.OPTIONS("/program", proc.Schema)

	// Runtime handlers
	api.POST("/program/:id/start", proc.Start)
	api.POST("/program/:id/stop", proc.Stop)
	api.POST("/program/:id/restart", proc.Restart)
	api.POST("/program/:id/status", proc.Status)
	api.POST("/program/:id/reload", proc.Reload)
	api.POST("/program/:id/enable", proc.Enable)
	api.POST("/program/:id/disable", proc.Disable)

	// Stream handlers
	api.POST("/program/:id/shell", proc.Shell)
	api.POST("/program/:id/tail", proc.Tail)

	return e
}
