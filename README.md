# gRPC Task Manager

Project for managing tasks using gRPC and PostgreSQL.

## Structure

- **cmd** — entry point and server startup.
- **internal** — internal logic and services.
- **proto** — gRPC service definitions.
- **migrations** — database migrations.
- **Dockerfile** — Docker image definition.
- **docker-compose.yml** — service definitions and interactions.
- **Makefile** — build and run commands.

## Running the project

1. Copy `.env.example` to `.env` and set database connection parameters.
2. Run the project:
   ```bash
   make up
