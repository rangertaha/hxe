# CLI Reference

The HXE command line interface provides a powerful way to manage programs and interact with the HXE server. This guide covers all available commands and options.

## Overview

```bash
hxe [OPTIONS] [COMMAND]
```

## Global Options

```bash
Options:
  -c, --config string    Configuration file path
  -d, --daemon          Run in daemon mode
  --debug               Enable debug logging
  -h, --help            Show help
  -v, --version         Show version
  --validate            Validate configuration file
  --check               Check configuration syntax
  --test-db             Test database connection
  --show-config         Show configuration summary
```

## Commands

### Server Commands

#### Daemon Mode

```bash
# Start HXE in daemon mode
hxe --daemon

# Start with custom configuration
hxe --config /etc/hxe/config.hcl --daemon

# Start with debug logging
hxe --daemon --debug
```

#### Configuration

```bash
# Validate configuration
hxe --config config.hcl --validate

# Check configuration syntax
hxe --config config.hcl --check

# Test database connection
hxe --config config.hcl --test-db

# Show configuration summary
hxe --config config.hcl --show-config
```

### Program Management Commands

#### List Programs

```bash
# List all programs
hxe list

# List with custom format
hxe list --format table
hxe list --format json
hxe list --format yaml

# List with filters
hxe list --status running
hxe list --status stopped
hxe list --enabled
hxe list --disabled
```

#### Program Information

```bash
# Show program details
hxe info <program-id>

# Show program status
hxe status <program-id>

# Show program logs
hxe logs <program-id>

# Follow program logs
hxe logs <program-id> --follow

# Show last N lines
hxe logs <program-id> --tail 100
```

#### Program Control

```bash
# Start a program
hxe start <program-id>

# Stop a program
hxe stop <program-id>

# Restart a program
hxe restart <program-id>

# Enable autostart
hxe enable <program-id>

# Disable autostart
hxe disable <program-id>

# Reload configuration
hxe reload <program-id>
```

#### Program Creation and Management

```bash
# Create a new program
hxe create --name "My Program" --command "/usr/bin/python3" --args "script.py"

# Create with all options
hxe create \
  --name "My Program" \
  --description "A test program" \
  --command "/usr/bin/python3" \
  --args "script.py" \
  --directory "/tmp" \
  --user "nobody" \
  --group "nobody" \
  --enabled \
  --auto-restart \
  --max-restarts 3

# Update a program
hxe update <program-id> --name "Updated Name" --description "Updated description"

# Delete a program
hxe delete <program-id>

# Delete multiple programs
hxe delete <id1> <id2> <id3>
```

#### Bulk Operations

```bash
# Start multiple programs
hxe start <id1> <id2> <id3>

# Stop multiple programs
hxe stop <id1> <id2> <id3>

# Restart multiple programs
hxe restart <id1> <id2> <id3>

# Enable multiple programs
hxe enable <id1> <id2> <id3>

# Disable multiple programs
hxe disable <id1> <id2> <id3>
```

### Execution Commands

#### Run Commands

```bash
# Run a simple command
hxe run echo "Hello, World!"

# Run with arguments
hxe run python script.py --arg1 value1 --arg2 value2

# Run in background
hxe run --background python long_running_script.py

# Run with custom user
hxe run --user nobody python script.py

# Run with custom directory
hxe run --directory /tmp python script.py

# Run with environment variables
hxe run --env VAR1=value1 --env VAR2=value2 python script.py
```

#### Interactive Commands

```bash
# Open shell for a program
hxe shell <program-id>

# Execute command in program context
hxe exec <program-id> ls -la

# Attach to program output
hxe attach <program-id>
```

### Monitoring Commands

#### Logs and Output

```bash
# Follow program logs
hxe logs <program-id> --follow

# Show last N lines
hxe logs <program-id> --tail 100

# Show logs with timestamps
hxe logs <program-id> --timestamps

# Show logs in JSON format
hxe logs <program-id> --format json

# Filter logs by level
hxe logs <program-id> --level error
hxe logs <program-id> --level warn
hxe logs <program-id> --level info
hxe logs <program-id> --level debug
```

#### Metrics and Statistics

```bash
# Show program metrics
hxe metrics <program-id>

# Show system metrics
hxe metrics --system

# Show all metrics
hxe metrics --all

# Export metrics
hxe metrics --export prometheus
hxe metrics --export json
```

### Authentication Commands

```bash
# Login to server
hxe login --username admin --password password

# Login with token
hxe login --token your-jwt-token

# Logout
hxe logout

# Show current authentication status
hxe auth status

# Refresh token
hxe auth refresh
```

## Command Options

### Global Options

