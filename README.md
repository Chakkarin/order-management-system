# Order Management System

**Order Management System** is a full-stack microservices-based application for managing user authentication, orders, inventory, and notifications. This project focuses on practicing Full Stack Development skills.

---

## Features

### Core Features
1. **User Authentication**:
   - Secure user registration, login, and password recovery using JWT.
2. **Order Management**:
   - Create, update, and track orders.
3. **Inventory Management**:
   - Real-time product stock updates and low inventory alerts.
4. **Notifications**:
   - Event-driven notifications for order updates.
5. **Customer Management**:
   - Manage user profiles and shipping addresses.

### Future Enhancements
1. **Role-Based Access Control (RBAC)** for advanced user management.
2. **WebSocket Integration** for real-time notifications.
3. **Payment Gateway Integration** for handling online payments.
4. **Monitoring with AWS X-Ray or Prometheus/Grafana**.

---

## Tech Stack

### Backend
- **Language**: Golang
- **Frameworks**: Gin (REST), gRPC
- **Database**: PostgreSQL, Redis
- **Messaging**: Kafka
- **Authentication**: JWT
- **ORM**: GORM

### Frontend
- **Framework**: Next.js
- **Styling**: Tailwind CSS
- **State Management**: React Query
- **Form Validation**: React Hook Form + Yup

### DevOps
- **Containerization**: Docker, Docker Compose
- **Cloud**: AWS (EC2, S3, RDS)
- **CI/CD**: GitHub Actions

---

## Project Structure

```plaintext
project-root/
├── api-gateway/
├── services/
│   ├── auth-service/             # User Authentication service
│   ├── order-service/            # Order Management service
│   ├── inventory-service/        # Inventory Management service
│   ├── notification-service/     # Notification service
│   └── customer-service/         # Customer Management service
├── frontend/                     # Frontend application
├── docker-compose.yml            # Compose file for all services
└── README.md                     # This file
```
---

## Setup Instructions

### Prerequisites

1. Install **Docker** and **Docker Compose**.
1. Install **Node.js** and **npm** (or Yarn) for the frontend.

### Steps

1. Clone the Repository
```plaintext
git clone https://github.com/your-username/order-management-system.git
cd order-management-system
```

2. Start Backend Services
```plaintext
docker-compose up --build
```

3. Start Frontend
```plaintext
cd frontend
npm install
npm run dev
```
Access the Frontend at `http://localhost:3000`.

4. Access Microservices via API Gateway

- Auth-Service: `http://localhost:8080`
- Order-Service: `http://localhost:8081`
- Inventory-Service: `http://localhost:8082`
- Notification-Service: `http://localhost:8083`
- Customer-Service: `http://localhost:8084`

---

## API Endpoints (via API Gateway)
| Method	|Endpoint	            |Description               |
|-----------|-----------------------|--------------------------|
| POST	    |/auth/register	        |Register a new user       |
| POST	    |/auth/login	        |Login with credentials    |
| POST	    |/auth/forgot-password	|Request password reset    |
| GET	    |/orders	            |List all orders           |
| POST	    |/orders	            |Create a new order        |
| GET	    |/inventory	            |View product stock levels |
| PUT	    |/inventory/:id	        |Update stock for a product|

---

## Future Enhancements

- **RBAC**: Enable fine-grained access control for different roles (Admin, Customer).
- **Real-Time Notifications**: Implement WebSocket for instant updates.
- **Payment Integration**: Add support for payment gateways like Stripe or PayPal.
- **Advanced Monitoring**: Use AWS X-Ray, Prometheus, or Grafana for detailed tracing and visualization.