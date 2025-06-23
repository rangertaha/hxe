# HXE Client

A Go client library for interacting with the HXE API server with JWT authentication support.

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
    // Create a client
    hxeClient := client.NewClient("http://localhost:8080", "admin", "password")
    
    // Get the program client
    programClient := hxeClient.Program
    
    // List programs
    programs, err := programClient.List()
    if err != nil {
        log.Fatal(err)
    }
    
    // Print programs in a table format
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
    // Create an authenticated client
    hxeClient := client.NewClient("http://localhost:8080", "admin", "password")
    
    // Login to get JWT token
    _, err := hxeClient.Login()
    if err != nil {
        log.Fatal("Login failed:", err)
    }
    
    // Get the program client
    programClient := hxeClient.Program
    
    // List programs
    programs, err := programClient.List()
    if err != nil {
        log.Fatal(err)
    }
    
    // Print programs
    programClient.Print(programs)
    
    // Logout
    hxeClient.Logout()
}
```

## API Reference

### Creating Clients

#### `NewClient(baseURL, username, password string) *Client`
Creates a new authenticated client with the given credentials.

```go
client := client.NewClient("http://localhost:8080", "admin", "password")
```

### Authentication Methods

#### `Login() (*Client, error)`
Authenticates with the server and retrieves a JWT token.

```go
_, err := client.Login()
if err != nil {
    // Handle login error
}
```

#### `Logout()`
Clears the current JWT token.

```go
client.Logout()
```

#### `RefreshToken() error`
Refreshes the JWT token.

```go
err := client.RefreshToken()
if err != nil {
    // Handle refresh error
}
```

#### `SetToken(token string)`
Manually sets a JWT token.

```go
client.SetToken("your-jwt-token-here")
```

#### `GetToken() string`
Returns the current JWT token.

```go
token := client.GetToken()
```

### Program Operations

#### `List() ([]*models.Program, error)`
Retrieves all programs.

```go
programs, err := programClient.List()
```

#### `Get(id string) (*models.Program, error)`
Retrieves a specific program by ID.

```go
program, err := programClient.Get("123")
```

#### `Status(id string) (*models.Program, error)`
Gets the status of a program (alias for Get).

```go
program, err := programClient.Status("123")
```

#### `MultiStatus(ids ...string) ([]*models.Program, error)`
Gets the status of multiple programs.

```go
programs, err := programClient.MultiStatus("123", "456", "789")
```

#### `Create(program *models.Program) (*models.Program, error)`
Creates a new program.

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
}

created, err := programClient.Create(newProgram)
```

#### `Update(id string, program *models.Program) (*models.Program, error)`
Updates an existing program.

```go
program.Name = "Updated Program Name"
updated, err := programClient.Update("123", program)
```

#### `Delete(id string) (*models.Program, error)`
Deletes a program.

```go
deleted, err := programClient.Delete("123")
```

#### `MultiDelete(ids ...string) ([]*models.Program, error)`
Deletes multiple programs.

```go
deleted, err := programClient.MultiDelete("123", "456", "789")
```

### Runtime Operations

#### `Start(id string) (*models.Program, error)`
Starts a program.

```go
program, err := programClient.Start("123")
```

#### `MultiStart(ids ...string) ([]*models.Program, error)`
Starts multiple programs.

```go
programs, err := programClient.MultiStart("123", "456", "789")
```

#### `Stop(id string) (*models.Program, error)`
Stops a program.

```go
program, err := programClient.Stop("123")
```

#### `MultiStop(ids ...string) ([]*models.Program, error)`
Stops multiple programs.

```go
programs, err := programClient.MultiStop("123", "456", "789")
```

#### `Restart(id string) (*models.Program, error)`
Restarts a program.

```go
program, err := programClient.Restart("123")
```

#### `MultiRestart(ids ...string) ([]*models.Program, error)`
Restarts multiple programs.

```go
programs, err := programClient.MultiRestart("123", "456", "789")
```

#### `Enable(id string) (*models.Program, error)`
Enables a program for autostart.

```go
program, err := programClient.Enable("123")
```

#### `MultiEnable(ids ...string) ([]*models.Program, error)`
Enables multiple programs for autostart.

```go
programs, err := programClient.MultiEnable("123", "456", "789")
```

#### `Disable(id string) (*models.Program, error)`
Disables a program from autostart.

