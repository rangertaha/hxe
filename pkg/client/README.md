# HXE Client

A Go client library for interacting with the HXE (HashiCorp Executor) API server with JWT authentication support.

## Installation

```bash
go get github.com/rangertaha/hxe/pkg/client
```

## Quick Start

### Basic Usage (No Authentication)

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/rangertaha/hxe/pkg/client"
)

func main() {
    // Create a new client
    client := client.NewClient("http://localhost:8080")
    
    // Get the program client
    programClient := client.Program()
    
    // List all programs
    programs, err := programClient.ListPrograms()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d programs\n", len(programs))
}
```

### With JWT Authentication

```go
package main

import (
    "fmt"
    "log"
    
    "github.com/rangertaha/hxe/pkg/client"
)

func main() {
    // Create an authenticated client
    client := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")
    
    // Login to get JWT token
    err := client.Login()
    if err != nil {
        log.Fatal("Failed to login:", err)
    }
    
    // Get the program client
    programClient := client.Program()
    
    // Use authenticated API calls
    programs, err := programClient.ListPrograms()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Found %d programs\n", len(programs))
    
    // Logout when done
    client.Logout()
}
```

## API Reference

### Client

#### `NewClient(baseURL string) *Client`
Creates a new HXE client instance without authentication.

**Parameters:**
- `baseURL` - The base URL of the HXE server (e.g., "http://localhost:8080")

**Returns:**
- `*Client` - A new client instance

#### `NewAuthenticatedClient(baseURL, username, password string) *Client`
Creates a new HXE client instance with authentication credentials.

**Parameters:**
- `baseURL` - The base URL of the HXE server
- `username` - Username for authentication
- `password` - Password for authentication

**Returns:**
- `*Client` - A new authenticated client instance

#### `client.Login() error`
Authenticates with the server and retrieves a JWT token.

**Returns:**
- `error` - Any error that occurred during login

#### `client.Logout()`
Clears the current JWT token.

#### `client.RefreshToken() error`
Refreshes the JWT token.

**Returns:**
- `error` - Any error that occurred during token refresh

#### `client.SetToken(token string)`
Sets a JWT token manually.

**Parameters:**
- `token` - The JWT token string

#### `client.GetToken() string`
Returns the current JWT token.

**Returns:**
- `string` - The current JWT token

#### `client.Program() *ProgramClient`
Returns a program client for managing programs.

### ProgramClient

#### `ListPrograms() ([]models.Program, error)`
Retrieves all programs from the server.

**Returns:**
- `[]models.Program` - List of programs
- `error` - Any error that occurred

#### `GetProgram(id uint) (*models.Program, error)`
Retrieves a specific program by ID.

**Parameters:**
- `id` - The program ID

**Returns:**
- `*models.Program` - The program
- `error` - Any error that occurred

#### `CreateProgram(program *models.Program) (*models.Program, error)`
Creates a new program.

**Parameters:**
- `program` - The program to create

**Returns:**
- `*models.Program` - The created program
- `error` - Any error that occurred

#### `UpdateProgram(id uint, program *models.Program) (*models.Program, error)`
Updates an existing program.

**Parameters:**
- `id` - The program ID
- `program` - The updated program data

**Returns:**
- `*models.Program` - The updated program
- `error` - Any error that occurred

#### `DeleteProgram(id uint) error`
Deletes a program.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `StartProgram(id uint) error`
Starts a program.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `StopProgram(id uint) error`
Stops a program.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `RestartProgram(id uint) error`
Restarts a program.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `ReloadProgram(id uint) error`
Reloads a program's configuration.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `EnableProgram(id uint) error`
Enables a program for autostart.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `DisableProgram(id uint) error`
Disables a program from autostart.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `ShellProgram(id uint) error`
Opens an interactive shell for a program.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `TailProgram(id uint) error`
Follows the logs of a program.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

#### `RunProgram(id uint) error`
Executes a program.

**Parameters:**
- `id` - The program ID

**Returns:**
- `error` - Any error that occurred

## Authentication

The HXE server uses JWT (JSON Web Tokens) for authentication. The client supports both automatic and manual token management.

### Default Credentials

For development, the server accepts these default credentials:
- **Username**: `admin`
- **Password**: `password`

### Authentication Flow

1. **Create authenticated client** with username and password
2. **Login** to get JWT token
3. **Use API** - token is automatically included in requests
4. **Refresh token** when needed
5. **Logout** to clear token

### Manual Token Management

You can also manage tokens manually:

```go
client := client.NewClient("http://localhost:8080")
client.SetToken("your-jwt-token-here")
```

### Token Refresh

Tokens expire after 24 hours. You can refresh them:

```go
err := client.RefreshToken()
if err != nil {
    // Handle refresh error
}
```

## Program Model

The `models.Program` struct contains the following fields:

```go
type Program struct {
    ID          uint          `json:"id"`
    AID         uint          `json:"aid"`
    GID         uint          `json:"gid"`
    Name        string        `json:"name"`
    Description string        `json:"desc"`
    Command     string        `json:"command"`
    Args        string        `json:"args"`
    Directory   string        `json:"cwd"`
    User        string        `json:"user"`
    Group       string        `json:"group"`
    Status      ProgramStatus `json:"status"`
    PID         int           `json:"pid"`
    ExitCode    int           `json:"exitCode"`
    StartTime   int64         `json:"startTime"`
    EndTime     int64         `json:"endTime"`
    Autostart   bool          `json:"autostart"`
    Enabled     bool          `json:"enabled"`
    Retries     int           `json:"retries"`
    MaxRetries  int           `json:"maxRetries"`
    Metrics     map[string]float64 `json:"metrics"`
    Actions     []Action      `json:"actions"`
}
```

## Error Handling

The client returns errors for various scenarios:

- Network errors
- HTTP errors (4xx, 5xx)
- Authentication errors (401 Unauthorized)
- API errors (server-side errors)
- JSON parsing errors

All errors include descriptive messages to help with debugging.

## Examples

See `examples/client/programs/main.go` for complete examples of how to use the client with and without authentication. 