# Development Guide

This guide covers how to set up a development environment for HXE, contribute to the project, and understand the codebase structure.

## Prerequisites

- **Go 1.23 or higher** - [Download Go](https://golang.org/dl/)
- **Git** - For version control
- **Make** - For build scripts (optional)
- **Docker** - For containerized development (optional)

## Development Setup

### 1. Clone the Repository

```bash
git clone https://github.com/rangertaha/hxe.git
cd hxe
```

### 2. Install Dependencies

```bash
# Download Go modules
go mod download

# Install development tools
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/go-delve/delve/cmd/dlv@latest
```

### 3. Build the Project

```bash
# Build for current platform
go build -o bin/hxe ./cmd/hxe

# Build with debug information
go build -gcflags="all=-N -l" -o bin/hxe ./cmd/hxe

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o bin/hxe-linux-amd64 ./cmd/hxe
GOOS=darwin GOARCH=amd64 go build -o bin/hxe-darwin-amd64 ./cmd/hxe
GOOS=windows GOARCH=amd64 go build -o bin/hxe-windows-amd64.exe ./cmd/hxe
```

### 4. Run Tests

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run tests with verbose output
go test -v ./...

# Run specific test
go test ./internal/api

# Run benchmarks
go test -bench=. ./...
```

## Project Structure

```
hxe/
├── cmd/hxe/                    # Main application entry point
│   └── main.go                # Application main function
├── internal/                   # Internal packages (not importable)
│   ├── agent/                 # Service agent implementation
│   │   ├── agent.go           # Main agent logic
│   │   ├── api.go             # Agent API handlers
│   │   ├── broker.go          # Message broker integration
│   │   └── supervisor.go      # Process supervision
│   ├── api/                   # API server implementation
│   │   ├── server.go          # HTTP server setup
│   │   ├── handlers/          # HTTP request handlers
│   │   │   └── program.go     # Program management handlers
│   │   └── services/          # Business logic services
│   │       └── program.go     # Program service layer
│   ├── config/                # Configuration management
│   │   ├── config.go          # Configuration structures
│   │   ├── broker.go          # Broker configuration
│   │   ├── seed.go            # Default configuration
│   │   └── hclparse.go        # HCL configuration parser
│   ├── models/                # Data models
│   │   ├── program.go         # Program model
│   │   ├── init.go            # Model initialization
│   │   └── metrics.go         # Metrics models
│   ├── log/                   # Logging utilities
│   │   └── logger.go          # Logger implementation
│   ├── interfaces/            # Interface definitions
│   │   ├── agent.go           # Agent interfaces
│   │   ├── bots.go            # Bot interfaces
│   │   ├── data.go            # Data interfaces
│   │   ├── metric.go          # Metric interfaces
│   │   ├── node.go            # Node interfaces
│   │   └── runtime.go         # Runtime interfaces
│   └── interfaces.go          # Main interface definitions
├── pkg/client/                # Go client library (importable)
│   ├── client.go              # Main client implementation
│   ├── programs.go            # Program client methods
│   └── README.md              # Client documentation
├── examples/                  # Example code and configurations
│   └── client/                # Client usage examples
│       └── programs/          # Program management examples
│           └── main.go        # Main example
├── docs/                      # Documentation
│   ├── index.md               # Documentation index
│   ├── installation.md        # Installation guide
│   ├── configuration.md       # Configuration guide
│   ├── api.md                 # API reference
│   ├── client.md              # Client library guide
│   ├── cli.md                 # CLI reference
│   ├── development.md         # This file
│   └── packages.md            # GitHub Packages guide
├── desktop/                   # Desktop application (future)
├── go.mod                     # Go module definition
├── go.sum                     # Go module checksums
├── LICENSE                    # GPLv3+ license
├── README.md                  # Project README
└── Makefile                   # Build scripts
```

## Development Workflow

### 1. Create a Feature Branch

```bash
# Create and switch to feature branch
git checkout -b feature/your-feature-name

# Or create branch from main
git checkout main
git pull origin main
git checkout -b feature/your-feature-name
```

### 2. Make Changes

```bash
# Edit files
vim internal/api/handlers/program.go

# Add new files
touch internal/api/handlers/new_handler.go
```

### 3. Run Tests and Linting

```bash
# Run tests
go test ./...

# Run linter
golangci-lint run

# Run linter with specific rules
golangci-lint run --enable=goimports,unused,deadcode
```

### 4. Commit Changes

```bash
# Stage changes
git add .

# Commit with descriptive message
git commit -m "feat: add new program handler

- Add support for program metrics
- Implement real-time status updates
- Add comprehensive error handling

Closes #123"
```

### 5. Push and Create Pull Request

```bash
# Push to remote
git push origin feature/your-feature-name

# Create pull request on GitHub
```

## Code Style and Standards

### Go Code Style

Follow the [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments):

```go
// Good: Clear, descriptive function name
func CreateProgram(ctx context.Context, program *models.Program) (*models.Program, error) {
    // Implementation
}

// Good: Clear variable names
programName := "web-server"
maxRetries := 3

// Good: Proper error handling
if err != nil {
    return nil, fmt.Errorf("failed to create program: %w", err)
}
```

### Package Organization

- **`internal/`**: Private packages, not importable by external code
- **`pkg/`**: Public packages, importable by external code
- **`cmd/`**: Application entry points

### Error Handling

```go
// Use wrapped errors
if err != nil {
    return fmt.Errorf("failed to start program %s: %w", programID, err)
}

// Use custom error types for specific errors
var (
    ErrProgramNotFound = errors.New("program not found")
    ErrProgramRunning  = errors.New("program already running")
)
```

### Logging

```go
// Use structured logging
logger.Info("program started",
    "program_id", program.ID,
    "pid", program.PID,
    "command", program.Command,
)

// Use appropriate log levels
logger.Debug("processing request", "request_id", reqID)
logger.Info("user logged in", "username", username)
logger.Warn("program restarting", "program_id", programID)
logger.Error("failed to start program", "error", err)
```

## Testing

### Unit Tests

```go
// internal/api/services/program_test.go
package services

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestCreateProgram(t *testing.T) {
    // Arrange
    service := NewProgramService()
    program := &models.Program{
        Name:    "test-program",
        Command: "/usr/bin/echo",
        Args:    "hello",
    }

    // Act
    result, err := service.CreateProgram(program)

    // Assert
    assert.NoError(t, err)
    assert.NotNil(t, result)
    assert.Equal(t, "test-program", result.Name)
}
```

### Integration Tests

```go
// tests/integration/program_test.go
package integration

import (
    "testing"
    "github.com/rangertaha/hxe/pkg/client"
)

func TestProgramLifecycle(t *testing.T) {
    // Setup
    client := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")
    _, err := client.Login()
    assert.NoError(t, err)

    // Test program creation
    program := &models.Program{
        Name:    "test-program",
        Command: "/usr/bin/echo",
        Args:    "hello",
    }
    
    created, err := client.Program.CreateProgram(program)
    assert.NoError(t, err)
    assert.NotNil(t, created)

    // Test program start
    _, err = client.Program.StartProgram(created.ID)
    assert.NoError(t, err)

    // Test program stop
    _, err = client.Program.StopProgram(created.ID)
    assert.NoError(t, err)

    // Cleanup
    err = client.Program.DeleteProgram(created.ID)
    assert.NoError(t, err)
}
```

### Benchmark Tests

```go
// internal/api/services/program_bench_test.go
package services

import (
    "testing"
)

func BenchmarkListPrograms(b *testing.B) {
    service := NewProgramService()
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _, err := service.ListPrograms()
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

## Debugging

### Using Delve

```bash
# Debug the application
dlv debug ./cmd/hxe

# Debug with arguments
dlv debug ./cmd/hxe -- --daemon

# Attach to running process
dlv attach <pid>
```

### Debug Configuration

```hcl
# config/debug.hcl
debug = true
log_level = "debug"

api {
  addr = "localhost"
  port = 8080
  username = "admin"
  password = "password"
}

database {
  type = "sqlite"
  path = "./debug.db"
  migrate = true
}
```

### Debug Logging

```go
// Enable debug logging
logger.SetLevel(log.DebugLevel)

// Add debug statements
logger.Debug("processing request",
    "method", r.Method,
    "path", r.URL.Path,
    "headers", r.Header,
)
```

## Performance Profiling

### CPU Profiling

```bash
# Run with CPU profiling
go run -cpuprofile=cpu.prof ./cmd/hxe --daemon

# Analyze profile
go tool pprof cpu.prof
```

### Memory Profiling

```bash
# Run with memory profiling
go run -memprofile=mem.prof ./cmd/hxe --daemon

# Analyze profile
go tool pprof mem.prof
```

### HTTP Profiling

```go
import _ "net/http/pprof"

// Add to main function
go func() {
    log.Println(http.ListenAndServe("localhost:6060", nil))
}()
```

## Documentation

### Code Documentation

```go
// ProgramService handles program management operations
type ProgramService struct {
    db     Database
    logger Logger
}

// CreateProgram creates a new program with the given configuration.
// It validates the program configuration and stores it in the database.
// Returns the created program or an error if creation fails.
func (s *ProgramService) CreateProgram(program *models.Program) (*models.Program, error) {
    // Implementation
}
```

### API Documentation

Update API documentation in `docs/api.md` when adding new endpoints.

### README Updates

Update relevant documentation when adding new features:

- `README.md` - Main project documentation
- `docs/` - Detailed documentation
- `pkg/client/README.md` - Client library documentation

## Release Process

### 1. Version Bumping

```bash
# Update version in code
# Update go.mod if needed
# Update documentation
```

### 2. Tagging

```bash
# Create annotated tag
git tag -a v0.1.0 -m "Release v0.1.0"

# Push tag
git push origin v0.1.0
```

### 3. Building Releases

```bash
# Build for multiple platforms
make release

# Or manually
GOOS=linux GOARCH=amd64 go build -o bin/hxe-linux-amd64 ./cmd/hxe
GOOS=darwin GOARCH=amd64 go build -o bin/hxe-darwin-amd64 ./cmd/hxe
GOOS=windows GOARCH=amd64 go build -o bin/hxe-windows-amd64.exe ./cmd/hxe
```

## Contributing Guidelines

### Pull Request Process

1. **Fork the repository**
2. **Create a feature branch**
3. **Make your changes**
4. **Add tests for new functionality**
5. **Ensure all tests pass**
6. **Update documentation**
7. **Submit a pull request**

### Commit Message Format

Use [Conventional Commits](https://www.conventionalcommits.org/):

```
<type>[optional scope]: <description>

[optional body]

[optional footer(s)]
```

Types:
- `feat`: New feature
- `fix`: Bug fix
- `docs`: Documentation changes
- `style`: Code style changes
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `chore`: Maintenance tasks

### Code Review Checklist

- [ ] Code follows Go style guidelines
- [ ] Tests are included and passing
- [ ] Documentation is updated
- [ ] No breaking changes (or documented)
- [ ] Error handling is appropriate
- [ ] Logging is appropriate
- [ ] Security considerations addressed

## Troubleshooting

### Common Issues

#### Build Errors
```bash
# Clean and rebuild
go clean
go mod tidy
go build ./cmd/hxe
```

#### Test Failures
```bash
# Run tests with verbose output
go test -v ./...

# Run specific test
go test -v ./internal/api -run TestCreateProgram
```

#### Linting Errors
```bash
# Fix imports
goimports -w .

# Fix formatting
gofmt -w .

# Run linter with auto-fix
golangci-lint run --fix
```

#### Database Issues
```bash
# Reset database
rm -f *.db
go run ./cmd/hxe --daemon
```

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 