package logger

import (
	"github.com/Sunf1ower113/grpc-task-manager/internal/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

var logInstance *logrus.Logger

func InitLogger() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logInstance = logrus.New()

	logLevel := config.GetConfig().LogLevel
	level, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.Fatal("Invalid log level")
	}
	logInstance.SetLevel(level)

	logFilePath := config.GetConfig().LogFilePath
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	logInstance.SetOutput(file)
	logInstance.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logInstance.Info("Logger initialized successfully")
}

func GetLogger() *logrus.Logger {
	return logInstance
}
