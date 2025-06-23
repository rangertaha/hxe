# HXE Client Library

A Go client library for interacting with the HXE API server with JWT authentication support.

## License

This project is licensed under the GNU General Public License v3.0 or later - see the [LICENSE](../../LICENSE) file for details.

## Installation

```bash
go get github.com/rangertaha/hxe/pkg/client
```

## Quick Start

### Basic Usage (without authentication)

```go
package main

import (
    "log"
    "github.com/rangertaha/hxe/pkg/client"
)

func main() {
    // Create client
    hxeClient := client.NewClient("http://localhost:8080")
    
    // Get program client
    programClient := hxeClient.Program
    
    // List programs
    programs, err := programClient.ListPrograms()
    if err != nil {
        log.Fatal("Failed to list programs:", err)
    }
    
    // Display programs
    programClient.PrintList(programs)
}
```

### Usage with JWT Authentication

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

## API Reference

### Creating Clients

#### NewClient
```go
client := client.NewClient("http://localhost:8080")
```
Creates a basic client without authentication.

#### NewAuthenticatedClient
```go
client := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")
```
Creates a client with authentication credentials.

### Authentication

#### Login
```go
_, err := client.Login()
```
Logs in with the provided credentials and retrieves a JWT token.

#### Logout
```go
err := client.Logout()
```
Logs out and clears the JWT token.

#### RefreshToken
```go
_, err := client.RefreshToken()
```
Refreshes the current JWT token.

#### SetToken
```go
client.SetToken("your-jwt-token-here")
```
Manually sets a JWT token.

### Program Management

#### ListPrograms
```go
programs, err := programClient.ListPrograms()
```
Retrieves all programs.

#### GetProgram
```go
program, err := programClient.GetProgram("program-id")
```
Retrieves a specific program by ID.

#### CreateProgram
```go
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
```
Creates a new program.

#### UpdateProgram
```go
program.Name = "Updated Name"
updated, err := programClient.UpdateProgram(program.ID, program)
```
Updates an existing program.

#### DeleteProgram
```go
err := programClient.DeleteProgram("program-id")
```
Deletes a program.

#### StartProgram
```go
_, err := programClient.StartProgram("program-id")
```
Starts a program.

#### StopProgram
```go
_, err := programClient.StopProgram("program-id")
```
Stops a program.

#### RestartProgram
```go
_, err := programClient.RestartProgram("program-id")
```
Restarts a program.

#### EnableAutostart
```go
_, err := programClient.EnableAutostart("program-id")
```
Enables autostart for a program.

#### DisableAutostart
```go
_, err := programClient.DisableAutostart("program-id")
```
Disables autostart for a program.

### Display Methods

#### PrintList
```go
programClient.PrintList(programs)
```
Displays programs in a beautiful table format.

#### PrintDetail
```go
programClient.PrintDetail(program)
```
Displays detailed information about a program.

## Authentication

The client supports JWT authentication with the following flow:

1. **Login**: Call `Login()` with credentials to get a JWT token
2. **Automatic Token Usage**: The client automatically includes the token in `Authorization: Bearer <token>` header
3. **Token Refresh**: Use `RefreshToken()` to refresh expired tokens
4. **Logout**: Call `Logout()` to clear the token

### Default Credentials
- **Username**: `admin`
- **Password**: `password`

### Manual Token Management
```go
// Set token manually
client.SetToken("your-jwt-token-here")

// Check if token is set
if client.HasToken() {
    fmt.Println("Token is set")
}

// Get current token
token := client.GetToken()
```

## Program Model

The `models.Program` struct contains the following fields:

```go
type Program struct {
    ID          string    `json:"id"`
    Name        string    `json:"name"`
    Description string    `json:"description"`
    Command     string    `json:"command"`
    Args        string    `json:"args"`
    Directory   string    `json:"directory"`
    User        string    `json:"user"`
    Group       string    `json:"group"`
    Enabled     bool      `json:"enabled"`
    AutoRestart bool      `json:"auto_restart"`
    MaxRestarts int       `json:"max_restarts"`
    Status      string    `json:"status"`
    PID         int       `json:"pid"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
```

## Error Handling

The client may return various types of errors:

- **Authentication errors**: Invalid credentials, expired tokens
- **Network errors**: Connection issues, timeouts
- **API errors**: Server errors, validation errors
- **Not found errors**: Programs that don't exist

```go
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

## Examples

See `examples/client/programs/main.go` for complete examples of client usage. 