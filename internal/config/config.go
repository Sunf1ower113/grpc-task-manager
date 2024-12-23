package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoggerConfig defines the configuration for the application's logging.
type LoggerConfig struct {
	Level       string
	OutputPaths []string
}

// AppConfig defines the application's configuration settings.
type AppConfig struct {
	DBUser          string
	DBPassword      string
	DBName          string
	DBHost          string
	DBPort          string
	GRPCHost        string
	GRPCPort        string
	Logger          *LoggerConfig
	GeneratedPath   string
	DebugMode       bool
	EnableProfiling bool
}

// InitConfig initializes the application configuration by reading environment variables.
// It attempts to load variables from a .env file if not running in Docker.
func InitConfig() (*AppConfig, error) {
	if os.Getenv("ENV_MODE") != "docker" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: .env file not found or could not be loaded")
		}
	}

	return &AppConfig{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		GRPCHost:   os.Getenv("GRPC_HOST"),
		GRPCPort:   os.Getenv("GRPC_PORT"),
		Logger: &LoggerConfig{
			Level:       os.Getenv("LOG_LEVEL"),
			OutputPaths: []string{"stdout", os.Getenv("LOG_FILE_PATH")},
		},
		GeneratedPath:   os.Getenv("GENERATED_PATH"),
		DebugMode:       os.Getenv("DEBUG_MODE") == "true",
		EnableProfiling: os.Getenv("ENABLE_PROFILING") == "true",
	}, nil
}
