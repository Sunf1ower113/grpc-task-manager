APP_NAME = grpc-task-manager
DB_SERVICE = db
APP_SERVICE = app
DOCKER_COMPOSE = docker-compose

build:
	@echo "Building the Go application..."
	docker build -t $(APP_NAME) .

up:
	@echo "Starting the application with Docker Compose..."
	$(DOCKER_COMPOSE) up --build -d

down:
	@echo "Stopping the application..."
	$(DOCKER_COMPOSE) down

migrate:
	@echo "Running database migrations..."
	$(DOCKER_COMPOSE) exec $(APP_SERVICE) go run ./migrations

logs:
	@echo "Viewing logs..."
	$(DOCKER_COMPOSE) logs -f $(APP_SERVICE)

dev:
	@echo "Starting the application in development mode..."
	$(DOCKER_COMPOSE) up --build

prod:
	@echo "Starting the application in production mode..."
	$(DOCKER_COMPOSE) up --build -d

test:
	@echo "Running tests..."
	$(DOCKER_COMPOSE) exec $(APP_SERVICE) go test ./...

clean:
	@echo "Cleaning up Docker images and containers..."
	$(DOCKER_COMPOSE) down --volumes --rmi all

rebuild: down build up
	@echo "Rebuilding the application..."
