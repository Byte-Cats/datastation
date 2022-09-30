package applogic

import (
	"database/sql"
	"fmt"
)

// Function that sets up all settings to database
func DatabaseMySql() *DB_Settings {
	database := new(DB_Settings)
	CheckDBSettings(database)
	return database
}

// Function that connects to database using based url
func ConnectToDatabase(database *DB_Settings) *sql.DB {
	databaseUrl := database.User + database.Password + "@(" + database.Hostname + ")/" + database.Name + "?parserTime=true"
	connection, err := sql.Open(database.Type, databaseUrl)
	if err != nil {
		fmt.Println(err)
	}
	return connection
}

// Function that checks if the connection to database is set
func CheckConnectionToDatabase(connection *sql.DB) {
	// Will change it to PingContext() in future
	err := connection.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

// Function that closes existing connection to database
func CloseConnectionToDatabase(connection *sql.DB) {
	err := connection.Close()
	if err != nil {
		fmt.Println(err)
	}
}
