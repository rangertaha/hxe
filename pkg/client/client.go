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

// import (
// 	"bytes"
// 	"context"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strings"
// 	"time"

// 	"github.com/hashicorp/hcl/v2/hclsimple"
// 	"github.com/rangertaha/hxe/internal/config"
// 	"github.com/urfave/cli/v3"

// 	_ "embed"
// )

// const (
// 	DefaultBaseURL  = "http://0.0.0.0:8080"
// 	DefaultUsername = "admin"
// 	DefaultPassword = "password"
// 	DefaultTimeout  = 30
// )

// //go:embed client.hcl
// var DefaultConfig []byte

// type (
// 	Config struct {
// 		ID         string  `hcl:"id,optional"`
// 		Debug      bool    `hcl:"debug,optional"`
// 		Version    string  `hcl:"version,optional"`
// 		Server     *Server `hcl:"server,block"`
// 		configPath string
// 	}

// 	// Client represents an API client for interacting with the hxe API
// 	Server struct {
// 		Host     string `hcl:"addr,optional"`
// 		Port     int    `hcl:"port,optional"`
// 		Username string `hcl:"username,optional"`
// 		Password string `hcl:"password,optional"`
// 		Token    string `hcl:"token,optional"`
// 		URL      string `hcl:"url,optional"`
// 		Timeout  int    `hcl:"timeout,optional"`
// 	}

// 	// Client represents an API client for interacting with the hxe API
// 	Client struct {
// 		config     *Config
// 		httpClient *http.Client

// 		Programs *ProgramClient
// 		// Credentials *CredentialClient
// 		// Variables   *VariableClient
// 		// Fields      *FieldClient
// 		// Apps        *AppClient

// 	}
// )

