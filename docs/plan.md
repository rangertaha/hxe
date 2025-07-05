# HXE Project Plan

## Current State Analysis

### Project Overview
HXE (Host-based Process Execution Engine) is a service management system with the following components:
- **CLI Interface**: Command-line tool for service management
- **API Client**: HTTP API for communicating to the API server, uses JWT based authentication
- **API Server**: HTTP API for service operations
- **Database**: SQLite storage for API server to use.

### Architecture
```
┌─────────────┐      ┌──────────────┐      ┌──────────────┐
│    CLI      │ ───► │ Client Lib   │ ───► │  API Server  │
│ (cmd/hxe)   │      │ (pkg/client) │      │ (internal/api)  │
└─────────────┘      └──────────────┘      └─────┬────────┘
                                                  │
                                                  ▼
                                         ┌───────────────┐
                                         │  SQLite DB    │
                                         │ (internal/db) │
                                         └───────────────┘
```

## Current Issues & Status

### ✅ Completed Components
1. **Basic CLI Structure**: Command framework with service subcommands
2. **NATS Integration**: Embedded NATS server for messaging
3. **Database Schema**: SQLite database with service models
4. **API Framework**: Echo-based HTTP server structure
5. **Configuration System**: HCL-based configuration management

### ❌ Critical Issues to Fix

#### 1. Service Handler Methods (internal/api/handlers/service.go)
**Problem**: Most handler methods are commented out, causing undefined method errors
**Impact**: API endpoints return 404/500 errors
**Files Affected**: 


**Required Actions**:
- Uncomment and implement missing handler methods
- Fix method signatures to match expected interfaces
- Implement proper error handling

#### 2. Service Client Methods (pkg/client/services.go)
**Problem**: Most client methods are commented out
**Impact**: CLI commands fail with undefined method errors
**Files Affected**: `pkg/client/services.go`

**Required Actions**:
- Uncomment and implement client methods
- Fix return types and error handling
- Ensure proper HTTP request/response handling



**Required Actions**:
- Complete service endpoint implementations
- Fix JSON handler wrapper
- Implement proper storage operations

#### 4. Type System Inconsistencies
**Problem**: Multiple conflicting type definitions
**Impact**: Compilation errors and runtime issues
**Files Affected**: Multiple files across the codebase

**Required Actions**:
- Standardize on single type system
- Remove duplicate type definitions
- Fix import conflicts

## Implementation Plan

### Phase 1: Fix Core Infrastructure (Priority: High)

#### 1.1 Complete Service Handlers
```go
// internal/api/handlers/service.go
func (s *Service) Get(c echo.Context) error
func (s *Service) Create(c echo.Context) error
func (s *Service) Update(c echo.Context) error
func (s *Service) Delete(c echo.Context) error
func (s *Service) Start(c echo.Context) error
func (s *Service) Stop(c echo.Context) error
func (s *Service) Restart(c echo.Context) error
func (s *Service) Status(c echo.Context) error
func (s *Service) Reload(c echo.Context) error
func (s *Service) Enable(c echo.Context) error
func (s *Service) Disable(c echo.Context) error
func (s *Service) Shell(c echo.Context) error
func (s *Service) Log(c echo.Context) error
```

#### 1.2 Complete Service Client Methods
```go
// pkg/client/services.go
func (c *ServiceClient) Get(id string) (*services.Response, error)
func (c *ServiceClient) Create(req *services.Request) (*services.Response, error)
func (c *ServiceClient) Update(req *services.Request) (*services.Response, error)
func (c *ServiceClient) Delete(id string) (*services.Response, error)
func (c *ServiceClient) Start(id string) (*services.Response, error)
func (c *ServiceClient) Stop(id string) (*services.Response, error)
func (c *ServiceClient) Restart(id string) (*services.Response, error)
func (c *ServiceClient) Status(id string) (*services.Response, error)
func (c *ServiceClient) Reload(id string) (*services.Response, error)
func (c *ServiceClient) Enable(id string) (*services.Response, error)
func (c *ServiceClient) Disable(id string) (*services.Response, error)
func (c *ServiceClient) Shell(id string) (*services.Response, error)
func (c *ServiceClient) Log(id string) (*services.Response, error)
```



