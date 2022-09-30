package main

import (
	"datastation/applogic"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Putting this here for now to hold the mysql driver dependency in go.mod
	mysql := applogic.DatabaseMySql()
	conn := applogic.ConnectToDatabase(mysql)
	applogic.CheckConnectionToDatabase(conn)
	applogic.CloseConnectionToDatabase(conn)
	fmt.Println(conn)
}
