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

## Endpoints with Descriptions

### Base Endpoints
- `GET /health`
  - **Description:** Returns the health status of the API to ensure it is operational.
  - **Example Response:**
    ```json
    {
      "status": "OK"
    }
    ```

- `GET /version`
  - **Description:** Returns the current version of the API.
  - **Example Response:**
    ```json
    {
      "version": "1.0.0"
    }
    ```

### Task Endpoints
- `GET /tasks/:id`
  - **Description:** Retrieves a specific task by its ID, including its details and dependencies.
  - **Example Response:**
    ```json
    {
      "id": "1",
      "name": "Design database schema",
      "description": "Create the initial database schema for the application.",
      "dependencies": ["2"]
    }
    ```

- `POST /tasks`
  - **Description:** Creates a new task with optional dependencies.
  - **Example Request:**
    ```json
    {
      "id": "3",
      "name": "Write API documentation",
      "description": "Document all the API endpoints for developers."
    }
    ```
  - **Example Response:**
    ```json
    {
      "id": "3",
      "name": "Write API documentation",
      "description": "Document all the API endpoints for developers.",
      "dependencies": []
    }
    ```

- `PUT /tasks/:id`
  - **Description:** Updates the details of an existing task by its ID.
  - **Example Request:**
    ```json
    {
      "name": "Update database schema",
      "description": "Make changes to the database schema to support new features."
    }
    ```
  - **Example Response:**
    ```json
    {
      "id": "1",
      "name": "Update database schema",
      "description": "Make changes to the database schema to support new features.",
      "dependencies": ["2"]
    }
    ```

- `DELETE /tasks/:id`
  - **Description:** Deletes a task by its ID and removes all its dependencies.
  - **Example Response:**
    ```json
    {
      "message": "task {1} deleted"
    }
    ```

- `GET /tasks`
  - **Description:** Lists all tasks along with their details and dependencies.
  - **Example Response:**
    ```json
    [
      {
        "id": "1",
        "name": "Design database schema",
        "description": "Create the initial database schema for the application.",
        "dependencies": ["2"]
      },
      {
        "id": "2",
        "name": "Set up development environment",
        "description": "Prepare the development environment with necessary tools and configurations.",
        "dependencies": []
      }
    ]
    ```

- `POST /tasks/:id/dependencies`
  - **Description:** Adds a dependency to an existing task.
  - **Example Request:**
    ```json
    {
      "dependency_id": "2"
    }
    ```
  - **Example Response:** No content (status 204).

- `DELETE /tasks/:id/dependencies/:depId`
  - **Description:** Removes a specific dependency from a task.
  - **Example Response:** No content (status 204).
