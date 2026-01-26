# Ayam Bangkok 2.0 Public Service API

[![Go Version](https://img.shields.io/badge/go-1.25+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

Public API service for the Ayam Bangkok 2.0, built with Go using an isolated feature-based architecture, structured logging, database integration MySQL.

## 🚀 Features

- **Feature-Based Architecture** - Each endpoint is a self-contained, isolated feature.
- **DB Support** - Support for MySQL & PostgresSQL using GORM.
- **Configuration Management** - Environment variables.
- **Health Monitoring** - Health checks for database and application.

## 📁 Project Structure

This project uses Go (Gin) with a Feature-Based + Clean Architecture approach so that the code is easy to read, scalable, and easy to develop by the team.

```
├── migrations/
│   └── main.go                 # Database migration (create/update table)
│
├── source/
│   ├── common/
│   │   └── glob_utils/         # Global utilities (hashing, helper, dll)
│   │
│   ├── health/                 # Health check endpoint
│   │
│   └── models/                 # Database models (ORM mapping)
│
├── config/
│   └── config.go               # Load & manage environment configuration
│
├── features/                   # Feature-based modules (1 endpoint = 1 feature)
│   └── exampletocopy/          # Template for creating new features
│       ├── handler.go          # Interface handler (HTTP layer)
│       ├── handler_impl.go     # Implementation handler
│       ├── usecase.go          # Interface business logic
│       ├── usecase_impl.go     # Implementation business logic
│       ├── repository.go       # Interface data access
│       └── repository_impl.go  # Implementation database access
│
├── pkg/
│   ├── db/                     # Database connection & configuration
│   └── logger/                 # Logger (zlog)
│
├── services/
│   └── route.go                # HTTP route registration
│
├── middleware/
│   ├── authMiddleware.go       # Authentication / authorization middleware
|   |── adminMiddleware.go      # Admin / authorization middleware
│   └── loggerRoute.go          # Request logging middleware
│
├── .env.example                # Example environment variable
├── main.go                     # Application entry point
└── version                     # Version app
```

## ⚡ Quick Start

### Installation

1. **Clone the repository**

   ```bash
   git clone https://github.com/LeoneLee1/ayam-bangkok-2.0-backend.git
   cd ayam-bangkok-2.0
   ```

2. **Install dependencies**

   ```bash
   go mod download
   ```

3. **Configure database**
   Copy `env_example` and adjust your database connections.

4. **Run database migration**

   ```bash
   go run migrations/main.go
   ```

5. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080/api` (default)

## 🏥 Health Checks

The application includes comprehensive health monitoring:

### Endpoints

- **`GET /health`** - Database connectivity and application status
