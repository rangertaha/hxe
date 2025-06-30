# HXE Terminal User Interface (TUI)

A beautiful terminal-based user interface for managing HXE services, built with [Bubble Tea](https://github.com/charmbracelet/bubbletea).

## Features

- **Service Management**: View, start, stop, and restart services
- **Real-time Updates**: Automatic refresh every 5 seconds
- **Interactive Interface**: Navigate with keyboard and mouse
- **Service Details**: View service status, PID, uptime, and more
- **Modern UI**: Clean, colorful interface with proper styling

## Installation

```bash
# Build the TUI
go build -o bin/hxe-tui cmd/hxe-tui/main.go

# Or install globally
go install cmd/hxe-tui/main.go
```

## Usage

### Prerequisites

1. **NATS Server**: Make sure NATS is running (default: `nats://localhost:4222`)
2. **HXE Services**: Ensure HXE services are available through NATS

### Running the TUI

```bash
# Run the TUI
./bin/hxe-tui

# Or if installed globally
hxe-tui
```

## Key Bindings

| Key | Action |
|-----|--------|
| `↑` / `k` | Navigate up |
| `↓` / `j` | Navigate down |
| `Enter` | Select service |
| `s` | Start selected service |
| `x` | Stop selected service |
| `r` | Restart selected service |
| `R` | Refresh services list |
| `d` | View service details |
| `?` | Toggle help |
| `q` / `Ctrl+C` | Quit |

## Interface Layout

```
HXE Service Manager
Services: 5 | Selected: 2 | Timer: 00:05

┌────┬────────────────────┬────────────┬──────┬────────────┬──────┬────────┬──────────────────────────────┐
│ ID │ Name               │ Status     │ PID  │ Uptime     │ CPU% │ Memory%│ Description                  │
├────┼────────────────────┼────────────┼──────┼────────────┼──────┼────────┼──────────────────────────────┤
│ 1  │ web-server         │ Running    │ 1234 │ 2h 15m     │ 2.1  │ 45.2   │ Web server for API           │
│ 2  │ database           │ Running    │ 1235 │ 1h 30m     │ 1.8  │ 32.1   │ PostgreSQL database          │
│ 3  │ cache              │ Stopped    │ -    │ -          │ -    │ -      │ Redis cache server           │
│ 4  │ monitoring         │ Running    │ 1237 │ 45m        │ 0.5  │ 12.3   │ System monitoring agent      │
│ 5  │ backup             │ Failed     │ -    │ -          │ -    │ -      │ Automated backup service     │
└────┴────────────────────┴────────────┴──────┴────────────┴──────┴────────┴──────────────────────────────┘

↑/k: up  ↓/j: down  s: start  x: stop  r: restart  R: refresh  q: quit
```

## Service Status

The TUI displays the following service information:

- **ID**: Unique service identifier
- **Name**: Service name
- **Status**: Current service status (Ready, Loading, Starting, Running, Stopping, Stopped, Failed, Success)
- **PID**: Process ID (if running)
- **Uptime**: How long the service has been running
- **CPU%**: CPU usage percentage (if available)
- **Memory%**: Memory usage percentage (if available)
- **Description**: Service description

## Service Actions

### Start Service
- Select a service and press `s`
- The service will be started and status will update

### Stop Service
- Select a service and press `x`
- The service will be stopped and status will update

### Restart Service
- Select a service and press `r`
- The service will be restarted and status will update

### View Details
- Select a service and press `d`
- Detailed service information will be displayed

## Configuration

### NATS Connection

The TUI connects to NATS using the default URL (`nats://localhost:4222`). To use a different NATS server:

```go
// Modify internal/tui/main.go
nc, err := nats.Connect("nats://your-nats-server:4222")
```

### Refresh Interval

The TUI automatically refreshes the service list every 5 seconds. To change this:

```go
// Modify internal/tui/app/model.go
timer: timer.NewWithInterval(10*time.Second, time.Second), // 10 seconds
```

## Development

### Building from Source

```bash
# Clone the repository
git clone https://github.com/rangertaha/hxe.git
cd hxe

# Build the TUI
go build -o bin/hxe-tui cmd/hxe-tui/main.go
```

### Dependencies

The TUI uses the following key dependencies:

- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Bubbles](https://github.com/charmbracelet/bubbles) - UI components
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Styling
- [NATS Go](https://github.com/nats-io/nats.go) - NATS client

### Project Structure

```
internal/tui/
├── app/
│   └── model.go      # Main application model and logic
├── main.go           # TUI package entry point
└── README.md         # This file

cmd/hxe-tui/
└── main.go           # Command-line entry point
```

## Troubleshooting

### Connection Issues

If the TUI fails to connect to NATS:

1. Ensure NATS server is running
2. Check NATS server URL and port
3. Verify network connectivity
4. Check NATS server logs

### Service Not Found

If services don't appear in the list:

1. Verify HXE services are running
2. Check NATS subject patterns
3. Ensure services are properly registered
4. Check service logs

### UI Rendering Issues

If the UI doesn't render correctly:

1. Ensure terminal supports UTF-8
2. Check terminal size (minimum 80x24 recommended)
3. Verify terminal supports colors
4. Try a different terminal emulator

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the GNU General Public License v3.0 or later - see the [LICENSE](../../LICENSE) file for details.

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 