package main

import (
	app "datastation/pkg/app"
	"datastation/pkg/mysql"
	"fmt"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Putting this here for now to hold the mysql driver dependency in go.mod
	mysqly := mysql.DatabaseMySql()
	conn := mysql.ConnectToMysql(mysqly)
	mysqlconn := mysql.NewMysqly(conn)
	mysqlconn.AddToMySqlTable("NewTable")
	app.CheckConnectionToDatabase(conn)
	app.CloseConnectionToDatabase(conn)
	fmt.Println(conn)
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