// // New creates a new client
// func New(options ...func(*Config) error) (*Client, error) {
// 	cfg := &Config{
// 		Server: &Server{
// 			URL:      DefaultBaseURL,
// 			Username: DefaultUsername,
// 			Password: DefaultPassword,
// 			Timeout:  DefaultTimeout,
// 		},
// 	}
// 	// Apply config options
// 	for _, opt := range options {
// 		err := opt(cfg)
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if cfg.Server.Host != "" && cfg.Server.Port != 0 {
// 		cfg.Server.URL = fmt.Sprintf("http://%s:%d", cfg.Server.Host, cfg.Server.Port)
// 	}

// 	c := &Client{
// 		config: cfg,
// 		httpClient: &http.Client{
// 			Timeout: time.Duration(cfg.Server.Timeout) * time.Second,
// 		},
// 	}

// 	c.Programs = &ProgramClient{client: c}
// 	// c.Credentials = &CredentialClient{client: c}
// 	// c.Variables = &VariableClient{client: c}
// 	// c.Fields = &FieldClient{client: c}
// 	// c.Apps = &AppClient{client: c}
// 	return c, nil
// }

// // func Options(options ...func(*Config) error) (*Config, error) {
// // 	c := &Config{
// // 		Server: &Server{
// // 			URL:      DefaultBaseURL,
// // 			Username: DefaultUsername,
// // 			Password: DefaultPassword,
// // 			Timeout:  DefaultTimeout,
// // 			Token:    "",
// // 		},
// // 	}

// // 	// Apply config options
// // 	for _, opt := range options {
// // 		err := opt(c)
// // 		if err != nil {
// // 			return nil, err
// // 		}
// // 	}

// // 	return c, nil
// // }

// func CliOptions(ctx context.Context, cmd *cli.Command) func(c *Config) error {
// 	return func(c *Config) error {
// 		if cmd.String("config") != "" {
// 			// Try to load connection details from config file
// 			if err := FileOptions(cmd.String("config"))(c); err != nil {
// 				return err
// 			}
// 		}

// 		if cmd.String("username") != "" {
// 			c.Server.Username = cmd.String("username")
// 		}

// 		if cmd.String("password") != "" {
// 			c.Server.Password = cmd.String("password")
// 		}

// 		if cmd.String("url") != "" {
// 			c.Server.URL = cmd.String("url")
// 		}

// 		if cmd.Duration("timeout").Seconds() > 1 {
// 			c.Server.Timeout = int(cmd.Duration("timeout").Seconds())
// 		}

// 		return nil
// 	}
// }

// func FileOptions(path string) func(*Config) error {
// 	return func(c *Config) (err error) {
// 		// Check if config file exists
// 		if _, err = os.Stat(path); os.IsNotExist(err) {
// 			if err = createConfig(path, DefaultConfig); err != nil {
// 				return
// 			}
// 		}

// 		c.configPath = path

// 		if err = hclsimple.DecodeFile(path, config.CtxFunctions, c); err != nil {
// 			return fmt.Errorf("error parsing config file: %w", err)
// 		}
// 		return
// 	}
// }

// func createConfig(filename string, contents []byte) error {
// 	// Ask user if they want to create a config file
// 	fmt.Printf("Config file does not exist: %s\n", filename)
// 	fmt.Println("Do you want to create a config file? (y/n)")
// 	var answer string
// 	fmt.Scanln(&answer)
// 	if strings.ToLower(answer) != "y" {
// 		return fmt.Errorf("config file does not exist")
// 	}

// 	// Check if file exists
// 	_, err := os.Stat(filename)
// 	if os.IsNotExist(err) {
// 		// Create directory if it doesn't exist
// 		dir := filepath.Dir(filename)
// 		if err := os.MkdirAll(dir, 0755); err != nil {
// 			return fmt.Errorf("failed to create config directory: %w", err)
// 		}

// 		// Write default config to file
// 		if err := os.WriteFile(filename, contents, 0644); err != nil {
// 			return fmt.Errorf("failed to create default config file: %w", err)
// 		}

// 		fmt.Printf("Created default configuration file: %s\n", filename)

// 	} else if err != nil {
// 		return fmt.Errorf("error checking config file: %w", err)
// 	}

// 	return nil
// }

// // func (c *Client) SaveToken() {
// // 	f := hclwrite.NewEmptyFile()
// // 	gohcl.EncodeIntoBody(&c.config, f.Body())
// // 	fmt.Printf("%s", f.Bytes())

// // 	if c.config.configPath != "" {
// // 		os.WriteFile(c.config.configPath+".new", f.Bytes(), 0644)
// // 	}
// // }

// // func (c *Config) SaveToken() {
// // 	f := hclwrite.NewEmptyFile()
// // 	gohcl.EncodeIntoBody(&c.config, f.Body())
// // 	fmt.Printf("%s", f.Bytes())

// // 	if c.config.configPath != "" {
// // 		os.WriteFile(c.config.configPath+".new", f.Bytes(), 0644)
// // 	}
// // }

// // SetToken sets the JWT token for authentication
// func (c *Client) SetToken(token string) {
// 	c.config.Server.Token = token
// 	// c.SaveToken()
// }

// // GetToken returns the current JWT token
// func (c *Client) GetToken() string {
// 	return c.config.Server.Token
// }

// // Login authenticates with the server and retrieves a JWT token
// func (c *Client) Login() (*Client, error) {
// 	loginData := map[string]string{
// 		"username": c.config.Server.Username,
// 		"password": c.config.Server.Password,
// 	}

// 	var response struct {
// 		Token string `json:"token"`
// 	}

// 	err := c.Post("/api/auth/login", loginData, &response)
// 	if err != nil {
// 		return c, fmt.Errorf("login failed: %w", err)
// 	}

// 	c.config.Server.Token = response.Token
// 	// c.SaveToken()
// 	return c, nil
// }

// // Logout clears the current token
// func (c *Client) Logout() {
// 	c.config.Server.Token = ""
// }

// // RefreshToken refreshes the JWT token
// func (c *Client) RefreshToken() error {
// 	var response struct {
// 		Token string `json:"token"`
// 	}

// 	err := c.Post("/api/auth/refresh", nil, &response)
// 	if err != nil {
// 		return fmt.Errorf("token refresh failed: %w", err)
// 	}

// 	c.config.Server.Token = response.Token
// 	return nil
// }

// // Generic HTTP methods

// func (c *Client) Get(path string, v interface{}) error {
// 	return c.DoRequest(http.MethodGet, path, nil, v)
// }

// func (c *Client) Post(path string, body interface{}, v interface{}) error {
// 	return c.DoRequest(http.MethodPost, path, body, v)
// }

// func (c *Client) Put(path string, body interface{}, v interface{}) error {
// 	return c.DoRequest(http.MethodPut, path, body, v)
// }

// func (c *Client) Delete(path string, v interface{}) error {
// 	return c.DoRequest(http.MethodDelete, path, nil, v)
// }

// func (c *Client) DoRequest(method, path string, body interface{}, v interface{}) error {
// 	var buf bytes.Buffer
// 	if body != nil {
// 		if err := json.NewEncoder(&buf).Encode(body); err != nil {
// 			return err
// 		}
// 	}

// 	req, err := http.NewRequest(method, c.config.Server.URL+path, &buf)
// 	if err != nil {
// 		return err
// 	}

// 	req.Header.Set("Content-Type", "application/json")

// 	// Add JWT token if available
// 	if c.config.Server.Token != "" {
// 		req.Header.Set("Authorization", "Bearer "+c.config.Server.Token)
// 	}

// 	resp, err := c.httpClient.Do(req)
// 	if err != nil {
// 		return err
// 	}
// 	defer resp.Body.Close()

// 	// Handle authentication errors
// 	if resp.StatusCode == http.StatusUnauthorized {
// 		return fmt.Errorf("authentication failed: invalid credentials or expired token")
// 	}

// 	if resp.StatusCode >= 400 {
// 		var errResp struct {
// 			Error string `json:"error"`
// 		}
// 		if err := json.NewDecoder(resp.Body).Decode(&errResp); err != nil {
// 			return fmt.Errorf("HTTP error: %d", resp.StatusCode)
// 		}
// 		return fmt.Errorf("API error: %s", errResp.Error)
// 	}

// 	if v != nil {
// 		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// // func (c *Client) Service() *ServiceClient {
// // 	return &ServiceClient{c}
// // }
