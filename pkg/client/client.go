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

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	DefaultBaseURL = "http://localhost:8080"
)

// Client represents an API client for interacting with the hxe API
type Client struct {
	baseURL    string
	httpClient *http.Client
	token      string
	username   string
	password   string
	Program    *ProgramClient
}

// // NewClient creates a new API client
// func NewClient(baseURL string) *Client {
// 	return &Client{
// 		baseURL: baseURL,
// 		httpClient: &http.Client{
// 			Timeout: time.Second * 30,
// 		},
// 	}
// }

// NewAuthenticatedClient creates a new API client with authentication
func NewClient(baseURL, username, password string) *Client {
	c := &Client{
		baseURL:  baseURL,
		username: username,
		password: password,
		httpClient: &http.Client{
			Timeout: time.Second * 30,
		},
	}
	c.Program = NewProgramClient(c)
	return c
}

// SetToken sets the JWT token for authentication
func (c *Client) SetToken(token string) {
	c.token = token
}

// GetToken returns the current JWT token
func (c *Client) GetToken() string {
	return c.token
}

// Login authenticates with the server and retrieves a JWT token
func (c *Client) Login() (*Client, error) {
	loginData := map[string]string{
		"username": c.username,
		"password": c.password,
	}

	var response struct {
		Token string `json:"token"`
	}

	err := c.Post("/api/auth/login", loginData, &response)
	if err != nil {
		return c, fmt.Errorf("login failed: %w", err)
	}

	c.token = response.Token
	return c, nil
}

// Logout clears the current token
func (c *Client) Logout() {
	c.token = ""
}

// RefreshToken refreshes the JWT token
func (c *Client) RefreshToken() error {
	var response struct {
		Token string `json:"token"`
	}

	err := c.Post("/api/auth/refresh", nil, &response)
	if err != nil {
		return fmt.Errorf("token refresh failed: %w", err)
	}

	c.token = response.Token
	return nil
}

// Generic HTTP methods

func (c *Client) Get(path string, v interface{}) error {
	return c.DoRequest(http.MethodGet, path, nil, v)
}

func (c *Client) Post(path string, body interface{}, v interface{}) error {
	return c.DoRequest(http.MethodPost, path, body, v)
}

func (c *Client) Put(path string, body interface{}, v interface{}) error {
	return c.DoRequest(http.MethodPut, path, body, v)
}

func (c *Client) Delete(path string, v interface{}) error {
	return c.DoRequest(http.MethodDelete, path, nil, v)
}

func (c *Client) DoRequest(method, path string, body interface{}, v interface{}) error {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return err
		}
	}

	req, err := http.NewRequest(method, c.baseURL+path, &buf)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	// Add JWT token if available
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Handle authentication errors
	if resp.StatusCode == http.StatusUnauthorized {
		return fmt.Errorf("authentication failed: invalid credentials or expired token")
	}

	if resp.StatusCode >= 400 {
		var errResp struct {
			Error string `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
			return fmt.Errorf("HTTP error: %d", resp.StatusCode)
		}
		return fmt.Errorf("API error: %s", errResp.Error)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return err
		}
	}

	return nil
}

// func (c *Client) Program() *ProgramClient {
// 	return &ProgramClient{c}
// }
