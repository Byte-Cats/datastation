package mysql

import (
	"database/sql"
	app "datastation/pkg/app"
	"fmt"
)

// Function to create a new struct to hold mysql settings
func DatabaseMySql() *app.DB_Settings {
	mysql := new(app.DB_Settings)
	app.CheckDBSettings(mysql)
	return mysql
}

// Function to connect to mysql database with provided settings
func ConnectToMysql(mysql *app.DB_Settings) *sql.DB {
	mysqlURL := mysql.User + mysql.Password + "@(" + mysql.Hostname + ")/" + mysql.Name + "?parseTime=true"
	conn := app.ConnectToDatabase(mysql.Type, mysqlURL)
	return conn
}

// I am not sure about the way I should create function like these
func Insert(connection *sql.DB) {
	insert := `insert into "Students"("Name", "Roll_Number") values('Jacob', 20)`
	_, err := connection.Exec(insert)
	if err != nil {
		fmt.Println(err)
	}
}
