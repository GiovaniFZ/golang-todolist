# Todo List API

A REST API for managing tasks, built with Go. Create an account, authenticate, and manage personal tasks with CRUD (create, read, update, and delete) operations.

## Features

- User registration and JWT authentication
- Create, list, update, and delete tasks
- Local persistence with SQLite and GORM
- Automated controller tests

## Architecture

The project follows an MVC-style architecture, separating responsibilities across models (data structures), controllers (HTTP request handling and application flow), and views (JSON response formatting). Repositories isolate data access from the controllers.

## Tech Stack

**Go** · **Gin** · **GORM** · **SQLite** · **JWT**

## Getting Started

### Prerequisites

- Go 1.27 or later

### Setup

1. Clone the repository and navigate to the project directory.
2. Create a `.env` file in the project root with a secret used to sign tokens:

   ```env
   JWT_SECRET=replace-with-a-strong-secret
   ```

3. Download dependencies and start the API:

   ```bash
   go mod download
   go run .
   ```

The API will be available at `http://localhost:8080`. The SQLite database (`database.db`) and its tables are created automatically on the first run.

## Endpoints

| Method | Route | Access | Description |
| --- | --- | --- | --- |
| `POST` | `/sign-up` | Public | Register a user |
| `POST` | `/login` | Public | Authenticate and receive an access token |
| `GET` | `/tasks` | Authenticated | List the user's tasks |
| `POST` | `/tasks` | Authenticated | Create a task |
| `PATCH` | `/tasks/:id` | Authenticated | Update a task |
| `DELETE` | `/tasks/:id` | Authenticated | Delete a task |

Task endpoints require a JWT in the `Authorization` header, using the `Bearer <token>` format. Access tokens expire after 15 minutes.

## Tests

Run the controller tests with:

```bash
go test github/giovanifz/todolist-go/controllers/tests
```

## Request Examples

[`requests.http`](./requests.http) includes examples for registration, login, and task operations. Run them with an HTTP client extension, such as the VS Code REST Client.
