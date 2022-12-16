package postgres

import (
	"database/sql"
	app "datastation/pkg/app"
)

// Function to create a new struct to hold postgres settings
func DatabasePostgres() *app.DB_Settings {
	postgres := new(app.DB_Settings)
	app.CheckDBSettings(postgres)
	return postgres
}

// Function to connect to Postgres database with provided settings
func ConnectToPostgres(postgres *app.DB_Settings) *sql.DB {
	postgresDSN := postgres.Type + "://" + postgres.User + ":" + postgres.Password + "@" + postgres.Hostname + "/" + postgres.Name + "?sslmode=disable"
	conn := app.ConnectToDatabase(postgres.Type, postgresDSN)
	return conn
}