### Phase 2: Enable CLI Commands (Priority: High)

#### 2.1 Uncomment CLI Commands
- Uncomment all service subcommands in `cmd/hxe/service.go`
- Fix method calls to use correct client methods
- Implement proper error handling and output formatting

#### 2.2 Test Basic Operations
- `hxe service list` - List all services
- `hxe service start <id>` - Start a service
- `hxe service stop <id>` - Stop a service
- `hxe service status <id>` - Check service status

### Phase 3: Database Integration (Priority: Medium)

#### 3.1 Complete Storage Layer
- Implement proper CRUD operations in storage layer
- Add service persistence and retrieval
- Implement service state management

#### 3.2 Add Service Lifecycle Management
- Service creation and configuration
- Service state tracking (running, stopped, error)
- Service dependency management

#### 4.2 Add Logging and Monitoring
- Service log collection and streaming
- Process monitoring and health checks
- Error reporting and recovery

### Phase 5: Advanced Features (Priority: Low)

#### 5.1 Shell Integration
- Interactive shell access to services
- Command execution in service context

#### 5.2 Configuration Management
- Service configuration reloading
- Dynamic configuration updates

#### 5.3 Security Features
- Authentication and authorization
- Service isolation and sandboxing

## Testing Strategy

### Unit Tests
- Test individual components in isolation
- Mock external dependencies (NATS, database)
- Test error conditions and edge cases

### Integration Tests
- Test complete request/response cycles
- Test CLI command execution
- Test API endpoint functionality

### End-to-End Tests
- Test complete service lifecycle
- Test multiple concurrent operations
- Test system recovery scenarios

## Success Criteria

### Phase 1 Success
- [ ] All API endpoints return proper responses
- [ ] CLI commands execute without errors
- [ ] NATS messaging works end-to-end
- [ ] No compilation errors

### Phase 2 Success
- [ ] Basic service operations work (list, start, stop, status)
- [ ] CLI provides meaningful output
- [ ] Error handling is robust

### Phase 3 Success
- [ ] Services persist across restarts
- [ ] Service state is properly tracked
- [ ] Database operations are reliable

### Phase 4 Success
- [ ] Process execution works correctly
- [ ] Service monitoring is functional
- [ ] Logging provides useful information

## Risk Assessment

### High Risk
- **Type System Conflicts**: Multiple competing type definitions could cause major refactoring
- **NATS Integration**: Complex messaging patterns may be difficult to debug
- **Process Management**: OS-level process handling can be platform-specific

### Medium Risk
- **Database Schema**: Changes to service models may require migrations
- **API Design**: REST API design may need iteration based on usage

### Low Risk
- **Configuration**: HCL configuration is well-established
- **CLI Framework**: urfave/cli is mature and stable

## Timeline Estimate

- **Phase 1**: 2-3 days (critical infrastructure)
- **Phase 2**: 1-2 days (CLI functionality)
- **Phase 3**: 2-3 days (database integration)
- **Phase 4**: 3-4 days (process management)
- **Phase 5**: 1-2 weeks (advanced features)

**Total Estimated Time**: 2-3 weeks for core functionality

## Next Steps

1. **Immediate**: Fix service handler methods in `internal/api/handlers/service.go`
2. **Next**: Complete service client methods in `pkg/client/services.go`
3. **Then**: Implement service module endpoints in `internal/modules/services/service.go`
4. **Finally**: Test and validate end-to-end functionality

## Notes

- Focus on getting basic CRUD operations working first
- Use existing database schema and models
- Maintain backward compatibility where possible
- Document any breaking changes clearly
- Test thoroughly at each phase before proceeding 

## HXE Architecture (No NATS)

**Flow:**
```
CLI  →  Client Library  →  API Server  →  SQLite Database
```

### Description

- **CLI**: The user interacts with the system via command-line commands (e.g., `hxe service list`).
- **Client Library**: The CLI uses a Go client library to format requests and handle responses.
- **API Server**: The client library communicates with the API server (HTTP/REST), which exposes endpoints for service management.
- **SQLite Database**: The API server performs all CRUD operations and service state management by reading from and writing to the SQLite database.

### Example Request Flow

