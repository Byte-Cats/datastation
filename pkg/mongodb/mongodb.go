package mongodb

import (
	"database/sql"
	app "datastation/pkg/app"
)

// Function to set settings to mongodb database
func DatabaseMongoDB() *app.DB_Settings {
	mongodb := new(app.DB_Settings)
	app.CheckDBSettings(mongodb)
	return mongodb
}

// Function to connect to mongoDB database with provided settings
func ConnectToMongoDB(mongodb *app.DB_Settings) *sql.DB {
	mongodbDSN := mongodb.User + mongodb.Password + "@(" + mongodb.Hostname + ")/" + mongodb.Name + "?parseTime=true"
	conn := app.ConnectToDatabase(mongodb.Type, mongodbDSN)
	return conn
}
