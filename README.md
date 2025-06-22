# Hxe

HXE - Host eXecution Engine

Hxe is a host based process execution engine with command line, desktop, and web user interfaces. 



[![Go Version](https://img.shields.io/badge/Go-1.23+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)


## Features

- 🚀 **Service Management**: Start, stop, restart, and monitor services
- 🖥️ **Desktop Interface**: Modern web-based UI for service management
- 💻 **Command Line Interface**: Full-featured CLI for automation
- 🔧 **Configuration**: HCL-based configuration files
- 📊 **Monitoring**: Real-time service status and logs
- 🔄 **Auto-restart**: Automatic service recovery on failure
- 🌐 **Web API**: RESTful API for integration
- 📝 **Logging**: Comprehensive logging with multiple levels
- 🔐 **Authentication**: Built-in security features

## Screenshots

### Desktop Interface
![Desktop Interface](docs/img/desktop.png)

### Terminal Interface
![Terminal Interface](docs/img/terminal.png)

## Installation

### Prerequisites

- Go 1.23 or higher
- Git

### From Source

```bash
# Clone the repository
git clone https://github.com/rangertaha/hxe.git
cd hxe

# Build the project
go build -o bin/hxe ./cmd/hxe

# Install globally (optional)
sudo cp bin/hxe /usr/local/bin/
```

### Using Make

```bash
# Build backend server
make server

# Build desktop application
make desktop

# Clean build artifacts
make clean
```

## Quick Start

### 1. Initialize Configuration

Hxe will automatically create a default configuration file on first run:

```bash
hxe --help
```

This creates a default configuration at `~/.config/hxe/config.hcl`.

### 2. Start a Service

```bash
# Run a simple command
hxe echo "Hello, World!"

# Run in daemon mode
hxe --daemon

# Execute with custom config
hxe --config /path/to/config.hcl ls -la
```

### 3. Desktop Interface

Start the desktop interface to manage services through a web browser:

```bash
hxe --daemon
```

Then open your browser to `http://localhost:8080`.

## Configuration

Hxe uses HCL (HashiCorp Configuration Language) for configuration files. The default configuration is created at `~/.config/hxe/config.hcl`.

### Basic Configuration

```hcl
// hxe Configuration
debug = false
version = "0.0.0"

api {
  addr = "0.0.0.0"
  port = 8080
  username = "hxe"
  password = "hxe"
}

broker { 
  name = "hxe"
  addr = "0.0.0.0"
  port = 8080
}
```

### Service Configuration

Create service configuration files (e.g., `service.hcl`):

```hcl
// Service Configuration
service "my-service" {
  command = "python"
  args = ["app.py"]
}
```

## Usage

### Command Line Options

```bash
hxe [OPTIONS] [COMMAND]

Options:
  -d, --daemon    Run in daemon mode
  -c, --config    Configuration file path
  --debug         Enable debug logging
  -h, --help      Show help
  -v, --version   Show version

Commands:
  run             Run a command or service
  list            List all services
  start           Start a service
  stop            Stop a service
  restart         Restart a service
  status          Show service status
  tail            Tail service logs
  reload          Reload configuration
  enable          Enable a service
  disable         Disable a service
  shell           Open shell for a service
```

### Subcommands

#### `run` - Execute Commands
```bash
# Run a simple command
hxe run echo "Hello, World!"

# Run a Python script
hxe run python script.py

# Run with arguments
hxe run node server.js --port 3000
```

#### `list` - List Services
```bash
# List all configured services
hxe list

# Alternative aliases
hxe ls
hxe l
```

#### `start` - Start Services
```bash
# Start a specific service
hxe start my-service

# Alternative alias
hxe s my-service
```

#### `stop` - Stop Services
```bash
# Stop a specific service
hxe stop my-service

# Alternative alias
hxe st my-service
```

#### `restart` - Restart Services
```bash
# Restart a specific service
hxe restart my-service

# Alternative alias
hxe rs my-service
```

#### `status` - Service Status
```bash
# Show status of all services
hxe status

# Show status of a specific service
hxe status my-service

# Alternative alias
hxe st my-service
```

#### `tail` - Follow Logs
```bash
# Tail logs with default settings (50 lines, follow enabled)
hxe tail my-service

# Tail with custom number of lines
hxe tail --lines 100 my-service

# Tail without following (show last N lines only)
hxe tail --follow=false my-service

# Alternative aliases
hxe t my-service
hxe tail -n 100 my-service
```

#### `reload` - Reload Configuration
```bash
# Reload configuration without restarting services
hxe reload

# Alternative alias
hxe rl
```

#### `enable` - Enable Services
```bash
# Enable a service to start automatically
hxe enable my-service

# Alternative alias
hxe e my-service
```

#### `disable` - Disable Services
```bash
# Disable a service from starting automatically
hxe disable my-service

# Alternative alias
hxe d my-service
```

#### `shell` - Interactive Shell
```bash
# Open an interactive shell for a service
hxe shell my-service

# Alternative alias
hxe sh my-service
```

### Examples

#### Basic Service Execution

```bash
# Run a simple command
hxe ls -la

# Run a Python script
hxe python script.py

# Run with arguments
hxe node server.js --port 3000
```

#### Daemon Mode

```bash
# Start hxe in daemon mode
hxe --daemon

# Start with custom configuration
hxe --config /etc/hxe/config.hcl --daemon
```

#### Debug Mode

```bash
# Enable debug logging
hxe --debug ls -la

# Debug with daemon mode
hxe --debug --daemon
```

## API Reference

### REST Endpoints

When running in daemon mode, Hxe exposes a REST API:

- `GET /api/services` - List all services
- `GET /api/services/{id}` - Get service details
- `POST /api/services` - Create a new service
- `PUT /api/services/{id}` - Update a service
- `DELETE /api/services/{id}` - Delete a service
- `POST /api/services/{id}/start` - Start a service
- `POST /api/services/{id}/stop` - Stop a service
- `POST /api/services/{id}/restart` - Restart a service

### Authentication

The API supports basic authentication with the credentials specified in your configuration file.

## Development

### Project Structure

```
hxe/
├── cmd/hxe/          # Main application entry point
├── internal/         # Internal packages
│   ├── agent/        # Service agent
│   ├── config/       # Configuration management
│   └── log/          # Logging utilities
├── pkg/              # Public packages
├── desktop/          # Desktop application
├── examples/         # Example configurations
└── docs/             # Documentation
```

### Building

```bash
# Build for current platform
go build -o bin/hxe ./cmd/hxe

# Build for specific platform
GOOS=linux GOARCH=amd64 go build -o bin/hxe-linux-amd64 ./cmd/hxe
GOOS=darwin GOARCH=amd64 go build -o bin/hxe-darwin-amd64 ./cmd/hxe
GOOS=windows GOARCH=amd64 go build -o bin/hxe-windows-amd64.exe ./cmd/hxe
```

### Testing

```bash
# Run all tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Run specific test
go test ./internal/agent
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Inspired by [Supervisord](http://supervisord.org/), [PM2](https://pm2.keymetrics.io/), and [Systemd](https://systemd.io/)
- Built with [Go](https://golang.org/), [NATS](https://nats.io/), and [React](https://reactjs.org/)

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki)



