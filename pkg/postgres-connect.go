package pkg

import (
	"database/sql"
	"fmt"
)

// PostgresString returns a connection string for a Postgres database.
func PostgresString(config DatabaseConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Hostname, config.Port, config.Username, config.Password, config.Name)
}

// OpenPostgres opens a new connection to a Postgres database using the provided connection string.
func OpenPostgres(connString string) (*sql.DB, error) {
	return sql.Open("postgres", connString)
}

// PingPostgres pings the Postgres database to test the connection.
func PingPostgres(db *sql.DB) error {
	return db.Ping()
}

// SetupPostgres sets up a connection to a Postgres database using the provided configuration.
func SetupPostgres(config DatabaseConfig) (*sql.DB, error) {
	connString := PostgresString(config)
	db, err := OpenPostgres(connString)
	if err != nil {
		return nil, err
	}

	err = PingPostgres(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// ClosePostgres closes the connection to the Postgres database.
func ClosePostgres(db *sql.DB) error {
	return db.Close()
}

// ExecPostgres executes a query that does not return rows (e.g. INSERT, UPDATE, DELETE).
func ExecPostgres(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

// QueryPostgres executes a query that returns rows (e.g. SELECT).
func QueryPostgres(db *sql.DB, query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args...)
}

// QueryRowPostgres executes a query that returns a single row.
// It is useful when you are expecting a single row as the result of a query (e.g. SELECT COUNT(*) FROM table).
func QueryRowPostgres(db *sql.DB, query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args...)
}

// BeginPostgres starts a new transaction.
// Transactions are useful for ensuring that a series of queries are executed as a single unit of work,
// and can be rolled back if any of the queries fail.
func BeginPostgres(db *sql.DB) (*sql.Tx, error) {
	return db.Begin()
}
