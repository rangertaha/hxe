# Installation Guide

This guide covers different ways to install HXE on various platforms.

## Prerequisites

- **Go 1.23 or higher** - [Download Go](https://golang.org/dl/)
- **Git** - For cloning the repository
- **Make** (optional) - For using build scripts

## Installation Methods

### 1. From Source (Recommended)

Clone the repository and build from source:

```bash
# Clone the repository
git clone https://github.com/rangertaha/hxe.git
cd hxe

# Build the project
go build -o bin/hxe ./cmd/hxe

# Install globally (optional)
sudo cp bin/hxe /usr/local/bin/
```

### 2. Using GitHub Packages

HXE is available as a Go module from GitHub Packages:

```bash
# Configure Go to use GitHub Packages
export GOPRIVATE=github.com/rangertaha/hxe
export GOPROXY=https://proxy.golang.org,direct

# Install HXE
go install github.com/rangertaha/hxe/cmd/hxe@latest

# Or add to your project
go get github.com/rangertaha/hxe/pkg/client
```

### 3. Using Make

If you have Make installed, you can use the provided build scripts:

```bash
# Clone the repository
git clone https://github.com/rangertaha/hxe.git
cd hxe

# Build backend server
make server

# Build desktop application
make desktop

# Clean build artifacts
make clean

# Install globally
make install
```

### 4. Cross-platform Builds

Build HXE for different platforms:

```bash
# Linux AMD64
GOOS=linux GOARCH=amd64 go build -o bin/hxe-linux-amd64 ./cmd/hxe

# macOS AMD64
GOOS=darwin GOARCH=amd64 go build -o bin/hxe-darwin-amd64 ./cmd/hxe

# Windows AMD64
GOOS=windows GOARCH=amd64 go build -o bin/hxe-windows-amd64.exe ./cmd/hxe

# Linux ARM64
GOOS=linux GOARCH=arm64 go build -o bin/hxe-linux-arm64 ./cmd/hxe
```

## Platform-Specific Instructions

### Linux

#### Ubuntu/Debian

```bash
# Install dependencies
sudo apt update
sudo apt install git golang-go make

# Clone and build
git clone https://github.com/rangertaha/hxe.git
cd hxe
go build -o bin/hxe ./cmd/hxe
sudo cp bin/hxe /usr/local/bin/
```

#### CentOS/RHEL/Fedora

```bash
# Install dependencies
sudo dnf install git golang make

# Clone and build
git clone https://github.com/rangertaha/hxe.git
cd hxe
go build -o bin/hxe ./cmd/hxe
sudo cp bin/hxe /usr/local/bin/
```

### macOS

#### Using Homebrew

```bash
# Install Go
brew install go

# Clone and build
git clone https://github.com/rangertaha/hxe.git
cd hxe
go build -o bin/hxe ./cmd/hxe
sudo cp bin/hxe /usr/local/bin/
```

#### Using MacPorts

```bash
# Install Go
sudo port install go

# Clone and build
git clone https://github.com/rangertaha/hxe.git
cd hxe
go build -o bin/hxe ./cmd/hxe
sudo cp bin/hxe /usr/local/bin/
```

### Windows

#### Using Chocolatey

```bash
# Install Go
choco install golang

# Clone and build
git clone https://github.com/rangertaha/hxe.git
cd hxe
go build -o bin/hxe.exe ./cmd/hxe
```

#### Using Scoop

```bash
# Install Go
scoop install go

# Clone and build
git clone https://github.com/rangertaha/hxe.git
cd hxe
go build -o bin/hxe.exe ./cmd/hxe
```

## Verification

After installation, verify that HXE is working correctly:

```bash
# Check version
hxe --version

# Show help
hxe --help

# Test basic functionality
hxe run echo "Hello, HXE!"
```

## Configuration

After installation, create your first configuration:

```bash
# Create configuration directory
mkdir -p ~/.config/hxe

# Generate default configuration
hxe --config ~/.config/hxe/config.hcl --daemon
```

## Next Steps

1. [Configure HXE](configuration.md)
2. [Start the daemon](cli.md#daemon-mode)
3. [Use the client library](client.md)
4. [Explore the CLI](cli.md)

## Troubleshooting

### Common Issues

#### Go not found
```bash
# Add Go to PATH
export PATH=$PATH:/usr/local/go/bin
# Add to ~/.bashrc or ~/.zshrc for persistence
```

#### Permission denied
```bash
# Make sure the binary is executable
chmod +x bin/hxe
```

#### Build errors
```bash
# Clean and rebuild
go clean
go mod tidy
go build -o bin/hxe ./cmd/hxe
```

### Getting Help

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 