package mysql

import (
	"database/sql"
	"fmt"
	"strconv"
)


type MySQLHandler struct {
	conn *sql.DB
}


func NewMySQLHandler(conn *sql.DB) MySQLHandler {
	return MySQLHandler{
				conn: conn,
	}
}


func (connection MySQLHandler) AddToMySqlTable(tableName string) error {
	query := `
	CREATE TABLE users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY (id)
	);`
	return connection.Exec(query)
}

func (connection MySQLHandler) Exec(query string, args ...interface{}) error {
	_, err := connection.conn.Exec(query, args...)
	return err
}

func (connection MySQLHandler) MySqlDeleteItem(tableName string, itemID int) error {
	query := "DELETE FROM " + tableName + " WHERE id=" + strconv.Itoa(itemID) + ";"
	return connection.Exec(query)
}

func (connection MySQLHandler) MySqlFindItemByID(tableName string, itemId int) error {
	query := "SELECT * FROM " + tableName + " WHERE ID = id LIMIT " + strconv.Itoa(itemId) + ";"
	return connection.Exec(query)
}