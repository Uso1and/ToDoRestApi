# ToDo REST API

REST API service for task management (ToDo), written in Go using Gin framework and PostgreSQL.

## 📌 Features

- Create, read, update and delete tasks (CRUD)
- API documentation with Swagger
- Fully tested handlers and repositories
- Ready-to-use Docker image with multi-stage build
- GitHub Actions integration for CI/CD
- HTML templates support

## 🛠 Technologies

- **Go 1.21+** - main programming language
- **Gin** - HTTP web framework
- **PostgreSQL** - database
- **Docker** - containerization
- **Swagger** - API documentation
- **GitHub Actions** - CI/CD

## 🚀 Quick Start

### Prerequisites

- Installed Go (version 1.21 or higher)
- PostgreSQL (version 15 or higher)
- Docker (optional)

### Local Setup

1. Clone the repository:
   git clone https://github.com/yourusername/ToDoRestApi.git
   cd ToDoRestApi

2. Set up environment variables (or use default values):
    export POSTGRES_HOST=localhost
    export POSTGRES_PORT=5432
    export POSTGRES_USER=postgres
    export POSTGRES_PASSWORD=root
    export POSTGRES_DB=todo

3. Start PostgreSQL and create database:
    docker-compose up -d db

4. Run the application:
    go run cmd/main.go

5. Application will be available at: http://localhost:8080

### Run with Docker

docker-compose up -d --build

### 📚 API Documentation
After starting the application, API documentation is available at:
- **Swagger UI:** - http://localhost:8080/swagger/index.html
- **Swagger JSON:** - http://localhost:8080/swagger/doc.json

### 🧪 Testing
To run tests:

go test -v ./...
Tests use test_todo database (configured in GitHub Action

