package main

import (
	"github.com/Sunf1ower113/grpc-task-manager/internal/composites"
	"github.com/Sunf1ower113/grpc-task-manager/internal/config"
	"github.com/Sunf1ower113/grpc-task-manager/pkg/client/postgres"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	appConfig, err := config.InitConfig()
	if err != nil {
		log.Fatalf("Failed to initialize configuration: %v", err)
	}

	logger, err := config.InitLogger(appConfig.Logger)
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
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
		logger.Fatal("Failed to initialize database", zap.Error(err))
	}
	defer database.Close()

	taskComposite, err := composites.NewTaskComposite(database, logger)

	if err != nil {
		logger.Fatal("Failed to initialize task composite", zap.Error(err))
	}

	start(taskComposite, appConfig, logger)
}

func start(taskComposite *composites.TaskComposite, cfg *config.AppConfig, logger *zap.Logger) {
	logger.Info("Starting the gRPC server...")

	server := grpc.NewServer()
	address := net.JoinHostPort(cfg.GRPCHost, cfg.GRPCPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		logger.Fatal("Failed to start listener", zap.Error(err))
	}

	logger.Info("gRPC server is listening", zap.String("address", address))

	if err := server.Serve(listener); err != nil {
		logger.Fatal("Failed to start gRPC server", zap.Error(err))
	}
}
