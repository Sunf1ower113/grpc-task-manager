APP_NAME = grpc-task-manager
DOCKER_COMPOSE = docker-compose
PROD_APP_SERVICE = grpc-task-manager
PROD_DB_SERVICE = grpc-task-manager-prod-db

build:
	@echo "Building the Go application..."
	docker build -t $(APP_NAME) .

up:
	@echo "Starting the application in production mode..."
	$(DOCKER_COMPOSE) up --build -d

down:
	@echo "Stopping the application..."
	$(DOCKER_COMPOSE) down

migrate:
	@echo "Running database migrations..."
	$(DOCKER_COMPOSE) exec $(PROD_DB_SERVICE) sh -c "psql -U $$POSTGRES_USER -d $$POSTGRES_DB -f /docker-entrypoint-initdb.d/0001_create_tasks_table.sql"

logs:
	@echo "Viewing logs..."
	$(DOCKER_COMPOSE) logs -f $(PROD_APP_SERVICE)

clean:
	@echo "Cleaning up Docker images and containers..."
	$(DOCKER_COMPOSE) down --volumes --rmi all

rebuild: clean build up
	@echo "Rebuilding the application..."
