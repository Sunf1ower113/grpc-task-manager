package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	DBUser        string
	DBPassword    string
	DBName        string
	DBHost        string
	DBPort        string
	GRPCPort      string
	LogLevel      string
	LogFilePath   string
	GeneratedPath string
}

var config *Config

func LoadConfig() error {
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("Error loading .env file")
	}

	config = &Config{
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBHost:        os.Getenv("DB_HOST"),
		DBPort:        os.Getenv("DB_PORT"),
		GRPCPort:      os.Getenv("GRPC_PORT"),
		LogLevel:      os.Getenv("LOG_LEVEL"),
		LogFilePath:   os.Getenv("LOG_FILE_PATH"),
		GeneratedPath: os.Getenv("GENERATED_PATH"),
	}

	return nil
}

func GetConfig() *Config {
	return config
}
