package grpc_test_test

import (
	"context"
	"database/sql"
	"github.com/Sunf1ower113/grpc-task-manager/internal/composites"
	"github.com/Sunf1ower113/grpc-task-manager/internal/config"
	"github.com/Sunf1ower113/grpc-task-manager/pkg/client/postgres"
	pb "github.com/Sunf1ower113/grpc-task-manager/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"testing"
	"time"
)

func cleanupDatabase(db *sql.DB, t *testing.T) {
	query := `DELETE FROM tasks`
	_, err := db.Exec(query)
	if err != nil {
		t.Fatalf("Failed to clean up database: %v", err)
	}
}

func TestCreateTask(t *testing.T) {
	appConfig, err := config.InitConfig()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}

	logger, err := config.InitLogger(appConfig.Logger)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	database, err := postgres.NewDB(
		appConfig.DBUser,
		appConfig.DBPassword,
		appConfig.DBName,
		appConfig.DBHost,
		appConfig.DBPort,
		logger,
	)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	cleanupDatabase(database, t)

	taskComposite, err := composites.NewTaskComposite(database, logger)
	if err != nil {
		t.Fatalf("Failed to initialize task composite: %v", err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	pb.RegisterTaskManagerServer(server, taskComposite.Handler)

	address := appConfig.GRPCHost + ":" + appConfig.GRPCPort
	go func() {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			t.Fatalf("Failed to start gRPC listener: %v", err)
		}
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	time.Sleep(2 * time.Second)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create gRPC client connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       "Test Task",
		Description: "This is a test task",
	})
	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	if resp.Title != "Test Task" && resp.Description != "This is a test task" {
		t.Errorf("CreateTask returned unexpected response: got %v want %v", resp.Title, "Test Task")
	}
	log.Printf("Task created successfully: %+v", resp)

	cleanupDatabase(database, t)
}

func TestListTasks(t *testing.T) {
	appConfig, err := config.InitConfig()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}

	logger, err := config.InitLogger(appConfig.Logger)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	database, err := postgres.NewDB(
		appConfig.DBUser,
		appConfig.DBPassword,
		appConfig.DBName,
		appConfig.DBHost,
		appConfig.DBPort,
		logger,
	)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	cleanupDatabase(database, t)

	taskComposite, err := composites.NewTaskComposite(database, logger)
	if err != nil {
		t.Fatalf("Failed to initialize task composite: %v", err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	pb.RegisterTaskManagerServer(server, taskComposite.Handler)

	address := appConfig.GRPCHost + ":" + appConfig.GRPCPort
	go func() {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			t.Fatalf("Failed to start gRPC listener: %v", err)
		}
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	time.Sleep(2 * time.Second)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create gRPC client connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       "Task 1",
		Description: "First test task",
	})
	if err != nil {
		t.Fatalf("Failed to create first task: %v", err)
	}

	_, err = client.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       "Task 2",
		Description: "Second test task",
	})
	if err != nil {
		t.Fatalf("Failed to create second task: %v", err)
	}

	resp, err := client.ListTasks(ctx, &pb.ListTasksRequest{})
	if err != nil {
		t.Fatalf("Failed to list tasks: %v", err)
	}

	if len(resp.Tasks) != 2 {
		t.Errorf("Expected 2 tasks, got %d", len(resp.Tasks))
	}

	for _, task := range resp.Tasks {
		if task.Id <= 0 || task.Title == "" || task.Description == "" {
			t.Errorf("Invalid task data: %+v", task)
		}
	}

	log.Printf("Tasks listed successfully: %+v", resp.Tasks)

	cleanupDatabase(database, t)
}

func TestGetTask(t *testing.T) {
	appConfig, err := config.InitConfig()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}

	logger, err := config.InitLogger(appConfig.Logger)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	database, err := postgres.NewDB(
		appConfig.DBUser,
		appConfig.DBPassword,
		appConfig.DBName,
		appConfig.DBHost,
		appConfig.DBPort,
		logger,
	)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	cleanupDatabase(database, t)

	taskComposite, err := composites.NewTaskComposite(database, logger)
	if err != nil {
		t.Fatalf("Failed to initialize task composite: %v", err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	pb.RegisterTaskManagerServer(server, taskComposite.Handler)

	address := appConfig.GRPCHost + ":" + appConfig.GRPCPort
	go func() {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			t.Fatalf("Failed to start gRPC listener: %v", err)
		}
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	time.Sleep(2 * time.Second)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create gRPC client connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createResp, err := client.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       "Task to Get",
		Description: "Task description",
	})
	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	getResp, err := client.GetTask(ctx, &pb.GetTaskRequest{
		Id: createResp.Id,
	})
	if err != nil {
		t.Fatalf("Failed to get task: %v", err)
	}

	if getResp.Id != createResp.Id || getResp.Title != "Task to Get" || getResp.Description != "Task description" {
		t.Errorf("Expected task data to match, got %+v", getResp)
	}

	cleanupDatabase(database, t)
}

func TestUpdateTask(t *testing.T) {
	appConfig, err := config.InitConfig()
	if err != nil {
		t.Fatalf("Failed to initialize configuration: %v", err)
	}

	logger, err := config.InitLogger(appConfig.Logger)
	if err != nil {
		t.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	database, err := postgres.NewDB(
		appConfig.DBUser,
		appConfig.DBPassword,
		appConfig.DBName,
		appConfig.DBHost,
		appConfig.DBPort,
		logger,
	)
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.Close()

	cleanupDatabase(database, t)

	taskComposite, err := composites.NewTaskComposite(database, logger)
	if err != nil {
		t.Fatalf("Failed to initialize task composite: %v", err)
	}

	server := grpc.NewServer()
	defer server.Stop()

	pb.RegisterTaskManagerServer(server, taskComposite.Handler)

	address := appConfig.GRPCHost + ":" + appConfig.GRPCPort
	go func() {
		listener, err := net.Listen("tcp", address)
		if err != nil {
			t.Fatalf("Failed to start gRPC listener: %v", err)
		}
		if err := server.Serve(listener); err != nil {
			t.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()
	time.Sleep(2 * time.Second)

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("Failed to create gRPC client connection: %v", err)
	}
	defer conn.Close()

	client := pb.NewTaskManagerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	createResp, err := client.CreateTask(ctx, &pb.CreateTaskRequest{
		Title:       "Original Title",
		Description: "Original Description",
	})
	if err != nil {
		t.Fatalf("Failed to create task: %v", err)
	}

	updateResp, err := client.UpdateTask(ctx, &pb.UpdateTaskRequest{
		Id:          createResp.Id,
		Title:       "Updated Title",
		Description: "Updated Description",
	})
	if err != nil {
		t.Fatalf("Failed to update task: %v", err)
	}

	if updateResp.Title != "Updated Title" || updateResp.Description != "Updated Description" {
		t.Errorf("Task was not updated correctly: %+v", updateResp)
	}

	cleanupDatabase(database, t)
}