| Option | Description | Default |
|--------|-------------|---------|
| `-c, --config` | Configuration file path | `~/.config/hxe/config.hcl` |
| `-d, --daemon` | Run in daemon mode | `false` |
| `--debug` | Enable debug logging | `false` |
| `-h, --help` | Show help | - |
| `-v, --version` | Show version | - |
| `--validate` | Validate configuration file | - |
| `--check` | Check configuration syntax | - |
| `--test-db` | Test database connection | - |
| `--show-config` | Show configuration summary | - |

### Program Options

| Option | Description | Default |
|--------|-------------|---------|
| `--name` | Program name | - |
| `--description` | Program description | - |
| `--command` | Program command | - |
| `--args` | Program arguments | - |
| `--directory` | Working directory | `/tmp` |
| `--user` | User to run as | `nobody` |
| `--group` | Group to run as | `nobody` |
| `--enabled` | Enable program | `false` |
| `--disabled` | Disable program | `false` |
| `--auto-restart` | Enable auto-restart | `false` |
| `--max-restarts` | Maximum restart attempts | `3` |

### Output Options

| Option | Description | Default |
|--------|-------------|---------|
| `--format` | Output format (table, json, yaml) | `table` |
| `--quiet` | Suppress output | `false` |
| `--verbose` | Verbose output | `false` |
| `--no-color` | Disable colored output | `false` |

## Examples

### Basic Usage

```bash
# Start the daemon
hxe --daemon

# List all programs
hxe list

# Create a simple program
hxe create --name "Test" --command "echo hello" --enabled

# Start the program
hxe start test

# Check status
hxe status test

# Stop the program
hxe stop test

# Delete the program
hxe delete test
```

### Advanced Usage

```bash
# Create a complex program
hxe create \
  --name "Web Server" \
  --description "Nginx web server" \
  --command "/usr/sbin/nginx" \
  --args "-g 'daemon off;'" \
  --directory "/var/www" \
  --user "www-data" \
  --group "www-data" \
  --enabled \
  --auto-restart \
  --max-restarts 5

# Monitor the program
hxe logs "Web Server" --follow

# Check metrics
hxe metrics "Web Server"

# Restart the program
hxe restart "Web Server"
```

### Bulk Operations

```bash
# Start multiple programs
hxe start web-server database cache

# Check status of all programs
hxe list --status running

# Stop all running programs
hxe stop $(hxe list --status running --format json | jq -r '.[].id')

# Enable all programs
hxe enable $(hxe list --disabled --format json | jq -r '.[].id')
```

### Scripting Examples

```bash
#!/bin/bash
# Start HXE daemon
hxe --daemon &

# Wait for daemon to start
sleep 2

# Create and start programs
hxe create --name "app1" --command "python app1.py" --enabled
hxe create --name "app2" --command "python app2.py" --enabled

hxe start app1 app2

# Monitor programs
hxe logs app1 --follow &
hxe logs app2 --follow &

# Wait for user input
read -p "Press Enter to stop programs..."

# Cleanup
hxe stop app1 app2
hxe delete app1 app2
```

## Environment Variables

HXE CLI supports environment variables for configuration:

```bash
# Server configuration
export HXE_SERVER_URL="http://localhost:8080"
export HXE_USERNAME="admin"
export HXE_PASSWORD="password"

# Output configuration
export HXE_FORMAT="json"
export HXE_NO_COLOR="true"
export HXE_QUIET="false"

# Authentication
export HXE_TOKEN="your-jwt-token"
```

## Exit Codes

HXE CLI uses the following exit codes:

| Code | Description |
|------|-------------|
| `0` | Success |
| `1` | General error |
| `2` | Configuration error |
| `3` | Authentication error |
| `4` | Program not found |
| `5` | Program already exists |
| `6` | Invalid command |
| `7` | Network error |
| `8` | Permission denied |

## Troubleshooting

### Common Issues

#### Connection Refused
```bash
# Check if daemon is running
ps aux | grep hxe

# Start daemon if not running
hxe --daemon
```

#### Authentication Failed
```bash
# Check credentials
hxe login --username admin --password password

# Check token
hxe auth status
```

#### Program Not Found
```bash
# List all programs
hxe list

# Check program ID
hxe info <program-id>
```

#### Permission Denied
```bash
# Check file permissions
ls -la ~/.config/hxe/

# Fix permissions
chmod 600 ~/.config/hxe/config.hcl
```

### Debug Mode

```bash
# Enable debug logging
hxe --debug list

# Show detailed error messages
hxe --debug --verbose start <program-id>
```

## Best Practices

1. **Use Configuration Files**: Store settings in configuration files
2. **Use Environment Variables**: For sensitive data like passwords
3. **Validate Configuration**: Always validate before deployment
4. **Monitor Programs**: Use logs and metrics for monitoring
5. **Use Bulk Operations**: For managing multiple programs
6. **Script Automation**: Use CLI in scripts for automation

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 