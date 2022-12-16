package pkg

import (
	"database/sql"
	"fmt"
)

// WithDBConnection establishes a connection to the database, creates the specified table if it doesn't exist,
// and then closes the connection when the provided function is finished executing.
func WithDBConnection(config ConnectionConfig, tableName string, s interface{}, f func(*sql.DB) error) error {
	db, err := NewDBConnection(config)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := CreateTable(db, tableName, s); err != nil {
		return err
	}

	return f(db)
}

// NewDBConnection establishes a connection to the MySQL database specified in the provided ConnectionConfig.
func NewDBConnection(config ConnectionConfig) (*sql.DB, error) {
	// Build the connection string
	dsn := BuildConnectionString(config)

	db, err := ConnectToDatabase(config.Type, dsn)
	if err != nil {
		return nil, err
	}

	// Set the connection options
	ApplyConfig(db, config)

	// Test the connection
	if err := PingDB(db); err != nil {
		return nil, err
	}
	return db, nil
}

// ApplyConfig sets the connection pool settings for a *sql.DB instance using the provided ConnectionConfig.
func ApplyConfig(db *sql.DB, config ConnectionConfig) {
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
}

// BuildConnectionString builds connection string for connecting to a MySQL database using the provided ConnectionConfig.
func BuildConnectionString(config ConnectionConfig) string {
	// Build the base connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	// Build the query parameters
	queryParams := fmt.Sprintf("timeout=%s&maxIdleConns=%d&maxOpenConns=%d&connMaxLifetime=%s&connMaxIdleTime=%s",
		config.Timeout, config.MaxIdleConns, config.MaxOpenConns, config.ConnMaxLifetime, config.ConnMaxIdleTime)
	// Return the full connection string
	return fmt.Sprintf("%s?%s", connString, queryParams)
}

// ConnectToDatabase establishes a connection to the database.
func ConnectToDatabase(dbType string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return db, nil
}

// PingDB pings the database to check the connection.
func PingDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}
	return nil
}
