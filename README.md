# GitHub Info Distributed System

A distributed microservices system built with **Golang** to fetch and display GitHub repository information. This project is an evolution of a CLI tool into a production-ready architecture following **Clean Architecture** principles.

## 🏗 System Architecture

The system consists of two decoupled microservices interacting via high-performance **gRPC**:

1.  **Collector Service**:
    *   **Role**: Internal data provider.
    *   **Transport**: gRPC Server.
    *   **Logic**: Encapsulates GitHub API interaction.
    *   **Architecture**: Divided into layers: **Domain**, **Use Cases**, **Adapters** (GitHub client), and **Handlers** (gRPC server).

2.  **API Gateway**:
    *   **Role**: Public entry point ("The Waiter").
    *   **Transport**: REST Server (HTTP) & gRPC Client.
    *   **Logic**: Receives external HTTP requests, forwards them to the Collector via gRPC, and returns JSON.
    *   **Features**: Automated Swagger (OpenAPI) documentation and intelligent error mapping (e.g., gRPC `NotFound` -> HTTP `404`).

### Interaction Flow:
`User -> [REST/HTTP] -> API Gateway -> [gRPC] -> Collector -> [REST/HTTPS] -> GitHub API`

---

## 🛠 Tech Stack

*   **Go (Golang)**: Core language.
*   **gRPC & Protocol Buffers**: Internal service communication.
*   **Swagger (swag)**: Automated API documentation.
*   **Docker & Docker Compose**: Containerization and orchestration.
*   **Clean Architecture**: Separation of concerns for maintainability.

---

## 📂 Project Structure

```text
.
├── api/proto/          # gRPC contract definitions (.proto files)
├── cmd/                # Entry points (main.go) for each service
│   ├── collector/
│   └── gateway/
├── docs/               # Auto-generated Swagger documentation
├── internal/           # Private application code
│   ├── collector/      # Domain, UseCase, Adapter, Handler layers
│   └── gateway/        # Client and REST Handler layers
├── Dockerfile.collector
├── Dockerfile.gateway
├── docker-compose.yaml # Orchestration for one-command startup
└── go.mod              # Project dependencies
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
First, start the Collector:
```bash
go run cmd/collector/main.go
```

Then, in a separate terminal, start the Gateway:
```bash
go run cmd/gateway/main.go
```

---

## 📖 API Documentation

The system provides a built-in Swagger interface for easy testing. 

1.  Open `http://localhost:8080/swagger/index.html` in your browser.
2.  Use the `GET /repo` endpoint.
3.  Parameters: `owner` (e.g., `google`) and `repo` (e.g., `go`).

### Error Mapping
The Gateway automatically translates internal gRPC errors into standard HTTP codes:
*   `200 OK`: Success.
*   `400 Bad Request`: Missing query parameters.
*   `404 Not Found`: Repository does not exist on GitHub.
*   `500 Internal Server Error`: Connection issues or API rate limits.

---

## 🛠 Development Commands

*   **Generate gRPC Code**:
    ```bash
    protoc --go_out=. --go-grpc_out=. api/proto/github.proto
    ```
*   **Update Swagger Docs**:
    ```bash
    swag init -g cmd/gateway/main.go
    ```
