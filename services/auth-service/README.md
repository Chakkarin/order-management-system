# Authen-Service

Authen-Service is a microservice designed to handle authentication and user management functionalities such as registration, login, and password recovery. It is built using **Golang** with a focus on Clean Architecture principles, ensuring scalability, maintainability, and testability.

---

## Features

- **User Registration**: Create new user accounts.
- **User Login**: Authenticate users with email and password.
- **Forgot Password**: Allow users to reset their password via email.
- **JWT Authentication**: Generate and validate JWT tokens for secure communication.

---

## Technologies Used

### Backend
- **Language**: Golang
- **Framework**: Gin
- **Database**: PostgreSQL (via GORM)
- **Cache**: Redis
- **Authentication**: JWT
- **Architecture**: Clean Architecture

### DevOps
- **Containerization**: Docker
- **Orchestration**: Docker Compose

---

## Project Structure

```plaintext
auth-service/
├── cmd/
│   └── main.go                # Entry point of the application
├── internal/
│   ├── domain/                # Business logic (Entities, Interfaces)
│   ├── usecases/              # Application logic (Use Cases)
│   ├── controllers/           # Interface adapters (HTTP handlers)
│   ├── repositories/          # Database access implementations
│   └── infrastructure/        # Frameworks and external systems
├── migrations/                # Database migrations
├── proto/                     # gRPC Protocol Files
├── Dockerfile                 # Docker build file
├── docker-compose.yml         # Compose configuration
├── go.mod                     # Dependency management
└── README.md                  # Project documentation
```
---

## Setup Instructions

### Prerequisites

- Install **Docker** and **Docker Compose**.

## Steps

1. Clone the repository:
```bash
git clone https://github.com/Chakkarin/order-management-system/auth-service.git
cd auth-service
```

2. Run the services using Docker Compose:
```bash
docker-compose up --build
```

3. Access the service:
- API: http://localhost:8080

---

## API Endpoints

### Authentication
| Method | Endpoint                | Description                 |
|--------|-------------------------|-----------------------------|
| POST   | `/auth/register`        | Register a new user         |
| POST   | `/auth/login`           | Login with email/password   |
| POST   | `/auth/forgot-password` | Request password reset      |

---

## Example API Requests

#### Register a New User
```bash
curl -X POST http://localhost:8080/auth/register \
-H "Content-Type: application/json" \
-d '{"username": "john_doe", "email": "john@example.com", "password": "password123"}'
```

#### Login
```bash
curl -X POST http://localhost:8080/auth/login \
-H "Content-Type: application/json" \
-d '{"email": "john@example.com", "password": "password123"}'
```

---

## Environment Variables

Make sure to set the following environment variables in a .env file or in Docker Compose:

```bash
DATABASE_URL=postgresql://auth_user:auth_password@auth-db:5432/auth_db
REDIS_URL=redis://redis:6379
JWT_SECRET=your_jwt_secret
```

---

## Testing
Use Postman or Curl to test the APIs. You can also write Unit Tests for the Use Cases and Controllers using Go's testing package.

### Run Unit Tests
```bash
go test ./...
```

---

## Future Enhancements

- Implement multi-factor authentication (MFA).
- Add role-based access control (RBAC).
- Enhance logging and monitoring.
- Add support for token refresh and account deactivation.