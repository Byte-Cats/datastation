package pkg

import (
	"database/sql"
	"fmt"
)

// WithDBConnection establishes a connection to the database, creates the specified table if it doesn't exist,
// and then closes the connection when the provided function is finished executing.
func WithDBConnection(config *MySqlDBConfig, tableName string, s interface{}, f func(*sql.DB) error) error {
	db, err := NewDBConnection(config)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := CreateMySqlTable(db, tableName, s); err != nil {
		return err
	}

	return f(db)
}

// NewDBConnection establishes a connection to the MySQL database specified in the provided ConnectionConfig.
func NewDBConnection(config *MySqlDBConfig) (*sql.DB, error) {
	// Build the connection string
	dsn := MySqlString(config)

	db, err := MySqlConn(config.Info.Type, dsn)
	if err != nil {
		return nil, err
	}

	// Set the connection options
	ApplyConfig(db, config)

	// Test the connection
	if err := MySqlPingDB(db); err != nil {
		return nil, err
	}
	return db, nil
}

// ApplyConfig sets the connection pool settings for a *sql.DB instance using the provided ConnectionConfig.
func ApplyConfig(db *sql.DB, config *MySqlDBConfig) {
	db.SetMaxIdleConns(config.MaxIdleConns)
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetConnMaxLifetime(config.ConnMaxLifetime)
	db.SetConnMaxIdleTime(config.ConnMaxIdleTime)
}

// MySqlString builds connection string for connecting to a MySQL database using the provided ConnectionConfig.
func MySqlString(config *MySqlDBConfig) string {
	// Build the base connection string
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Info.Username, config.Info.Password, config.Info.Hostname, config.Info.Port, config.DBName)
	// Build the query parameters
	queryParams := fmt.Sprintf("timeout=%s&maxIdleConns=%d&maxOpenConns=%d&connMaxLifetime=%s&connMaxIdleTime=%s",
		config.Timeout, config.MaxIdleConns, config.MaxOpenConns, config.ConnMaxLifetime, config.ConnMaxIdleTime)
	// Return the full connection string
	return fmt.Sprintf("%s?%s", connString, queryParams)
}

// MySqlConn establishes a connection to the database.
func MySqlConn(dbType string, dsn string) (*sql.DB, error) {
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %w", err)
	}
	return db, nil
}

// MySqlPingDB pings the database to check the connection.
func MySqlPingDB(db *sql.DB) error {
	err := db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}
	return nil
}
