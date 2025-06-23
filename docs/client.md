# HXE Client Library

The HXE Go client library provides a comprehensive API for interacting with HXE servers. It supports JWT authentication, program management, and beautiful table formatting.

## Installation

```bash
# Add to your project
go get github.com/rangertaha/hxe/pkg/client

# Or install from GitHub Packages
go get github.com/rangertaha/hxe/pkg/client@latest
```

## Quick Start

```go
package main

import (
    "log"
    "github.com/rangertaha/hxe/pkg/client"
)

func main() {
    // Create authenticated client
    hxeClient := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")
    
    // Login to get JWT token
    _, err := hxeClient.Login()
    if err != nil {
        log.Fatal("Login failed:", err)
    }
    
    // Get program client
    programClient := hxeClient.Program
    
    // List programs
    programs, err := programClient.ListPrograms()
    if err != nil {
        log.Fatal("Failed to list programs:", err)
    }
    
    // Display programs in beautiful table format
    programClient.PrintList(programs)
    
    // Logout when done
    hxeClient.Logout()
}
```

## Client Types

### Authenticated Client

```go
// Create client with authentication
client := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")

// Login to get JWT token
_, err := client.Login()
if err != nil {
    log.Fatal("Login failed:", err)
}

// Use the client...
client.Logout()
```

### Basic Client

```go
// Create basic client without authentication
client := client.NewClient("http://localhost:8080")

// Set token manually if needed
client.SetToken("your-jwt-token-here")
```

## Authentication

### Login

```go
// Login with credentials
_, err := client.Login()
if err != nil {
    log.Fatal("Login failed:", err)
}
```

### Manual Token

```go
// Set token manually
client.SetToken("your-jwt-token-here")

// Check if token is set
if client.HasToken() {
    fmt.Println("Token is set")
}
```

### Refresh Token

```go
// Refresh the current token
_, err := client.RefreshToken()
if err != nil {
    log.Fatal("Token refresh failed:", err)
}
```

### Logout

```go
// Logout and clear token
err := client.Logout()
if err != nil {
    log.Fatal("Logout failed:", err)
}
```

## Program Management

### List Programs

```go
// Get all programs
programs, err := programClient.ListPrograms()
if err != nil {
    log.Fatal("Failed to list programs:", err)
}

// Display in table format
programClient.PrintList(programs)
```

### Get Program

```go
// Get specific program
program, err := programClient.GetProgram("program-id")
if err != nil {
    log.Fatal("Failed to get program:", err)
}

// Display detailed view
programClient.PrintDetail(program)
```

### Create Program

```go
// Create new program
newProgram := &models.Program{
    Name:        "My Program",
    Description: "A test program",
    Command:     "/usr/bin/python3",
    Args:        "script.py",
    Directory:   "/tmp",
    User:        "nobody",
    Group:       "nobody",
    Enabled:     true,
    AutoRestart: true,
    MaxRestarts: 3,
}

created, err := programClient.CreateProgram(newProgram)
if err != nil {
    log.Fatal("Failed to create program:", err)
}

fmt.Printf("Created program: %s\n", created.ID)
```

### Update Program

```go
// Update existing program
program.Name = "Updated Program Name"
program.Description = "Updated description"

updated, err := programClient.UpdateProgram(program.ID, program)
if err != nil {
    log.Fatal("Failed to update program:", err)
}

fmt.Printf("Updated program: %s\n", updated.ID)
```

### Delete Program

```go
// Delete program
err := programClient.DeleteProgram("program-id")
if err != nil {
    log.Fatal("Failed to delete program:", err)
}

fmt.Println("Program deleted successfully")
```

## Runtime Operations

### Start Program

```go
// Start a program
_, err := programClient.StartProgram("program-id")
if err != nil {
    log.Fatal("Failed to start program:", err)
}

fmt.Println("Program started successfully")
```

### Stop Program

```go
// Stop a program
_, err := programClient.StopProgram("program-id")
if err != nil {
    log.Fatal("Failed to stop program:", err)
}

fmt.Println("Program stopped successfully")
```

### Restart Program

```go
// Restart a program
_, err := programClient.RestartProgram("program-id")
if err != nil {
    log.Fatal("Failed to restart program:", err)
}

fmt.Println("Program restarted successfully")
```

### Enable/Disable Autostart

```go
// Enable autostart
_, err := programClient.EnableAutostart("program-id")
if err != nil {
    log.Fatal("Failed to enable autostart:", err)
}

// Disable autostart
_, err = programClient.DisableAutostart("program-id")
if err != nil {
    log.Fatal("Failed to disable autostart:", err)
}
```

