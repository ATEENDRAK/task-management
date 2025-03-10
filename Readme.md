Task Management Service

Problem Breakdown

The Task Management Service is a RESTful API designed to:

1) Create, fetch, update, and delete tasks.

2) Implement pagination for fetching tasks along with filtering by status

3) Follow Domain-Driven Design (DDD) principles.

4) Using PostgreSQL as the persistent storage.

Design Decisions

1) Microservices Architecture: Follow DDD pattern with api, domain application and infrastructure for the separation of concern.

2) Gin Framework : HTTP router for handling API requests.

3) GORM: ORM for database operations with PostgreSQL.

4) Scalability: Can interact with other microservices like User service using rest protocols

Running the Service

Prerequisites

Ensure you have the following installed:

1) Install golang

2) Install PostgreSQL 

3)  Configure the Database

    Open PostgreSQL CLI:

    psql -U postgres

    Create the database:

    CREATE DATABASE task_management;

    Exit PostgreSQL:

    \q

4️) Install Dependencies

go mod tidy

5️) Run the Service

go run main.go

If everything is correct, you should see:

Connected to PostgresSql database!
Listening and serving HTTP on :8080

📌 API Documentation

1) Create a Task

Endpoint: POST /tasks

curl --location 'http://localhost:8080/tasks' \
--header 'Content-Type: application/json' \
--data '{"title": "Learn PostgreSQL", "status": "Pending"}'

 Response:
{
    "id": 1,
    "title": "Learn PostgreSQL",
    "status": "Pending",
    "created_at": "2025-03-10T23:16:45.490923+05:30",
    "updated_at": "2025-03-10T23:16:45.490923+05:30"
}

2️) Get All Tasks (With Pagination & Filtering)

Endpoint: GET /tasks

curl --location 'http://localhost:8080/tasks?status=Completed&limit=5&offset=1'

📌 Response:

[
  {
    "id": 1,
    "title": "Learn PostgreSQL",
    "status": "Completed",
    "created_at": "2024-03-06T12:00:00Z",
    "updated_at": "2024-03-06T12:00:00Z"
  }
]

3️⃣ Update a Task

Endpoint: PUT /tasks/{id}

curl -X PUT "http://localhost:8080/tasks/1" \n     -H "Content-Type: application/json" \n     -d '{"title": "Learn Golang", "status": "Completed"}'

📌 Response:

{
  "message": "Task updated"
}

4️⃣ Delete a Task

Endpoint: DELETE /tasks/{id}

curl -X DELETE "http://localhost:8080/tasks/1"

📌 Response:

{
  "message": "Task deleted"
}

Microservices Concepts Demonstrated

Single Responsibility Principle: Handlers, services, and repositories follow separation of concerns.

Database Persistence Layer: Uses PostgreSQL with GORM.

RESTful API Design: Uses HTTP methods to perform CRUD operations.

Scalability: Can integrate with a User Service via REST/gRPC.


Future Improvements

Logging & Monitoring: Can be extended with Prometheus & Grafana.

Add JWT authentication for user access control.

Implement gRPC for inter-service communication.


