# API Reference

The HXE API provides a comprehensive RESTful interface for managing programs and interacting with the HXE server. All endpoints support JWT authentication and return JSON responses.

## Base URL

```
http://localhost:8080/api
```

## Authentication

HXE uses JWT (JSON Web Tokens) for authentication. Include the token in the `Authorization` header:

```
Authorization: Bearer <your-jwt-token>
```

### Default Credentials
- **Username**: `admin`
- **Password**: `password`

## Authentication Endpoints

### Login

**POST** `/api/auth/login`

Authenticate and receive a JWT token.

**Request Body:**
```json
{
  "username": "admin",
  "password": "password"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2025-01-01T12:00:00Z",
  "user": {
    "username": "admin",
    "role": "admin"
  }
}
```

**Status Codes:**
- `200` - Authentication successful
- `401` - Invalid credentials
- `400` - Invalid request body

### Refresh Token

**POST** `/api/auth/refresh`

Refresh the current JWT token.

**Headers:**
```
Authorization: Bearer <current-token>
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expires_at": "2025-01-01T12:00:00Z"
}
```

**Status Codes:**
- `200` - Token refreshed successfully
- `401` - Invalid or expired token

### Logout

**POST** `/api/auth/logout`

Logout and invalidate the current token.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Logged out successfully"
}
```

**Status Codes:**
- `200` - Logout successful
- `401` - Invalid token

## Program Management Endpoints

### List Programs

**GET** `/api/program`

Retrieve all programs.

**Headers:**
```
Authorization: Bearer <token>
```

**Query Parameters:**
- `status` - Filter by status (`running`, `stopped`, `error`)
- `enabled` - Filter by enabled status (`true`, `false`)
- `limit` - Maximum number of programs to return (default: 100)
- `offset` - Number of programs to skip (default: 0)

**Response:**
```json
[
  {
    "id": "program-1",
    "name": "Web Server",
    "description": "Nginx web server",
    "command": "/usr/sbin/nginx",
    "args": "-g 'daemon off;'",
    "directory": "/var/www",
    "user": "www-data",
    "group": "www-data",
    "enabled": true,
    "auto_restart": true,
    "max_restarts": 3,
    "status": "running",
    "pid": 12345,
    "created_at": "2025-01-01T00:00:00Z",
    "updated_at": "2025-01-01T00:00:00Z"
  }
]
```

**Status Codes:**
- `200` - Programs retrieved successfully
- `401` - Authentication required

### Get Program

**GET** `/api/program/{id}`

Retrieve a specific program by ID.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "program-1",
  "name": "Web Server",
  "description": "Nginx web server",
  "command": "/usr/sbin/nginx",
  "args": "-g 'daemon off;'",
  "directory": "/var/www",
  "user": "www-data",
  "group": "www-data",
  "enabled": true,
  "auto_restart": true,
  "max_restarts": 3,
  "status": "running",
  "pid": 12345,
  "exit_code": 0,
  "start_time": "2025-01-01T00:00:00Z",
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

**Status Codes:**
- `200` - Program retrieved successfully
- `404` - Program not found
- `401` - Authentication required

### Create Program

**POST** `/api/program`

Create a new program.

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "My Program",
  "description": "A test program",
  "command": "/usr/bin/python3",
  "args": "script.py",
  "directory": "/tmp",
  "user": "nobody",
  "group": "nobody",
  "enabled": true,
  "auto_restart": true,
  "max_restarts": 3
}
```