1. **User runs a CLI command** (e.g., `hxe service start <id>`).
2. **CLI calls the client library** to build and send an HTTP request.
3. **Client library sends the request** to the API server.
4. **API server processes the request**, interacts with the SQLite database, and returns a response.
5. **Client library receives the response** and displays the result in the CLI.

---

**This architecture is simple, robust, and easy to debug. All communication is via HTTP and the database, with no message queue or NATS dependency.**

Let me know if you want this added to your plan or as a separate doc section! 

## Group Resource: Hierarchical Structure

### Concept

- **Group** is a resource that can belong to a parent group.
- This allows you to build a tree or hierarchy of groups (e.g., categories, organizational units).

### GORM Model Example

```go
import "gorm.io/gorm"

type Group struct {
    ID          uint      `gorm:"primaryKey"`
    Title       string    `gorm:"not null"`
    Icon        string
    Description string
    ParentID    *uint
    Parent      *Group    `gorm:"foreignKey:ParentID"`
    Children    []Group   `gorm:"foreignKey:ParentID"`
    gorm.Model  // includes CreatedAt, UpdatedAt, DeletedAt
}
```

- `ParentID` is a pointer to allow `NULL` (no parent).
- `Parent` and `Children` set up the self-referential relationship for GORM.

### Example Usage

**Create a top-level group:**
```go
root := Group{Title: "Root", Icon: "root.png", Description: "Top-level group"}
db.Create(&root)
```

**Create a child group:**
```go
child := Group{
    Title: "Child",
    Icon: "child.png",
    Description: "A subgroup",
    ParentID: &root.ID,
}
db.Create(&child)
```

**Query a group with its children:**
```go
var group Group
db.Preload("Children").First(&group, root.ID)
fmt.Println(group.Children)
```

**Query a group with its parent:**
```go
var child Group
db.Preload("Parent").First(&child, childID)
fmt.Println(child.Parent) // Parent group
```

### API Endpoints (Example)

- `GET    /api/groups` — List all groups
- `POST   /api/groups` — Create a new group
- `GET    /api/groups/:id` — Get group details (optionally with children)
- `PUT    /api/groups/:id` — Update group
- `DELETE /api/groups/:id` — Delete group

### Notes

- You can build arbitrarily deep hierarchies (trees) of groups.
- For advanced queries (e.g., all descendants), you may need recursive logic in Go.

---

**Summary:**  
The `Group` resource supports parent-child relationships, enabling hierarchical organization. This is implemented in GORM with a self-referential foreign key (`ParentID`). The API and client can use this to build and navigate group trees.

Let me know if you want a sample handler, migration, or more details! 

## Service Resource: Model

### Fields

- **ID**: Unique identifier (auto-incremented)
- **Title**: Name/title of the service
- **Description**: (string) Description of the service

### GORM Model Example

```go
import "gorm.io/gorm"

type Service struct {
    ID          uint   `gorm:"primaryKey"`
    Title       string `gorm:"not null"`
    Description string
    gorm.Model  // includes CreatedAt, UpdatedAt, DeletedAt
}
```

### Example Usage

**Create a service:**
```go
svc := Service{
    Title: "Web Server",
    Description: "Handles HTTP requests for the application.",
}
db.Create(&svc)
```

**Query a service:**
```go
var svc Service
db.First(&svc, 1) // Find service with ID=1
```

### API Endpoints (Example)

- `GET    /api/services` — List all services
- `POST   /api/services` — Create a new service
- `GET    /api/services/:id` — Get service details
- `PUT    /api/services/:id` — Update service
- `DELETE /api/services/:id` — Delete service

---

**Summary:**  
The `Service` resource includes `ID`, `Title`, and `Description` fields, and is managed via GORM for all database operations. This structure is suitable for representing any logical or physical service in your system.

Let me know if you want a sample handler, migration, or OpenAPI schema! 

## Tag Resource: Model

### Fields

- **ID**: Unique identifier (auto-incremented)
- **Title**: Name/title of the tag
- **Description**: (string) Description of the tag
- **Values**: (string or JSON/text) Possible values for the tag (see note below)
- **Query**: (string) Query or filter expression associated with the tag
- **Weight**: (int) Weight or priority for sorting/ranking

### GORM Model Example