```go
program, err := programClient.Disable("123")
```

#### `MultiDisable(ids ...string) ([]*models.Program, error)`
Disables multiple programs from autostart.

```go
programs, err := programClient.MultiDisable("123", "456", "789")
```

#### `Reload(id string) (*models.Program, error)`
Reloads a program's configuration.

```go
program, err := programClient.Reload("123")
```

#### `MultiReload(ids ...string) ([]*models.Program, error)`
Reloads multiple programs' configurations.

```go
programs, err := programClient.MultiReload("123", "456", "789")
```

### Advanced Operations

#### `Run(command string) (*models.Program, error)`
Creates and starts a program with the given command.

```go
program, err := programClient.Run("python3 /path/to/script.py")
```

#### `Shell(id string) (*models.Program, error)`
Opens an interactive shell for a program.

```go
program, err := programClient.Shell("123")
```

#### `Tail(id string) (*models.Program, error)`
Follows the logs of a program.

```go
program, err := programClient.Tail("123")
```

### Display Methods

#### `Print(programs []*models.Program)`
Intelligently displays programs based on count:
- 0 programs: Shows "No programs found"
- 1 program: Shows detailed view
- Multiple programs: Shows list view

```go
programs, err := programClient.List()
if err != nil {
    log.Fatal(err)
}
programClient.Print(programs)
```

#### `PrintDetail(program *models.Program)`
Displays a single program in detailed table format.

```go
program, err := programClient.Get("123")
if err != nil {
    log.Fatal(err)
}
programClient.PrintDetail(program)
```

#### `PrintList(programs []*models.Program)`
Displays multiple programs in a table format.

```go
programs, err := programClient.List()
if err != nil {
    log.Fatal(err)
}
programClient.PrintList(programs)
```

## Authentication

The client supports JWT authentication with the following features:

- **Automatic token management**: Tokens are automatically included in API requests
- **Token refresh**: Tokens can be refreshed before expiration
- **Manual token management**: Tokens can be set manually if needed
- **Default credentials**: Username `admin`, password `password`

### Authentication Flow

1. Create a client with credentials
2. Call `Login()` to authenticate and get a JWT token
3. The token is automatically included in subsequent API requests
4. Use `RefreshToken()` to extend the token's validity
5. Call `Logout()` to clear the token

## Program Model

The `models.Program` struct contains the following fields:

```go
type Program struct {
    ID          uint          `json:"id"`
    AID         uint          `json:"aid"`
    GID         uint          `json:"gid"`
    SID         uint          `json:"sid"`
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

### Complete Example

```go
package main

import (
    "log"
    "github.com/rangertaha/hxe/pkg/client"
)

func main() {
    // Create authenticated client
    hxeClient := client.NewClient("http://localhost:8080", "admin", "password")
    
    // Login
    _, err := hxeClient.Login()
    if err != nil {
        log.Fatal("Login failed:", err)
    }
    
    programClient := hxeClient.Program
    
    // List all programs
    programs, err := programClient.List()
    if err != nil {
        log.Fatal("Failed to list programs:", err)
    }
    
    // Display programs
    programClient.Print(programs)
    
    // Create a new program
    newProgram := &models.Program{
        Name:        "Test Program",
        Description: "A test program",
        Command:     "/usr/bin/python3",
        Args:        "test.py",
        Directory:   "/tmp",
        User:        "nobody",
        Group:       "nobody",
        Enabled:     true,
    }
    
    created, err := programClient.Create(newProgram)
    if err != nil {
        log.Fatal("Failed to create program:", err)
    }
    
    // Start the program
    started, err := programClient.Start(created.ID)
    if err != nil {
        log.Fatal("Failed to start program:", err)
    }
    
    // Get program status
    status, err := programClient.Status(created.ID)
    if err != nil {
        log.Fatal("Failed to get status:", err)
    }
    
    // Display program details
    programClient.PrintDetail(status)
    
    // Stop the program
    stopped, err := programClient.Stop(created.ID)
    if err != nil {
        log.Fatal("Failed to stop program:", err)
    }
    
    // Delete the program
    deleted, err := programClient.Delete(created.ID)
    if err != nil {
        log.Fatal("Failed to delete program:", err)
    }
    
    // Logout
    hxeClient.Logout()
} 