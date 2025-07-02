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

package api

import (
	"fmt"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rangertaha/hxe/internal/api/handlers"
	"github.com/rangertaha/hxe/internal/client"
	"github.com/rangertaha/hxe/internal/config"
	"github.com/rangertaha/hxe/internal/log"
	"github.com/rs/zerolog"
)

type Server struct {
	config *config.API
	log    zerolog.Logger
	client *client.Client
	router *echo.Echo
	done   chan bool
}

// New creates a new API server
func New(cfg *config.API) (server *Server, err error) {
	server = &Server{
		config: cfg,
		client: client.New(cfg.NC),
		log:    log.With().Logger(),
		done:   make(chan bool),
	}

	if err := server.Init(); err != nil {
		return nil, err
	}

	return
}

func (a *Server) Init() error {
	a.log.Info().Msg("initializing api")

	a.router = echo.New()

	// Load routes
	if err := a.Load(a.config); err != nil {
		return err
	}

	return nil
}

func (a *Server) Load(cfg *config.API) error {
	a.router.HideBanner = true

	// Add CORS middleware
	a.router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Add logging middleware
	a.router.Use(middleware.Logger())
	a.router.Use(middleware.Recover())

	// Auth routes (no JWT required)
	auth := a.router.Group("/api/auth")
	authHandler := &AuthHandler{}
	auth.POST("/login", authHandler.Login)

	// Protected API routes (JWT required)
	api := a.router.Group("/api")
	api.Use(echojwt.WithConfig(JWTConfig()))

	// Auth routes that require JWT
	api.POST("/auth/refresh", authHandler.Refresh)
	api.POST("/auth/logout", authHandler.Logout)

	// Routes
	handlers.ServiceRoutes(api, a.client)
	handlers.CategoryRoutes(api, a.client)

	return nil
}

// Start starts the task
func (a *Server) Start() (err error) {
	a.log.Info().Msg("starting api")

	// Start server on port
	a.log.Info().Str("host", a.config.Host).Int("port", a.config.Port).Msg("starting api")
	if err = a.router.Start(fmt.Sprintf("%s:%d", a.config.Host, a.config.Port)); err != nil {
		a.log.Fatal().Err(err).Msg("failed to start api")
	}
	return
}

func (a *Server) Stop() error {
	a.router.Close()
	return nil
}

func (a *Server) Restart() error {
	a.log.Info().Msg("restarting api")

	err := a.Stop()
	if err != nil {
		a.log.Error().Err(err).Msg("failed to stop api")
		return err
	}
	return a.Start()
}