## Display Methods

### Print List

```go
// Display programs in table format
programs, _ := programClient.ListPrograms()
programClient.PrintList(programs)
```

### Print Detail

```go
// Display detailed program information
program, _ := programClient.GetProgram("program-id")
programClient.PrintDetail(program)
```

## Error Handling

The client library provides comprehensive error handling:

```go
// Check for specific error types
if err != nil {
    switch {
    case client.IsUnauthorized(err):
        fmt.Println("Authentication required")
    case client.IsNotFound(err):
        fmt.Println("Program not found")
    case client.IsServerError(err):
        fmt.Println("Server error occurred")
    default:
        fmt.Printf("Unexpected error: %v\n", err)
    }
}
```

## Configuration

### Custom HTTP Client

```go
// Create custom HTTP client
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}

// Create client with custom HTTP client
client := client.NewClientWithHTTP("http://localhost:8080", httpClient)
```

### Custom Headers

```go
// Add custom headers
client.SetHeader("X-Custom-Header", "custom-value")

// Set multiple headers
headers := map[string]string{
    "X-API-Version": "1.0",
    "X-Client-ID":   "my-app",
}
client.SetHeaders(headers)
```

## Examples

### Complete Example

```go
package main

import (
    "log"
    "time"
    "github.com/rangertaha/hxe/pkg/client"
)

func main() {
    // Create authenticated client
    hxeClient := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")
    
    // Login
    _, err := hxeClient.Login()
    if err != nil {
        log.Fatal("Login failed:", err)
    }
    defer hxeClient.Logout()
    
    programClient := hxeClient.Program
    
    // List existing programs
    programs, err := programClient.ListPrograms()
    if err != nil {
        log.Fatal("Failed to list programs:", err)
    }
    
    fmt.Printf("Found %d programs\n", len(programs))
    programClient.PrintList(programs)
    
    // Create a new program
    newProgram := &models.Program{
        Name:        "Test Program",
        Description: "A test program for demonstration",
        Command:     "/usr/bin/python3",
        Args:        "-c 'import time; time.sleep(10); print(\"Hello from HXE!\")'",
        Directory:   "/tmp",
        User:        "nobody",
        Group:       "nobody",
        Enabled:     true,
        AutoRestart: true,
        MaxRestarts: 3,
    }
    
    created, err := programClient.CreateProgram(newProgram)
    if err != nil {
        log.Fatal("Failed to create program:", err)
    }
    
    fmt.Printf("Created program: %s\n", created.ID)
    
    // Start the program
    _, err = programClient.StartProgram(created.ID)
    if err != nil {
        log.Fatal("Failed to start program:", err)
    }
    
    fmt.Println("Program started")
    
    // Wait a bit and check status
    time.Sleep(2 * time.Second)
    
    program, err := programClient.GetProgram(created.ID)
    if err != nil {
        log.Fatal("Failed to get program:", err)
    }
    
    programClient.PrintDetail(program)
    
    // Stop the program
    _, err = programClient.StopProgram(created.ID)
    if err != nil {
        log.Fatal("Failed to stop program:", err)
    }
    
    fmt.Println("Program stopped")
    
    // Clean up
    err = programClient.DeleteProgram(created.ID)
    if err != nil {
        log.Fatal("Failed to delete program:", err)
    }
    
    fmt.Println("Program deleted")
}
```

## API Reference

### Client Methods

- `NewClient(baseURL string) *Client`
- `NewAuthenticatedClient(baseURL, username, password string) *Client`
- `NewClientWithHTTP(baseURL string, httpClient *http.Client) *Client`
- `Login() (*models.AuthResponse, error)`
- `Logout() error`
- `RefreshToken() (*models.AuthResponse, error)`
- `SetToken(token string)`
- `GetToken() string`
- `HasToken() bool`
- `SetHeader(key, value string)`
- `SetHeaders(headers map[string]string)`

### Program Client Methods

- `ListPrograms() ([]*models.Program, error)`
- `GetProgram(id string) (*models.Program, error)`
- `CreateProgram(program *models.Program) (*models.Program, error)`
- `UpdateProgram(id string, program *models.Program) (*models.Program, error)`
- `DeleteProgram(id string) error`
- `StartProgram(id string) (*models.Program, error)`
- `StopProgram(id string) (*models.Program, error)`
- `RestartProgram(id string) (*models.Program, error)`
- `EnableAutostart(id string) (*models.Program, error)`
- `DisableAutostart(id string) (*models.Program, error)`
- `PrintList(programs []*models.Program)`
- `PrintDetail(program *models.Program)`

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 