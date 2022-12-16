package pkg

import (
	"database/sql"
	"fmt"
)

func buildConnectionString(config ConnectionConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.Hostname, config.Port, config.Username, config.Password, config.Name)
}

func openDatabase(connString string) (*sql.DB, error) {
	return sql.Open("postgres", connString)
}

func pingDatabase(db *sql.DB) error {
	return db.Ping()
}

func SetupDatabase(config ConnectionConfig) (*sql.DB, error) {
	connString := buildConnectionString(config)
	db, err := openDatabase(connString)
	if err != nil {
		return nil, err
	}

	err = pingDatabase(db)
	if err != nil {
		return nil, err
	}

	return db, nil
}