**Response:**
```json
{
  "id": "program-2",
  "name": "My Program",
  "description": "A test program",
  "command": "/usr/bin/python3",
  "args": "script.py",
  "directory": "/tmp",
  "user": "nobody",
  "group": "nobody",
  "enabled": true,
  "auto_restart": true,
  "max_restarts": 3,
  "status": "stopped",
  "pid": null,
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

**Status Codes:**
- `201` - Program created successfully
- `400` - Invalid request body
- `409` - Program with same name already exists
- `401` - Authentication required

### Update Program

**PUT** `/api/program/{id}`

Update an existing program.

**Headers:**
```
Authorization: Bearer <token>
Content-Type: application/json
```

**Request Body:**
```json
{
  "name": "Updated Program",
  "description": "Updated description",
  "command": "/usr/bin/python3",
  "args": "updated_script.py",
  "directory": "/tmp",
  "user": "nobody",
  "group": "nobody",
  "enabled": true,
  "auto_restart": true,
  "max_restarts": 5
}
```

**Response:**
```json
{
  "id": "program-1",
  "name": "Updated Program",
  "description": "Updated description",
  "command": "/usr/bin/python3",
  "args": "updated_script.py",
  "directory": "/tmp",
  "user": "nobody",
  "group": "nobody",
  "enabled": true,
  "auto_restart": true,
  "max_restarts": 5,
  "status": "running",
  "pid": 12345,
  "created_at": "2025-01-01T00:00:00Z",
  "updated_at": "2025-01-01T00:00:00Z"
}
```

**Status Codes:**
- `200` - Program updated successfully
- `400` - Invalid request body
- `404` - Program not found
- `401` - Authentication required

### Delete Program

**DELETE** `/api/program/{id}`

Delete a program.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "message": "Program deleted successfully"
}
```

**Status Codes:**
- `200` - Program deleted successfully
- `404` - Program not found
- `401` - Authentication required

## Runtime Control Endpoints

### Start Program

**POST** `/api/program/{id}/start`

Start a program.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "program-1",
  "status": "running",
  "pid": 12345,
  "start_time": "2025-01-01T00:00:00Z",
  "message": "Program started successfully"
}
```

**Status Codes:**
- `200` - Program started successfully
- `400` - Program already running
- `404` - Program not found
- `401` - Authentication required

### Stop Program

**POST** `/api/program/{id}/stop`

Stop a program.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "program-1",
  "status": "stopped",
  "pid": null,
  "exit_code": 0,
  "message": "Program stopped successfully"
}
```

**Status Codes:**
- `200` - Program stopped successfully
- `400` - Program not running
- `404` - Program not found
- `401` - Authentication required

### Restart Program

**POST** `/api/program/{id}/restart`

Restart a program.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "program-1",
  "status": "running",
  "pid": 12346,
  "start_time": "2025-01-01T00:00:00Z",
  "message": "Program restarted successfully"
}
```

**Status Codes:**
- `200` - Program restarted successfully
- `404` - Program not found
- `401` - Authentication required

### Enable Autostart

**POST** `/api/program/{id}/enable`

Enable autostart for a program.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "program-1",
  "enabled": true,
  "message": "Autostart enabled successfully"
}
```

**Status Codes:**
- `200` - Autostart enabled successfully
- `404` - Program not found
- `401` - Authentication required

### Disable Autostart

**POST** `/api/program/{id}/disable`

Disable autostart for a program.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "id": "program-1",
  "enabled": false,
  "message": "Autostart disabled successfully"
}
```

**Status Codes:**
- `200` - Autostart disabled successfully
- `404` - Program not found
- `401` - Authentication required

## Monitoring Endpoints

### Get Program Logs

**GET** `/api/program/{id}/logs`

Retrieve program logs.

**Headers:**
```
Authorization: Bearer <token>
```

**Query Parameters:**
- `lines` - Number of lines to return (default: 100)
- `follow` - Follow logs in real-time (`true`, `false`)
- `level` - Filter by log level (`debug`, `info`, `warn`, `error`)

**Response:**
```json
{
  "program_id": "program-1",
  "logs": [
    {
      "timestamp": "2025-01-01T00:00:00Z",
      "level": "info",
      "message": "Program started successfully"
    },
    {
      "timestamp": "2025-01-01T00:00:01Z",
      "level": "info",
      "message": "Listening on port 8080"
    }
  ]
}
```

**Status Codes:**
- `200` - Logs retrieved successfully
- `404` - Program not found
- `401` - Authentication required

### Get Program Metrics

**GET** `/api/program/{id}/metrics`

Retrieve program metrics.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "program_id": "program-1",
  "metrics": {
    "cpu_usage": 2.5,
    "memory_usage": 1024,
    "uptime": 3600,
    "restart_count": 0,
    "last_restart": null
  }
}
```

