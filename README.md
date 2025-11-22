# Event Registration REST API (Go + Gin)

A simple, clean, and fully-functional **REST API built using Go, Gin, SQLite, and JWT authentication**.
This project was created with the core purpose of **learning Go backend development, structuring a real API, implementing authentication, and handling common CRUD operations**.

---

## Features

### Authentication

-   User signup (password hashing using bcrypt)
-   User login (JWT token generation)
-   Middleware-protected routes

### Event Management

-   Create event (protected)
-   Update event (protected)
-   Delete event (protected)
-   Fetch all events (public)
-   Fetch a single event (public)

### Event Registration

-   Authenticated users can register for events
-   Prevent registering for non-existing events
-   SQLite database with:

    -   Users table
    -   Events table
    -   Event Registrations table

-   Foreign key constraints

### Security

-   JWT authentication middleware
-   Password hashing (bcrypt)
-   CORS configured for frontend integration

---

## Tech Stack

| Layer            | Technology              |
| ---------------- | ----------------------- |
| Language         | **Go (Golang)**         |
| Framework        | **Gin Web Framework**   |
| Database         | **SQLite**              |
| Auth             | **JWT (golang-jwt v5)** |
| Password hashing | **bcrypt**              |
| Live Reload      | **Air**                 |
| Env config       | `godotenv`              |

---

## Setup & Installation

### 1 Clone the repository

```bash
git clone https://github.com/TaushifReza/go-event-booking.git .
cd Backend
```

### 2 Install dependencies

```bash
go mod tidy
```

### 3 Create `.env` file

```
FRONTEND_URL=http://localhost:5173
```

### 4 Run the server

```bash
go run main.go
```

API runs on:

```
http://localhost:8080
```
