package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"go.uber.org/zap"

	_ "github.com/lib/pq" // PostgreSQL driver
)

// ConnectDB attempts to establish a connection to the PostgreSQL database.
// It retries up to 5 times with a 5-second interval between retries in case of failures.
func ConnectDB(dsn string) (*sql.DB, error) {
	var db *sql.DB
	var err error

	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
		}
		if err == nil {
			log.Println("Database connection established successfully")
			return db, nil
		}

		log.Printf("Failed to connect to database: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	return nil, fmt.Errorf("failed to connect to database after 5 attempts: %w", err)
}

// NewDB initializes a new PostgreSQL database connection with the given parameters.
// It also ensures the necessary tables are created.
func NewDB(user, password, dbName, host, port string, logger *zap.Logger) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbName, host, port)

	db, err := ConnectDB(connStr)
	if err != nil {
		logger.Error("Error creating database connection", zap.Error(err))
		return nil, err
	}

	err = createTables(db, logger)
	if err != nil {
		logger.Error("Error creating tables", zap.Error(err))
		return nil, err
	}

	return db, nil
}

// createTables ensures the required tables exist in the database.
func createTables(db *sql.DB, logger *zap.Logger) error {
	var queries []string

	tasks := `
CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
)
`
	queries = append(queries, tasks)

	for _, query := range queries {
		_, err := db.Exec(query)
		if err != nil {
			logger.Error("Failed to execute query", zap.String("query", query), zap.Error(err))
			return err
		}
	}

	logger.Info("Tables created successfully")
	return nil
}
