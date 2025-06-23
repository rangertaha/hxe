# CHANGELOG

All notable changes to this project will be documented in this file.

## [Unreleased] - 2025-01-27

### Added
- **JWT Authentication System**
  - JWT middleware for Echo v4 with authentication endpoints
  - Default credentials: username `admin`, password `password`
  - Token expiration and automatic refresh capabilities
  - CORS support for cross-origin requests

- **Enhanced HXE Client Library**
  - Complete Go client library with JWT authentication support
  - Beautiful table formatting using `go-pretty/v6` library
  - Multi-operation support (MultiStart, MultiStop, MultiDelete, etc.)
  - Cross-platform user/group detection

- **Program Management Features**
  - Full CRUD operations and runtime operations (Start, Stop, Restart, etc.)
  - Advanced operations: Shell, Tail, Run
  - Status monitoring and schema-driven table display

- **API Server & Database**
  - RESTful API endpoints with JWT protection
  - SQLite database with GORM integration
  - Service layer architecture with proper error handling

- **Documentation**
  - Comprehensive README for client library
  - API reference with code examples

### Changed
- **API Structure**: Refactored handlers and service methods for better type safety
- **Client Interface**: Standardized method signatures and enhanced table formatting
- **Authentication Flow**: Simplified client creation with automatic token management

### Fixed
- **Type Safety**: Fixed return value mismatches and function signatures
- **API Endpoints**: Corrected HTTP methods and parameter handling
- **Client Methods**: Fixed HTTP client integration and response parsing
- **Database**: Resolved GORM relationships and migration issues

### Technical Improvements
- **Code Quality**: Enhanced error handling and code organization
- **Performance**: Optimized database queries and HTTP client operations
- **Security**: Implemented proper JWT validation and CORS configuration

---

## [0.1.0] - 2025-06-20
### Added
- Initial project structure
- Basic program management functionality
- Database models and migrations
- API server foundation

---
