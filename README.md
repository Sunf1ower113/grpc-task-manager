# gRPC Task Manager

Task Manager is a gRPC-based application for managing tasks. It supports operations to create, list, retrieve, update, and delete tasks, backed by a PostgreSQL database.

---

## Project Structure

This project follows the principles of Clean Architecture to ensure maintainability and scalability.

- **cmd**: Entry point for starting the application, including the gRPC server.
- **internal**: Contains the core application logic, divided into multiple layers:
   - **adapters**: Adapters for infrastructure-level integrations such as gRPC handlers and database repositories.
      - **grpc**: gRPC handlers for interacting with the application logic.
      - **db**: Database repository implementations (e.g., PostgreSQL).
   - **composites**: Dependency injection and composition of services, repositories, and handlers.
   - **config**: Configuration loading and initialization.
   - **domain**: Core business logic and domain models.
      - **models**: Business entities such as `Task`.
      - **services**: Core business logic operations implemented as services.
      - **repository**: Interfaces defining the contracts for database operations.
- **proto**: Contains gRPC service definitions (`.proto` files) and generated code.
- **pkg**: Common utilities and helpers (e.g., database clients).
- **Dockerfile**: Defines the Docker image for the application.
- **docker-compose.yml**: Configuration for running the application and dependent services like PostgreSQL in containers.
- **Makefile**: Build, deployment, and utility commands for automation.

---

## Running the Project

### 1. Running Locally

#### Prerequisites
- Install `Go` (minimum version 1.22).
- Install `PostgreSQL` and ensure it is running.
- Install `protoc` for generating gRPC files.

#### Steps
1. configure the necessary environment variables in `.env`.
2. Generate gRPC protobuf files:  
   ```
   protoc --go_out=./proto --go-grpc_out=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/task.proto
   ```
3. Run tests:
   ```
   go test ./... -v
   ```
4. Run the application:  
   ```
   go run ./cmd/server/main.go
   ```
### 2. Running with Makefile

#### Prerequisites
- Install `Make`.

#### Steps
1. configure the necessary environment variables in `.env.prod` .
2. Run the application:  
   ```
   make up
   ```
3. View logs:  
   ```
   make logs
   ```
4. Clean up Docker containers and images:  
   ```
   make clean
   ```
5. Rebuild the application:  
   ```
   make rebuild
   ```
6. Generate gRPC protobuf files:  
   ```
   make generate-protobuf
   ```

### 3. Running with Docker (without Makefile)

#### Prerequisites
- Install `Docker`.
- Install `Docker Compose`.

#### Steps
1. configure the necessary environment variables in `.env.prod` .
2. Build and start the application:  
   ```
   docker build -t <name-of-container> .
   docker-compose --profile default up --build -d
   ```
3. View logs:  
   ```
   docker logs -f <name-of-container>
   ```
4. Clean up Docker containers and images:  
   ```
   docker-compose down --volumes --rmi all
   ```

### Example gRPC Tests with `grpcurl`

#### Prerequisites
- Ensure the application is running.
- `grpcurl` must be installed or accessible through Docker.

#### Tests

1. **CreateTask (Valid Request)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"title":"Valid Task","description":"A proper description"}" localhost:50051 taskmanager.TaskManager/CreateTask
   ```
2. **CreateTask (Invalid Request - Empty Title)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"title":"","description":"Description without title"}" localhost:50051 taskmanager.TaskManager/CreateTask 
   ```
3. **ListTasks**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{}" localhost:50051 taskmanager.TaskManager/ListTasks
   ```
4. **GetTask (Valid Request)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"id":1}" localhost:50051 taskmanager.TaskManager/GetTask
   ```
5. **GetTask (Invalid Request - Nonexistent ID)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"id":9999}" localhost:50051 taskmanager.TaskManager/GetTask
   ```
6. **UpdateTask (Valid Request)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"id":1,"title":"Updated Task","description":"Updated description"}" localhost:50051 taskmanager.TaskManager/UpdateTask
   ```
7. **UpdateTask (Invalid Request - Empty Title and Description)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"id":1,"title":"","description":""}" localhost:50051 taskmanager.TaskManager/UpdateTask
   ```
8. **DeleteTask (Valid Request)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"id":1}" localhost:50051 taskmanager.TaskManager/DeleteTask
   ```
9. **DeleteTask (Invalid Request - Nonexistent ID)**  
   ```
   docker run --rm --network="host" fullstorydev/grpcurl -plaintext -d "{"id":9999}" localhost:50051 taskmanager.TaskManager/DeleteTask
   ```
