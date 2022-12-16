package pkg

import (
	"time"
)

// ConnectionConfig is a struct that contains the configuration for the database connection.
type ConnectionConfig struct {
	Type            string
	Host            string
	Port            int
	Username        string
	Password        string
	Database        string
	Timeout         time.Duration
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime time.Duration
	ConnMaxIdleTime time.Duration
}

// DefaultConnectionConfig returns a ConnectionConfig struct with default values for the fields.
func DefaultConnectionConfig() ConnectionConfig {
	return ConnectionConfig{
		Type:            "mysql",
		Host:            "localhost",
		Port:            3306,
		Username:        "user",
		Password:        "password",
		Database:        "database",
		Timeout:         5 * time.Second,
		MaxIdleConns:    2,
		MaxOpenConns:    10,
		ConnMaxLifetime: 30 * time.Minute,
		ConnMaxIdleTime: 10 * time.Minute,
	}
}
