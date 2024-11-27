APP_NAME = grpc-task-manager
DOCKER_COMPOSE = docker-compose
PROD_APP_SERVICE = grpc-task-manager
PROD_DB_SERVICE = grpc-task-manager-prod-db
PROFILE = default

build:
	@echo "Building the Go application..."
	docker build -t $(APP_NAME) .

up:
	@echo "Starting the application in production mode..."
	$(DOCKER_COMPOSE) --profile $(PROFILE) up --build -d

clean:
	@echo "Cleaning up Docker images and containers..."
	$(DOCKER_COMPOSE) --profile $(PROFILE) down --volumes --rmi all

logs:
	@echo "Viewing logs..."
	$(DOCKER_COMPOSE) --profile $(PROFILE) logs -f $(PROD_APP_SERVICE)

rebuild: clean build up
	@echo "Rebuilding the application..."

generate-protobuf:
	@echo "Generating Protobuf files locally..."
	protoc --go_out=./proto --go-grpc_out=./proto \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative proto/task.proto
