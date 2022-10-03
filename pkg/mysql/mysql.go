package mysql

import (
	"database/sql"
	app "datastation/pkg/app"
	"fmt"
	"strconv"
)

type Mysqly struct {
	conn *sql.DB
}

// Function to create a new struct to hold mysql settings
func DatabaseMySql() *app.DB_Settings {
	mysql := new(app.DB_Settings)
	app.CheckDBSettings(mysql)
	return mysql
}

// Function to connect to mysql database with provided settings
func ConnectToMysql(mysql *app.DB_Settings) *sql.DB {
	mysqlDSN := mysql.User + ":" + mysql.Password + "@" + mysql.Protocol + "(" + mysql.Hostname + ")/" + mysql.Name + "?parseTime=true"
	conn := app.ConnectToDatabase(mysql.Type, mysqlDSN)
	return conn
}

// Function that applies basic congiguration for database connection
func SetupConnection(connection *sql.DB) *sql.DB {
	conn := app.SetConnectionSettings(connection)
	return conn
}

func NewMysqly(conn *sql.DB) Mysqly {
	return Mysqly{
		conn: conn,
	}
}

// Function to add a table
func (connection Mysqly) AddToMySqlTable(tableName string) string {
	query := `
    CREATE TABLE users (
        id INT AUTO_INCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        created_at DATETIME,
        PRIMARY KEY (id)
    );`
	return query
}

// Function to delete a row from table by it's ID
func (connection Mysqly) MySqlDeleteItem(tableName string, itemID int) sql.Result {
	query := "DELETE FROM " + tableName + "WHERE id=" + strconv.Itoa(itemID) + ";"
	result, err := connection.conn.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// Function to find a row from a table by it's ID
func (connection Mysqly) MySqlFindItemByID(tableName string, itemId int) sql.Result {
	query := "SELECT * FROM " + tableName + "WHERE ID = id LIMIT " + strconv.Itoa(itemId) + ";"
	result, err := connection.conn.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

// I am not sure about the way I should create functions like these
func Insert(connection *sql.DB) {
	insert := `insert into "Students"("Name", "Roll_Number") values('Jacob', 20)`
	_, err := connection.Exec(insert)
	if err != nil {
		fmt.Println(err)
	}
}
