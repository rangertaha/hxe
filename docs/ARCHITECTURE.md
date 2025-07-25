# HXE Architecture Plan

## Overview
Host-based Process Execution Engine (HXE) is a service management platform that provides process lifecycle management, metrics collection, alerting, and health monitoring. The system uses HCL (HashiCorp Configuration Language) for configuration management and supports both internal service-to-service communication via NATS and external API access via HTTP.

## Project Structure

```markdown:docs/ARCHITECTURE.md
```
hxe/
├── cmd/
│   ├── hxe/           # Main CLI binary
├── internal/
│   ├── config/        # HCL configuration management
│   ├── agent/         # Agent implementation
│   ├── client/        # Internal service-to-service client
│   └── services/      # Service modules (models, handlers, clients)
├── pkg/
│   └── client/        # External HTTP API client
├── docs/              # Documentation
└── examples/          # Example configurations and usage
```

## Core Components

### 1. Configuration Management (`internal/config/`)
- **HCL Parser**: Parse HCL configuration files using HashiCorp's HCL library
- **Config Structs**: Go structs representing HCL configuration blocks
- **Validation**: Validate HCL configuration syntax and semantic values
- **Environment Overrides**: Support for environment variable overrides
- **Hot Reload**: Watch for configuration changes and reload without restart
- **Default Values**: Provide sensible defaults for configuration options

**Key Features:**
- Multi-file configuration support
- Configuration inheritance and merging
- Type-safe configuration access
- Configuration validation with detailed error messages

### 2. Agent Implementation (`internal/agent/`)
The agent is responsible for running on each host and managing local services.

**Core Responsibilities:**
- **Process Manager**: Lifecycle management of services (start, stop, restart, monitor)
- **Metrics Collector**: System and service metrics collection (CPU, memory, disk, network)
- **Alert Manager**: Alert generation and management based on thresholds
- **Health Monitor**: Service health monitoring and status reporting
- **Config Manager**: HCL configuration management and hot reloading
- **Service Communication**: NATS-based communication with other services

**Key Features:**
- Graceful service shutdown
- Automatic service recovery
- Resource usage monitoring
- Alert aggregation and deduplication
- Service dependency management

### 3. Internal Clients (`internal/client/`)
- **Service Clients**: Import and use service clients for internal communication
- **NATS Integration**: Direct NATS messaging for service-to-service communication
- **Connection Pooling**: Efficient connection management
- **Retry Logic**: Automatic retry with exponential backoff
- **Circuit Breaker**: Fault tolerance patterns

**Usage:**
```go
// Internal service-to-service communication
import "github.com/rangertaha/hxe/internal/client"

servicesClient := client.NewServicesClient(natsConn)
service, err := servicesClient.Get(ctx, "service-id")
```

### 4. External Client (`pkg/client/`)
- **HTTP API Client**: External applications use this via HTTP API
- **Authentication**: API key/token management
- **Request/Response**: HTTP request/response handling
- **Rate Limiting**: Client-side rate limiting
- **Retry Logic**: Automatic retry for transient failures

**Usage:**
```go
// External application integration
import "github.com/rangertaha/hxe/pkg/client"

