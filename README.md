# GitHub Info Distributed System

A distributed microservices system built with **Golang (1.25+)** to fetch and display GitHub repository information. This project is a production-ready architecture following **Clean Architecture** principles.

## 🏗 System Architecture

The system consists of two decoupled microservices interacting via high-performance **gRPC**:

1.  **Collector Service** (Internal):
    *   **Role**: Internal data provider encapsulating GitHub API interactions.
    *   **Transport**: gRPC Server (Port `50051`).
    *   **Architecture**: Structured into `adapter` (GitHub client), `delivery` (gRPC controller), `domain`, and `usecase`. Includes custom data formatting (e.g., parsing RFC3339 dates).

2.  **API Gateway** (Public):
    *   **Role**: Public entry point for users/clients.
    *   **Transport**: REST Server (HTTP Port `8080`) & gRPC Client.
    *   **Logic**: Receives external HTTP requests, forwards them to the Collector via gRPC, and returns the response in JSON format.

### Interaction Flow:
`User -> [HTTP GET] -> API Gateway -> [gRPC] -> Collector -> [HTTPS] -> GitHub API`

---

## ✨ Key Features
*   **Clean Architecture**: Strict separation of concerns (`Domain` -> `UseCase` -> `Adapter`/`Delivery`) in both services.
*   **Graceful Shutdown**: Both HTTP and gRPC servers correctly handle `SIGINT`/`SIGTERM` signals with configurable timeouts.
*   **Intelligent Error Mapping**: Translates GitHub HTTP errors to gRPC status codes (e.g., `codes.NotFound`), which the Gateway maps back to standard HTTP codes (`404 Not Found`).
*   **Automated Documentation**: Integrated Swagger UI for testing REST endpoints.

---

## 📂 Project Structure

```text
.
├── api/
│   ├── proto/          # gRPC contracts (.proto and generated .pb.go files)
│   └── swagger/        # Auto-generated OpenAPI/Swagger documentation
├── collector/          # Collector Microservice
│   ├── cmd/            # Entry point (main.go)
│   └── internal/       # Core logic (adapter, config, delivery, domain, usecase)
├── gateway/            # API Gateway Microservice
│   ├── cmd/            # Entry point (main.go)
│   └── internal/       # Core logic (adapter, config, delivery, domain, usecase)
├── docker-compose.yaml # Orchestration for one-command startup
├── go.mod              # Dependencies (Go 1.25.7)
└── README.md
```

---

## 🚀 Getting Started

### 1. Run with Docker (Recommended)
You can start the entire system with a single command. Ensure Docker and Docker Compose are installed:

```bash
docker-compose up --build
```

The services will be available at:
*   **API Gateway**: `http://localhost:8080`
*   **Swagger UI**: `http://localhost:8080/swagger/index.html`

### 2. Run Locally (Manual)
First, start the Collector service:
```bash
go run collector/cmd/main.go
```

Then, in a separate terminal, start the API Gateway:
```bash
go run gateway/cmd/main.go
```

---

## 📖 API Documentation (Swagger)

The system provides a built-in Swagger interface for easy testing. 

1.  Open `http://localhost:8080/swagger/index.html` in your browser.
2.  Use the `GET /repo` endpoint.
3.  **Query Parameters**: 
    *   `owner` (e.g., `google`) 
    *   `repo` (e.g., `go`)

**Example Request**:
```bash
curl -X GET "http://localhost:8080/repo?owner=google&repo=go"
```

**Example Response**:
```json
{
  "name": "google/go",
  "description": "The Go programming language",
  "stars": 120500,
  "forks": 17000,
  "created_at": "15:04:05 02.01.06"
}
```

### Error Handling
The Gateway translates internal gRPC errors into standard HTTP codes:
*   `200 OK`: Success.
*   `400 Bad Request`: Missing `owner` or `repo` query parameters.
*   `404 Not Found`: Repository does not exist on GitHub.
*   `500 Internal Server Error`: Connection issues or API rate limits.

---

## 🛠 Development Commands

If you change the `.proto` files or the Swagger annotations, run the following commands to regenerate the code:

*   **Generate gRPC Code**:
    ```bash
    protoc --go_out=. --go-grpc_out=. api/proto/github.proto
    ```
*   **Update Swagger Docs**:
    ```bash
    swag init -g gateway/cmd/main.go -o api/swagger
    ```