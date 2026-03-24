# Layered Inventory Management API

## Overview

A production-ready **REST API built with Gin (Go)**, designed using a layered architecture for scalability, maintainability, and clear separation of concerns.

This project serves as a boilerplate for building structured Go APIs while gaining hands-on experience with modern backend practices.

## Third Party Integrations

- **Gin** (*v1.12.0*) - Web Framework
- **modernc.org/sqlite** (*v1.46.1*) - SQLite Driver

## Folder Structure

- `cmd/`
  - `api/`
    - `main.go` - *Application entry point*
- `data/`
  - `app.db` - *Database file*
- `internal/`
  - `config/`
    - `db.go` - *Database Configuration & Initialization*
  - `handler/` - *HTTP Handlers*
    - `product_handler.go`
  - `model/` - *Domain Models*
    - `user.go`
  - `repository/` - *Database Access Layer*
    - `product_repository.go`
  - `route/` - *Routing Definitions*
    - `routes.go`
  - `service/` - *Dependency Wiring*
    - `service.go`
- `go.mod`
- `go.sum`

## Design Principles

- **Layered Architecture** – Clear separation between handler, service, and repository layers
- **Dependency Injection** – Promotes testability and loose coupling
- **Repository Pattern** – Abstracts database operations
- **Input Validation & Error Handling** – Consistent and secure API behavior
