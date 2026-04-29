# GoTasks

<p align="center">
  <img src="https://go.dev/images/go-logo-blue.svg" alt="Go Logo" width="200">
</p>

## Features

- User registration and login with JWT authentication
- CRUD operations for tasks
- Task completion status tracking
- Secure API with middleware protection
- MongoDB for persistent data storage

## Prerequisites

- Go 1.26.2 or later
- MongoDB running on localhost:27017

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/sriraghariharan/gotasks.git
   cd gotasks
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

## Usage

### Development

Use Air for live reloading during development:
```bash
air
```

The API will be available at `http://localhost:4000`.

## API Endpoints

### Authentication
- `POST /auth/signup` - Register a new user
- `POST /auth/login` - Authenticate user and return JWT token

### Tasks (JWT required)
- `POST /task/` - Create a new task
- `PUT /task/{id}` - Update an existing task
- `DELETE /task/{id}` - Delete a task
- `GET /task/all` - Retrieve all tasks for the authenticated user

## Project Structure

- `cmd/server/` - Application entry point
- `internal/db/` - Database connection and configuration
- `internal/handlers/` - HTTP request handlers
- `internal/middleware/` - Authentication middleware
- `internal/models/` - Data models
- `internal/routes/` - Route definitions
- `internal/services/` - Business logic
- `internal/validators/` - Input validation

---

*This project is a small step in embracing the Gopher spirit.*