client := hxe.NewClient("http://localhost:8080", "api-key")
service, err := client.Services.Get(ctx, "service-id")
```

### 5. Services (`internal/services/`)
Organized service modules with clear separation of concerns.

**Structure per service:**
```
internal/services/
├── models/            # Data structures and types
├── handlers/          # HTTP API handlers
├── clients/           # Service-specific clients
├── seeders/           # Data seeding functions
├── service.go         # Service business logic
└── server.go          # Service server setup
```

**Current Services:**
- **Services**: Process and service management
- **Groups**: Service grouping and organization
- **Tags**: Service tagging and categorization
- **Fields**: Custom field definitions
- **Variables**: Environment variable management
- **Credentials**: Secure credential storage
- **Apps**: Application definitions

## Configuration Format (HCL)

### Agent Configuration (`configs/agent.hcl`)
```hcl
agent {
  name = "hxe-agent-01"
  data_dir = "/var/lib/hxe"
  log_level = "info"
  
  metrics {
    enabled = true
    interval = "30s"
    retention = "24h"
    prometheus_endpoint = ":9090"
  }
  
  alerts {
    enabled = true
    webhook_url = "https://hooks.slack.com/..."
    email {
      smtp_server = "smtp.gmail.com:587"
      username = "alerts@company.com"
      password = "env:SMTP_PASSWORD"
    }
  }
  
  nats {
    servers = ["nats://localhost:4222"]
    cluster_name = "hxe-cluster"
  }
}
```

### Services Configuration (`configs/services.hcl`)
```hcl
services {
  service "web-app" {
    command = "node app.js"
    working_dir = "/opt/web-app"
    user = "www-data"
    
    environment = {
      NODE_ENV = "production"
      PORT = "3000"
      DATABASE_URL = "env:DATABASE_URL"
    }
    
    health_check {
      type = "http"
      url = "http://localhost:3000/health"
      interval = "10s"
      timeout = "5s"
      retries = 3
    }
    
    resources {
      cpu_limit = "1000m"
      memory_limit = "1Gi"
      restart_policy = "always"
    }
    
    tags = ["web", "frontend", "production"]
  }
  
  service "database" {
    command = "postgres -D /var/lib/postgresql/data"
    working_dir = "/var/lib/postgresql"
    user = "postgres"
    
    health_check {
      type = "tcp"
      host = "localhost"
      port = 5432
      interval = "30s"
    }
    
    resources {
      cpu_limit = "2000m"
      memory_limit = "4Gi"
      restart_policy = "on-failure"
    }
    
    tags = ["database", "postgresql", "production"]
  }
}
```

### Server Configuration (`configs/server.hcl`)
```hcl
server {
  host = "0.0.0.0"
  port = 8080
  
  tls {
    enabled = true
    cert_file = "/etc/hxe/certs/server.crt"
    key_file = "/etc/hxe/certs/server.key"
  }
  
  auth {
    type = "jwt"
    secret = "env:JWT_SECRET"
    token_expiry = "24h"
  }
  
  cors {
    allowed_origins = ["https://app.company.com"]
    allowed_methods = ["GET", "POST", "PUT", "DELETE"]
  }
  
  rate_limit {
    requests_per_minute = 1000
    burst_size = 100
  }
}
```

## Communication Patterns

### Internal Communication (NATS)
- **Service Discovery**: Automatic service discovery via NATS
- **Event Streaming**: Real-time events and metrics
- **Request/Response**: Synchronous service calls
- **Pub/Sub**: Asynchronous event publishing

### External Communication (HTTP)
- **RESTful API**: Standard REST endpoints
- **WebSocket**: Real-time updates for dashboards
- **GraphQL**: Optional GraphQL endpoint for complex queries
- **gRPC**: High-performance RPC for internal services

## Security Considerations

### Authentication & Authorization
- **JWT Tokens**: Stateless authentication
- **API Keys**: Simple API key authentication
- **OAuth2**: Integration with external identity providers
- **RBAC**: Role-based access control

### Data Protection
- **Encryption at Rest**: Sensitive data encryption
- **TLS**: Transport layer security
- **Secrets Management**: Secure credential storage
- **Audit Logging**: Comprehensive audit trails

## Monitoring & Observability

### Metrics
- **System Metrics**: CPU, memory, disk, network
- **Service Metrics**: Response times, error rates, throughput
- **Business Metrics**: Custom application metrics
- **Prometheus Integration**: Standard metrics format

### Logging
- **Structured Logging**: JSON format with correlation IDs
- **Log Levels**: Debug, info, warn, error
- **Log Aggregation**: Centralized log collection
- **Log Retention**: Configurable retention policies

### Tracing
- **Distributed Tracing**: Request tracing across services
- **Jaeger Integration**: OpenTelemetry compatible
- **Performance Profiling**: Service performance analysis

## Deployment

### Single Binary
- **Static Linking**: Single executable with all dependencies
- **Cross-Platform**: Support for Linux, macOS, Windows
- **Minimal Dependencies**: Minimal runtime dependencies

### Container Support
- **Docker Images**: Official Docker images
- **Kubernetes**: Helm charts and manifests
- **Service Mesh**: Istio/Envoy integration

### Configuration Management
- **Environment Variables**: Runtime configuration
- **Configuration Files**: HCL configuration files
- **Secrets**: External secrets management
- **Dynamic Configuration**: Hot reloading support

## Development Workflow

### Local Development
- **Hot Reload**: Automatic restart on code changes
- **Mock Mode**: Development without external dependencies
- **Debug Mode**: Enhanced logging and debugging
- **Test Mode**: Isolated testing environment

### Testing Strategy
- **Unit Tests**: Individual component testing
- **Integration Tests**: Service interaction testing
- **End-to-End Tests**: Full system testing
- **Performance Tests**: Load and stress testing

### CI/CD Pipeline
- **Automated Testing**: Run tests on every commit
- **Code Quality**: Linting and static analysis
- **Security Scanning**: Vulnerability scanning
- **Automated Deployment**: Deployment automation

## Future Enhancements

### Planned Features
- **Multi-Tenancy**: Support for multiple organizations
- **Service Mesh**: Advanced service networking
- **Machine Learning**: Predictive analytics and auto-scaling
- **Plugin System**: Extensible plugin architecture
- **Multi-Cloud**: Cloud provider integration

### Scalability
- **Horizontal Scaling**: Multi-node deployments
- **Load Balancing**: Intelligent load distribution
- **Caching**: Multi-level caching strategy
- **Database Sharding**: Data distribution across nodes

This architecture provides a solid foundation for a scalable, maintainable, and secure service management platform with clear separation of concerns and modern development practices.
```

This updated plan document now properly reflects the HCL configuration format and provides comprehensive details about the architecture, including configuration examples, security considerations, monitoring strategies, and deployment approaches. 