**Status Codes:**
- `200` - Metrics retrieved successfully
- `404` - Program not found
- `401` - Authentication required

### Get System Metrics

**GET** `/api/metrics`

Retrieve system-wide metrics.

**Headers:**
```
Authorization: Bearer <token>
```

**Response:**
```json
{
  "system": {
    "cpu_usage": 15.2,
    "memory_usage": 8192,
    "disk_usage": 75.5,
    "uptime": 86400
  },
  "programs": {
    "total": 10,
    "running": 8,
    "stopped": 2,
    "error": 0
  }
}
```

**Status Codes:**
- `200` - Metrics retrieved successfully
- `401` - Authentication required

## Error Responses

All endpoints return consistent error responses:

### Validation Error

```json
{
  "error": "validation_error",
  "message": "Invalid request body",
  "details": {
    "field": "name",
    "issue": "Name is required"
  }
}
```

### Authentication Error

```json
{
  "error": "authentication_error",
  "message": "Invalid or expired token"
}
```

### Not Found Error

```json
{
  "error": "not_found",
  "message": "Program not found",
  "resource": "program",
  "id": "program-1"
}
```

### Server Error

```json
{
  "error": "internal_error",
  "message": "Internal server error",
  "request_id": "req-12345"
}
```

## Rate Limiting

The API implements rate limiting to prevent abuse:

- **Default Limit**: 100 requests per minute per IP
- **Headers**: Rate limit information is included in response headers
  - `X-RateLimit-Limit`: Maximum requests per window
  - `X-RateLimit-Remaining`: Remaining requests in current window
  - `X-RateLimit-Reset`: Time when the rate limit resets

**Rate Limit Exceeded Response:**
```json
{
  "error": "rate_limit_exceeded",
  "message": "Rate limit exceeded",
  "retry_after": 60
}
```

## CORS Support

The API supports Cross-Origin Resource Sharing (CORS):

- **Allowed Origins**: Configurable via configuration
- **Methods**: GET, POST, PUT, DELETE, OPTIONS
- **Headers**: Authorization, Content-Type
- **Credentials**: Supported

## WebSocket Support

For real-time updates, the API supports WebSocket connections:

**WebSocket URL:**
```
ws://localhost:8080/api/ws
```

**Authentication:**
```json
{
  "type": "auth",
  "token": "your-jwt-token"
}
```

**Event Types:**
- `program.started` - Program started
- `program.stopped` - Program stopped
- `program.restarted` - Program restarted
- `program.error` - Program error
- `log.new` - New log entry

## SDKs and Libraries

### Go Client

```go
import "github.com/rangertaha/hxe/pkg/client"

// Create client
hxeClient := client.NewAuthenticatedClient("http://localhost:8080", "admin", "password")

// Login
_, err := hxeClient.Login()

// Use the client
programs, err := hxeClient.Program.ListPrograms()
```

### cURL Examples

```bash
# Login
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"password"}'

# List programs
curl -X GET http://localhost:8080/api/program \
  -H "Authorization: Bearer <token>"

# Create program
curl -X POST http://localhost:8080/api/program \
  -H "Authorization: Bearer <token>" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Program",
    "command": "/usr/bin/echo",
    "args": "Hello, World!"
  }'

# Start program
curl -X POST http://localhost:8080/api/program/program-1/start \
  -H "Authorization: Bearer <token>"
```

## Versioning

The API is versioned through the URL path:

- **Current Version**: `/api/v1/` (default)
- **Future Versions**: `/api/v2/`, `/api/v3/`, etc.

## Support

- 📧 Email: rangertaha@gmail.com
- 🐛 Issues: [GitHub Issues](https://github.com/rangertaha/hxe/issues)
- 📖 Documentation: [GitHub Wiki](https://github.com/rangertaha/hxe/wiki) 