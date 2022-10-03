package pkg

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

var (
	ctx context.Context
)

// Function that sets up all settings to database
func DatabaseMySql() *DB_Settings {
	database := new(DB_Settings)
	CheckDBSettings(database)
	return database
}

// Function to retry connection to database if it fails
func RetryConnection(databaseType string, databaseURL string) (error, *sql.DB) {
	connection, err := sql.Open(databaseType, databaseURL)
	return err, connection
}

// Function that connects to database using based url
func ConnectToDatabase(databaseType string, databaseURL string) *sql.DB {
	connection, err := sql.Open(databaseType, databaseURL)
	if err != nil {
		log.Println(err)
		for i := 0; i < 5; i++ {
			log.Printf("Retrying connection to database for the %v time", i)
			err, connection = RetryConnection(databaseType, databaseURL)
			if err == nil {
				break
			}
			time.Sleep(5 * time.Second)
		}

	}
	return connection
}

// Function that checks if the connection to database is set
func CheckConnectionToDatabase(connection *sql.DB) {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	status := "up"
	err := connection.PingContext(ctx)
	if err != nil {
		status = "down"
	}
	log.Printf("Database connection is %s\n %v", status, err)
}

// Function that closes existing connection to database
func CloseConnectionToDatabase(connection *sql.DB) {
	err := connection.Close()
	if err != nil {
		fmt.Println(err)
	}
}
