package pkg

import (
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
