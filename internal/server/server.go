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

package server

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/server/handlers"
)

var jwtSecret = []byte("your-secret-key") // In production, use environment variable

// Claims represents JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// JWTConfig returns JWT middleware configuration
func JWTConfig() echojwt.Config {
	return echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(Claims)
		},
		SigningKey: jwtSecret,
	}
}

// LoginRequest represents login request
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents login response
type LoginResponse struct {
	Token string `json:"token"`
}

// AuthHandler handles authentication
type AuthHandler struct{}

// Login handles user login
func (h *AuthHandler) Login(c echo.Context) error {
	var req LoginRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Simple authentication - in production, validate against database
	if req.Username == "admin" && req.Password == "password" {
		// Create JWT token
		claims := &Claims{
			Username: req.Username,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(jwtSecret)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create token"})
		}

		return c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
	}

	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Invalid credentials"})
}

// Refresh handles token refresh
func (h *AuthHandler) Refresh(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*Claims)

	// Create new token with extended expiration
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	claims.IssuedAt = jwt.NewNumericDate(time.Now())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to refresh token"})
	}

	return c.JSON(http.StatusOK, LoginResponse{Token: tokenString})
}

// Logout handles user logout
func (h *AuthHandler) Logout(c echo.Context) error {
	// In a real implementation, you might want to blacklist the token
	return c.JSON(http.StatusOK, map[string]string{"message": "Logged out successfully"})
}

func New(b internal.Broker) *echo.Echo {
	e := echo.New()
	e.HideBanner = true

	// Add CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// Add logging middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Auth routes (no JWT required)
	auth := e.Group("/api/auth")
	authHandler := &AuthHandler{}
	auth.POST("/login", authHandler.Login)

	// Protected API routes (JWT required)
	api := e.Group("/api")
	api.Use(echojwt.WithConfig(JWTConfig()))

	// Auth routes that require JWT
	api.POST("/auth/refresh", authHandler.Refresh)
	api.POST("/auth/logout", authHandler.Logout)

	// Routes
	ServiceRoutes(e, b)
	CategoryRoutes(e, b)

	return e
}

func ServiceRoutes(e *echo.Echo, b internal.Broker) {
	// Services
	service := handlers.NewService(b)
	svc := e.Group("/service")
	svc.GET("", service.List)
	svc.GET("/:id", service.Get)
	svc.POST("", service.Create)
	svc.PUT("/:id", service.Update)
	svc.DELETE("/:id", service.Delete)
	// svc.OPTIONS("", service.Schema)

	// Runtime handlers
	svc.POST("/:id/start", service.Start)
	svc.POST("/:id/stop", service.Stop)
	svc.POST("/:id/restart", service.Restart)
	svc.POST("/:id/status", service.Status)
	svc.POST("/:id/reload", service.Reload)
	svc.POST("/:id/enable", service.Enable)
	svc.POST("/:id/disable", service.Disable)

	// Stream handlers
	svc.POST("/:id/shell", service.Shell)
	svc.POST("/:id/log", service.Log)
}

func CategoryRoutes(e *echo.Echo, b internal.Broker) {
	category := handlers.NewCategory(b)
	cat := e.Group("/category")
	cat.GET("", category.List)
	cat.GET("/:id", category.Get)
	cat.POST("", category.Create)
	cat.PUT("/:id", category.Update)
	cat.DELETE("/:id", category.Delete)
}
