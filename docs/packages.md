# GitHub Packages

HXE is available as a Go module from GitHub Packages, making it easy to install and use in your projects.

## Overview

GitHub Packages provides a secure, reliable way to distribute Go modules. HXE packages are published to GitHub Packages and can be installed directly using `go install` or `go get`.

## Installation

### Prerequisites

- Go 1.23 or higher
- Git configured with your GitHub credentials

### Configure Go for GitHub Packages

```bash
# Set GOPRIVATE to include HXE repository
export GOPRIVATE=github.com/rangertaha/hxe

# Configure GOPROXY to use direct mode for private modules
export GOPROXY=https://proxy.golang.org,direct

# Add to your shell profile for persistence
echo 'export GOPRIVATE=github.com/rangertaha/hxe' >> ~/.bashrc
echo 'export GOPROXY=https://proxy.golang.org,direct' >> ~/.bashrc
```

### Install HXE Binary

```bash
# Install the latest version
go install github.com/rangertaha/hxe/cmd/hxe@latest

# Install a specific version
go install github.com/rangertaha/hxe/cmd/hxe@v0.1.0

# Install from main branch
go install github.com/rangertaha/hxe/cmd/hxe@main
```

### Add to Your Project

```bash
# Add HXE client library to your project
go get github.com/rangertaha/hxe/pkg/client

# Add specific version
go get github.com/rangertaha/hxe/pkg/client@v0.1.0
```

## Authentication

### Personal Access Token

For private repositories or to avoid rate limits, use a GitHub Personal Access Token:

```bash
# Configure Git to use token
git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# Or set environment variable
export GITHUB_TOKEN=your_github_token_here
```

### SSH Authentication

Alternatively, use SSH authentication:

```bash
# Configure Git to use SSH
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

## Usage Examples

### Using the Binary

```bash
# Check version
hxe --version

# Start daemon
hxe --daemon

# List programs
hxe list

# Run a command
hxe run "echo Hello, World!"
```

### Using the Client Library

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
    
    // Display in table format
    programClient.PrintList(programs)
    
    // Logout when done
    hxeClient.Logout()
}
```

## Version Management

### Available Versions

```bash
# List available versions
go list -m -versions github.com/rangertaha/hxe

# Install specific version
go install github.com/rangertaha/hxe/cmd/hxe@v0.1.0
```

### Updating

```bash
# Update to latest version
go install github.com/rangertaha/hxe/cmd/hxe@latest

# Update client library
go get -u github.com/rangertaha/hxe/pkg/client
```

## Troubleshooting

### Common Issues

#### Module not found
```bash
# Ensure GOPRIVATE is set
export GOPRIVATE=github.com/rangertaha/hxe

# Clear module cache
go clean -modcache

# Try again
go install github.com/rangertaha/hxe/cmd/hxe@latest
```

#### Authentication errors
```bash
# Check Git configuration
git config --list | grep url

# Verify GitHub token
curl -H "Authorization: token $GITHUB_TOKEN" https://api.github.com/user
```

#### Rate limiting
```bash
# Use GitHub token to avoid rate limits
export GITHUB_TOKEN=your_github_token_here
go install github.com/rangertaha/hxe/cmd/hxe@latest
```

### Environment Variables

```bash
# Required for GitHub Packages
export GOPRIVATE=github.com/rangertaha/hxe
export GOPROXY=https://proxy.golang.org,direct

# Optional for authentication
export GITHUB_TOKEN=your_github_token_here
```

## Package Structure

```
github.com/rangertaha/hxe/
├── cmd/hxe/              # Main binary
├── pkg/client/           # Go client library
└── internal/             # Internal packages (not importable)
```

## Security

- All packages are signed and verified
- Source code is publicly available
- Licensed under GPLv3+
- No telemetry or data collection

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 