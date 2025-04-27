# Task Dependency Manager

A REST API built with Go and the Gin framework to manage tasks with dependencies.

## Setup

1. Clone the repository:
   ```bash
   git clone https://github.com/Soulevoker/task-dependency-manager.git
   cd task-dependency-manager
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the API:
   ```bash
   go run cmd/api/main.go
   ```

   The API runs on `http://localhost:8080` by default. Set the `PORT` environment variable to change the port:
   ```bash
   PORT=9090 go run cmd/api/main.go
   ```

## Endpoints

- `GET /health`: Returns `{"status": "OK"}` (status 200).
- `GET /version`: Returns `{"version": "1.0.0"}` (status 200).