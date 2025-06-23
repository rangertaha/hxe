# Configuration Guide

HXE uses HCL (HashiCorp Configuration Language) for configuration files. This guide covers all configuration options and provides examples for different use cases.

## Configuration File Location

By default, HXE looks for configuration files in the following locations:

1. `~/.config/hxe/config.hcl` (Linux/macOS)
2. `%APPDATA%\hxe\config.hcl` (Windows)
3. Custom path specified with `--config` flag

## Basic Configuration

### Minimal Configuration

```hcl
// HXE Configuration
debug = false
version = "0.1.0"

api {
  addr = "0.0.0.0"
  port = 8080
  username = "admin"
  password = "password"
}

database {
  type = "sqlite"
  migrate = true
}

broker { 
  name = "hxe"
  addr = "0.0.0.0"
  port = 4222
}
```

### Complete Configuration

```hcl
// HXE Configuration
debug = false
version = "0.1.0"
log_level = "info"
log_file = "/var/log/hxe.log"

api {
  addr = "0.0.0.0"
  port = 8080
  username = "admin"
  password = "password"
  jwt_secret = "your-secret-key-here"
  jwt_expiry = "24h"
  cors_origins = ["http://localhost:3000", "https://yourdomain.com"]
  rate_limit = 100
  rate_limit_window = "1m"
}

database {
  type = "sqlite"
  path = "/var/lib/hxe/hxe.db"
  migrate = true
  max_connections = 10
  connection_timeout = "30s"
}

broker { 
  name = "hxe"
  addr = "0.0.0.0"
  port = 4222
  cluster_name = "hxe-cluster"
  max_payload = "1MB"
  max_connections = 100
  ping_interval = "2m"
  ping_max_out = 3
}

agent {
  name = "hxe-agent"
  hostname = "localhost"
  port = 4223
  heartbeat_interval = "30s"
  reconnect_interval = "5s"
  max_reconnect_attempts = 10
}

programs {
  default_user = "nobody"
  default_group = "nobody"
  default_directory = "/tmp"
  max_restarts = 3
  restart_delay = "5s"
  log_rotation = {
    max_size = "100MB"
    max_files = 5
    max_age = "30d"
  }
}

security {
  tls_enabled = false
  tls_cert_file = "/path/to/cert.pem"
  tls_key_file = "/path/to/key.pem"
  tls_ca_file = "/path/to/ca.pem"
  allowed_hosts = ["localhost", "127.0.0.1"]
  firewall_rules = [
    {
      action = "allow"
      source = "192.168.1.0/24"
      port = 8080
    }
  ]
}
```

## Configuration Sections

### Global Settings

```hcl
// Global configuration
debug = false                    // Enable debug logging
version = "0.1.0"               // HXE version
log_level = "info"              // Log level: debug, info, warn, error
log_file = "/var/log/hxe.log"   // Log file path
```

### API Configuration

```hcl
api {
  addr = "0.0.0.0"              // API server address
  port = 8080                   // API server port
  username = "admin"            // Default username
  password = "password"         // Default password
  jwt_secret = "secret-key"     // JWT signing secret
  jwt_expiry = "24h"            // JWT token expiry
  cors_origins = ["*"]          // CORS allowed origins
  rate_limit = 100              // Rate limit per window
  rate_limit_window = "1m"      // Rate limit window
  timeout = "30s"               // Request timeout
  max_body_size = "10MB"        // Maximum request body size
}
```

### Database Configuration

```hcl
database {
  type = "sqlite"               // Database type: sqlite, postgres, mysql
  path = "/var/lib/hxe/hxe.db"  // Database file path (SQLite)
  host = "localhost"            // Database host (PostgreSQL/MySQL)
  port = 5432                   // Database port (PostgreSQL/MySQL)
  name = "hxe"                  // Database name (PostgreSQL/MySQL)
  username = "hxe"              // Database username (PostgreSQL/MySQL)
  password = "password"         // Database password (PostgreSQL/MySQL)
  migrate = true                // Run database migrations
  max_connections = 10          // Maximum database connections
  connection_timeout = "30s"    // Connection timeout
  ssl_mode = "disable"          // SSL mode (PostgreSQL)
}
```

### Broker Configuration (NATS)

```hcl
broker { 
  name = "hxe"                  // Broker name
  addr = "0.0.0.0"             // Broker address
  port = 4222                   // Broker port
  cluster_name = "hxe-cluster"  // Cluster name
  max_payload = "1MB"           // Maximum message payload
  max_connections = 100         // Maximum client connections
  ping_interval = "2m"          // Ping interval
  ping_max_out = 3              // Maximum ping failures
  auth_required = false         // Require authentication
  username = "nats"             // Broker username
  password = "password"         // Broker password
}
```

### Agent Configuration

```hcl
agent {
  name = "hxe-agent"            // Agent name
  hostname = "localhost"        // Agent hostname
  port = 4223                   // Agent port
  heartbeat_interval = "30s"    // Heartbeat interval
  reconnect_interval = "5s"     // Reconnect interval
  max_reconnect_attempts = 10   // Maximum reconnect attempts
  log_level = "info"            // Agent log level
  metrics_enabled = true        // Enable metrics collection
  metrics_port = 9090           // Metrics server port
}
```