```go
import "gorm.io/gorm"

type Tag struct {
    ID          uint   `gorm:"primaryKey"`
    Title       string `gorm:"not null"`
    Description string
    Values      string // Could be a comma-separated string, JSON, or use a custom type
    Query       string
    Weight      int
    gorm.Model  // includes CreatedAt, UpdatedAt, DeletedAt
}
```

#### Note on `Values` Field
- If you want to store multiple values, you can use:
  - A comma-separated string (e.g., `"red,green,blue"`)
  - JSON-encoded array (e.g., `'["red","green","blue"]'`)
  - Or define a custom GORM type for more advanced use

### Example Usage

**Create a tag:**
```go
tag := Tag{
    Title: "Environment",
    Description: "Deployment environment",
    Values: "dev,staging,prod",
    Query: "env=prod",
    Weight: 10,
}
db.Create(&tag)
```

**Query tags by weight:**
```go
var tags []Tag
db.Order("weight desc").Find(&tags)
```

### API Endpoints (Example)

- `GET    /api/tags` — List all tags
- `POST   /api/tags` — Create a new tag
- `GET    /api/tags/:id` — Get tag details
- `PUT    /api/tags/:id` — Update tag
- `DELETE /api/tags/:id` — Delete tag

---

**Summary:**  
The `Tag` resource includes `ID`, `Title`, `Description`, `Values`, `Query`, and `Weight`. It is managed via GORM and can be used for categorization, filtering, or metadata in your system.

Let me know if you want a more advanced example for the `Values` field or a sample handler! 

## Field Resource: Model

### Fields

- **ID**: Unique identifier (auto-incremented)
- **Type**: (string) The type of the field (e.g., "string", "int", "date", etc.)
- **Description**: (string) Description of the field
- **Weight**: (int) Weight or priority for ordering

### GORM Model Example

```go
import "gorm.io/gorm"

type Field struct {
    ID          uint   `gorm:"primaryKey"`
    Type        string `gorm:"not null"`
    Description string
    Weight      int
    gorm.Model  // includes CreatedAt, UpdatedAt, DeletedAt
}
```

### Example Usage

**Create a field:**
```go
field := Field{
    Type:        "string",
    Description: "A generic string field",
    Weight:      1,
}
db.Create(&field)
```

**Query fields ordered by weight:**
```go
var fields []Field
db.Order("weight asc").Find(&fields)
```

### API Endpoints (Example)

- `GET    /api/fields` — List all fields
- `POST   /api/fields` — Create a new field
- `GET    /api/fields/:id` — Get field details
- `PUT    /api/fields/:id` — Update field
- `DELETE /api/fields/:id` — Delete field

---

**Summary:**  
The `Field` resource includes `ID`, `Type`, `Description`, and `Weight`. It is managed via GORM and can be used to define dynamic or custom fields in your system.

Let me know if you want a sample migration, handler, or OpenAPI schema! 

## Variable Resource: Model

### Fields

- **ID**: Unique identifier (auto-incremented)
- **Key**: (string) The variable's name or key (should be unique)
- **Value**: (string) The variable's value

### GORM Model Example

```go
import "gorm.io/gorm"

type Variable struct {
    ID    uint   `gorm:"primaryKey"`
    Key   string `gorm:"uniqueIndex;not null"`
    Value string
    gorm.Model  // includes CreatedAt, UpdatedAt, DeletedAt
}
```

### Example Usage

**Create a variable:**
```go
v := Variable{
    Key:   "APP_ENV",
    Value: "production",
}
db.Create(&v)
```

**Query a variable by key:**
```go
var v Variable
db.Where("key = ?", "APP_ENV").First(&v)
```

### API Endpoints (Example)

- `GET    /api/variables` — List all variables
- `POST   /api/variables` — Create a new variable
- `GET    /api/variables/:id` — Get variable details
- `PUT    /api/variables/:id` — Update variable
- `DELETE /api/variables/:id` — Delete variable

---

**Summary:**  
The `Variable` resource includes `ID`, `Key`, and `Value`. It is managed via GORM and is ideal for storing configuration, environment variables, or other key-value pairs in your system.

Let me know if you want a sample handler, migration, or OpenAPI schema! 