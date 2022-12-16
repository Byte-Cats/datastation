package pkg

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"time"
)

// MySqlDBConfig is a struct that contains the configuration for the database connection.
type MySqlDBConfig struct {
	Info            *ConnectionConfig
	DBName          string
	Timeout         time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// DefaultMySqlConf returns a struct with default values for the fields.
func DefaultMySqlConf() *MySqlDBConfig {
	newInfo := DefaultMySql()
	return &MySqlDBConfig{
		Info:            &newInfo,
		DBName:          "database",
		Timeout:         5 * time.Second,
		MaxIdleConns:    2,
		MaxOpenConns:    10,
		ConnMaxLifetime: 30 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
	}
}

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

// ExecuteSQL is a generic function that can be used to execute any SQL query.
func ExecuteSQL(db *sql.DB, query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args...)
}

// QuerySQL is a generic function that can be used to execute any SELECT SQL query.
// It returns a slice of map[string]interface{}, with each map representing a row in the result set.
// The keys of the map are the column names, and the values are the column values for that row.
func QuerySQL(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	result := make([]map[string]interface{}, 0)
	for rows.Next() {
		columns, err := rows.Columns()
		if err != nil {
			return nil, err
		}

		values := make([]interface{}, len(columns))
		valuePtrs := make([]interface{}, len(columns))
		for i := range values {
			valuePtrs[i] = &values[i]
		}

		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, err
		}

		row := make(map[string]interface{})
		for i, column := range columns {
			row[column] = values[i]
		}
		result = append(result, row)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateMySqlTable creates a new table in the database based on the provided struct.
func CreateMySqlTable(db *sql.DB, tableName string, s interface{}) error {
	// Get the type and value of the struct
	val := reflect.ValueOf(s)
	typ := reflect.TypeOf(s)

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
		val = val.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return fmt.Errorf("CreateTable: s must be a struct or a pointer to a struct")
	}

	// Build the column definitions
	columnDefs := make([]string, 0)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		name := field.Tag.Get("db")
		if name == "" {
			name = field.Name
		}

		// Get the column type from the field type
		var colType string
		switch field.Type.Kind() {
		case reflect.Bool:
			colType = "BOOLEAN"
		case reflect.Int:
			colType = "INT"
		case reflect.Int8:
			colType = "TINYINT"
		case reflect.Int16:
			colType = "SMALLINT"
		case reflect.Int32:
			colType = "INT"
		case reflect.Int64:
			colType = "BIGINT"
		case reflect.Uint:
			colType = "INT UNSIGNED"
		case reflect.Uint8:
			colType = "TINYINT UNSIGNED"
		case reflect.Uint16:
			colType = "SMALLINT UNSIGNED"
		case reflect.Uint32:
			colType = "INT UNSIGNED"
		case reflect.Uint64:
			colType = "BIGINT UNSIGNED"
		case reflect.Float32:
			colType = "FLOAT"
		case reflect.Float64:
			colType = "DOUBLE"
		case reflect.String:
			colType = "VARCHAR(255)"
		default:
			colType = "VARCHAR(255)"
		}

		// Add any additional column options specified in the field tag
		options := field.Tag.Get("options")
		if options != "" {
			colType += " " + options
		}

		columnDefs = append(columnDefs, fmt.Sprintf("%s %s", name, colType))
	}

	// Build the CREATE TABLE statement
	query := fmt.Sprintf("CREATE TABLE %s (%s)", tableName, strings.Join(columnDefs, ", "))
	// Execute the query
	_, err := db.Exec(query)
	return err
}