### Programs Configuration

```hcl
programs {
  default_user = "nobody"       // Default program user
  default_group = "nobody"      // Default program group
  default_directory = "/tmp"    // Default working directory
  max_restarts = 3              // Maximum restart attempts
  restart_delay = "5s"          // Delay between restarts
  log_rotation = {
    max_size = "100MB"          // Maximum log file size
    max_files = 5               // Maximum log files to keep
    max_age = "30d"             // Maximum log file age
    compress = true             // Compress old log files
  }
  environment = {               // Default environment variables
    PATH = "/usr/local/bin:/usr/bin:/bin"
    HOME = "/tmp"
  }
}
```

### Security Configuration

```hcl
security {
  tls_enabled = false           // Enable TLS
  tls_cert_file = "/path/to/cert.pem"  // TLS certificate file
  tls_key_file = "/path/to/key.pem"    // TLS private key file
  tls_ca_file = "/path/to/ca.pem"      // TLS CA certificate file
  allowed_hosts = ["localhost", "127.0.0.1"]  // Allowed hosts
  firewall_rules = [            // Firewall rules
    {
      action = "allow"
      source = "192.168.1.0/24"
      port = 8080
    },
    {
      action = "deny"
      source = "0.0.0.0/0"
      port = 22
    }
  ]
}
```

## Environment Variables

HXE supports environment variables for configuration:

```bash
# API Configuration
export HXE_API_ADDR="0.0.0.0"
export HXE_API_PORT="8080"
export HXE_API_USERNAME="admin"
export HXE_API_PASSWORD="password"

# Database Configuration
export HXE_DB_TYPE="sqlite"
export HXE_DB_PATH="/var/lib/hxe/hxe.db"

# Broker Configuration
export HXE_BROKER_ADDR="0.0.0.0"
export HXE_BROKER_PORT="4222"

# Logging
export HXE_LOG_LEVEL="info"
export HXE_LOG_FILE="/var/log/hxe.log"
```

## Configuration Validation

HXE validates configuration files on startup:

```bash
# Validate configuration
hxe --config config.hcl --validate

# Check configuration syntax
hxe --config config.hcl --check
```

## Configuration Examples

### Development Configuration

```hcl
debug = true
log_level = "debug"

api {
  addr = "localhost"
  port = 8080
  username = "admin"
  password = "password"
  cors_origins = ["http://localhost:3000"]
}

database {
  type = "sqlite"
  path = "./hxe-dev.db"
  migrate = true
}

broker { 
  name = "hxe-dev"
  addr = "localhost"
  port = 4222
}
```

### Production Configuration

```hcl
debug = false
log_level = "info"
log_file = "/var/log/hxe.log"

api {
  addr = "0.0.0.0"
  port = 8080
  username = "admin"
  password = "secure-password"
  jwt_secret = "your-super-secret-key"
  jwt_expiry = "24h"
  cors_origins = ["https://yourdomain.com"]
  rate_limit = 100
  rate_limit_window = "1m"
}

database {
  type = "postgres"
  host = "localhost"
  port = 5432
  name = "hxe"
  username = "hxe"
  password = "secure-db-password"
  migrate = true
  max_connections = 20
}

broker { 
  name = "hxe-prod"
  addr = "0.0.0.0"
  port = 4222
  cluster_name = "hxe-cluster"
  auth_required = true
  username = "nats"
  password = "secure-broker-password"
}

security {
  tls_enabled = true
  tls_cert_file = "/etc/ssl/certs/hxe.crt"
  tls_key_file = "/etc/ssl/private/hxe.key"
  allowed_hosts = ["yourdomain.com", "192.168.1.0/24"]
}
```

### Docker Configuration

```hcl
debug = false
log_level = "info"

api {
  addr = "0.0.0.0"
  port = 8080
  username = "admin"
  password = "password"
}

database {
  type = "sqlite"
  path = "/data/hxe.db"
  migrate = true
}

broker { 
  name = "hxe"
  addr = "0.0.0.0"
  port = 4222
}

agent {
  name = "hxe-agent"
  hostname = "hxe-agent"
  port = 4223
}
```

## Troubleshooting

### Common Configuration Issues

#### Invalid HCL Syntax
```bash
# Check HCL syntax
hclfmt -check config.hcl
```

#### Missing Required Fields
```bash
# HXE will show specific error messages for missing fields
hxe --config config.hcl
```

#### Permission Issues
```bash
# Ensure proper file permissions
chmod 600 config.hcl
chown hxe:hxe config.hcl
```

#### Database Connection Issues
```bash
# Test database connection
hxe --config config.hcl --test-db
```

### Configuration Validation

```bash
# Validate configuration file
hxe --config config.hcl --validate

# Show configuration summary
hxe --config config.hcl --show-config
```

## Best Practices

1. **Security**: Use strong passwords and JWT secrets
2. **Backup**: Regularly backup configuration files
3. **Version Control**: Use version control for configuration
4. **Environment**: Use different configs for dev/staging/prod
5. **Validation**: Always validate configuration before deployment
6. **Documentation**: Document custom configuration changes

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 