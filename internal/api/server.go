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
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rangertaha/hxe/internal"
	"github.com/rangertaha/hxe/internal/api/handlers"
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

	// Initialize CRUD handlers
	proc := handlers.NewProgram(b)
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
