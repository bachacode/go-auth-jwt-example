# Go Authentication Example with JWT

This repository contains a simple Go API example for user authentication using JSON Web Tokens (JWT). It demonstrates a clean architecture with modular components, including handlers, middleware, and database integration.

## Features

- **User Signup**: Allows users to register with an email and password.
- **User Login**: Authenticates users and generates a JWT token.
- **Token Validation**: Validates JWT tokens for protected routes.
- **REST API**: Exposes endpoints for signup, login, and token validation.
- **Password Hashing**: Uses bcrypt for secure password storage.
- **Middleware**: Protects routes with JWT-based authentication.

## Project Structure

```
    go-auth-jwt-example/
    ├── cmd/
    │   └── api/
    │       └── api.go                # Entry point for the application
    ├── internal/
    │   ├── database/
    │   │   ├── database.go           # Database initialization and migration
    │   │   └── models.go             # Database models
    │   ├── handlers/
    │   │   └── users.go              # Handlers for user-related endpoints
    │   └── middleware/
    │       └── auth.go               # JWT authentication middleware
    └── go.mod                        # Go module definition
```

## Getting Started

### Prerequisites

- Go 1.24.1 or later installed on your system.
- SQLite installed on your system.
- A `.env` file with the following variables:
  ```
  DB_NAME=database.sqlite
  SECRET=your_jwt_secret
  ```

### Installation

1. Clone the repository:

```bash
git clone https://github.com/bachacode/go-auth-jwt-example.git
cd go-auth-jwt-example
```

2. Install dependencies:

```bash
go mod tidy
```

### Running the Service

To start the service, run the following command:

```bash
go run cmd/api/api.go
```

The service will start on port `8080`.

## API Documentation

### `POST /signup`
Registers a new user.

**Request Body**

| Field    | Type   | Description              |
|----------|--------|--------------------------|
| email    | string | The user's email address |
| password | string | The user's password      |

**Response**

| Field    | Type   | Description              |
|----------|--------|--------------------------|
| status   | int    | HTTP status code         |
| message  | string | Success or error message |

**Error responses**
- **400 Bad Request:** Invalid input data.
- **500 Internal Server Error:** Failed to create user.

---

### `POST /login`
Authenticates a user and generates a JWT token.

**Request Body**

| Field    | Type   | Description              |
|----------|--------|--------------------------|
| email    | string | The user's email address |
| password | string | The user's password      |

**Response**

| Field    | Type   | Description              |
|----------|--------|--------------------------|
| status   | int    | HTTP status code         |
| message  | string | Success or error message |

**Error responses**
- **400 Bad Request:** Invalid email or password.
- **500 Internal Server Error:** Failed to generate token.

---

### `GET /validate`
Validates the user's JWT token.

**Headers**

| Header           | Description              |
|-------------------|--------------------------|
| Authorization     | JWT token as a cookie   |

**Response**

| Field    | Type   | Description              |
|----------|--------|--------------------------|
| status   | int    | HTTP status code         |
| message  | string | Success or error message |

**Error responses**
- **401 Unauthorized:** Invalid or expired token.

---

### `GET /ping`
Health check endpoint.

**Response**

```json
{
    "message": "pong"
}
```

## Notes

- Ensure the `SECRET` environment variable is set to a strong, random value for secure token signing.
- The database file (`database.sqlite`) is ignored by Git as specified in `.gitignore`.
