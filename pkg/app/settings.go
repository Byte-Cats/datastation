package pkg

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Struct of database settings
type DB_Settings struct {
	// Database's type
	Type string
	// Database's protocol
	Protocol string
	// Database's hostname
	Hostname string
	// Database's port
	Port string
	// Database's name
	Name string
	// Database's username
	User string
	// Database's password
	Password string
}

// Struct to store default database's settings
type Template struct {
	Type     string
	Protocol string
	Hostname string
	Port     string
	Name     string
	User     string
	Password string
}

// I guess this function is temporary, just to describe the
// principle of how stuff works here
func Default_Template() Template {
	return Template{
		// Defaul Database type
		Type: "Dunno",
		// Default Protocol
		Protocol: "TCP",
		// Default Hostname (There is no place like 127.0.0.1)
		Hostname: "127.0.0.1",
		// Default Port
		Port: ":8080",
		// Default database's name
		Name: "Datastore",
		// Default database username
		User: "admin",
		// Default, but the strongest database password
		Password: "1234",
	}
}

// Function that checks if database's type is set from .env
func CheckType(database *DB_Settings, defaults string) {
	database.Type = os.Getenv("DATABASE_TYPE")
	if database.Type == "" {
		database.Type = defaults
	}
}

func CheckProtocol(database *DB_Settings, defaults string) {
	database.Protocol = os.Getenv("DATABASE_PROTOCOL")
	if database.Protocol == "" {
		database.Protocol = defaults
	}
}

// Function that checks if database's hostname is set from .env
func CheckHostname(database *DB_Settings, defaults string) {
	database.Hostname = os.Getenv("DATABASE_HOST")
	if database.Hostname == "" {
		database.Hostname = defaults
	}
}

// Function that checks if database's port is set from .env
func CheckPort(database *DB_Settings, defaults string) {
	database.Port = os.Getenv("DATABASE_PORT")
	if database.Port == "" {
		database.Port = defaults
	}
}

// Function that checks if database's name is set from .env
func CheckName(database *DB_Settings, defaults string) {
	database.Name = os.Getenv("DATABASE_NAME")
	if database.Name == "" {
		database.Name = defaults
	}
}

// Function that checks if database's user is set from .env
func CheckUser(database *DB_Settings, defaults string) {
	database.User = os.Getenv("DATABASE_NAME")
	if database.User == "" {
		database.User = defaults
	}
}

// Function that checks if database's password is set from .env
func CheckPassword(database *DB_Settings, defaults string) {
	database.Password = os.Getenv("DATABASE_PASSWORD")
	if database.Password == "" {
		database.Password = defaults
	}
}

// Function that contains all CheckFunctions and invokes full check
func CheckDBSettings(database *DB_Settings) {
	CheckType(database, Default_Template().Type)
	CheckProtocol(database, Default_Template().Protocol)
	CheckHostname(database, Default_Template().Hostname)
	CheckPort(database, Default_Template().Port)
	CheckName(database, Default_Template().Name)
	CheckUser(database, Default_Template().User)
	CheckPassword(database, Default_Template().Password)
}

// Function that sets database's type value
func (database *DB_Settings) SetDatabaseType(dtype string) {
	database.Type = dtype
}

// Function that returns database's type value
func GetDatabaseType(database *DB_Settings) string {
	return database.Type
}

// Function that prints database's type value
func ShowDatabaseType(database *DB_Settings) {
	fmt.Println(database.Type)
}

// Function that sets database's protocol value
func (database *DB_Settings) SetDatabaseProtocol(dprotocol string) {
	database.Protocol = dprotocol
}

// Functions that returns database's protocol value
func GetDatabaseProtocol(database *DB_Settings) string {
	return database.Protocol
}

// Function that prints database's protocol value
func ShowDatabaseProtocol(database *DB_Settings) {
	fmt.Println(database.Protocol)
}

// Function that sets database's hostname value
func (database *DB_Settings) SetDatabaseHost(dhost string) {
	database.Hostname = dhost
}

// Function that returns database's hostname value
func GetDatabaseHost(database *DB_Settings) string {
	return database.Hostname
}

// Function that prints database's hostname value
func ShowDatabaseHost(database *DB_Settings) {
	fmt.Println(database.Hostname)
}

// Function that sets database's port value
func (database *DB_Settings) SetDatabasePort(dport string) {
	database.Port = dport
}

// Function that returns database's port value
func GetDatabasePort(database *DB_Settings) string {
	return database.Port
}

// Function that prints database's port value
func ShowDatabasePort(database *DB_Settings) {
	fmt.Println(database.Port)
}

// Function that sets database's name value
func (database *DB_Settings) SetDatabaseName(dname string) {
	database.Name = dname
}

// Function that returns database's name value
func GetDatabaseName(database *DB_Settings) string {
	return database.Name
}

// Function that prints database's name value
func ShowDatabaseName(database *DB_Settings) {
	fmt.Println(database.Name)
}

// Function that sets database's username value
func (database *DB_Settings) SetDatabaseUser(dusername string) {
	database.User = dusername
}

// Function that returns database's username value
func GetDatabaseUser(database *DB_Settings) string {
	return database.User
}

// Function that prints database's username value
func ShowDatabaseUser(database *DB_Settings) {
	fmt.Println(database.User)
}

// Function that sets database's password value
func (database *DB_Settings) SetDatabasePassword(dpassword string) {
	database.Password = dpassword
}

// Function that returns database's password value
func GetDatabasePassword(database *DB_Settings) string {
	return database.Password
}

// Function that prints database's password value
func ShowDatabasePassword(database *DB_Settings) {
	fmt.Println(database.Password)
}

// Set database's connection settings
func SetConnectionSettings(database *sql.DB) *sql.DB {
	database.SetConnMaxLifetime(time.Minute * 3)
	database.SetConnMaxIdleTime(time.Minute * 2)
	database.SetMaxOpenConns(12)
	database.SetMaxIdleConns(8)
	return database
}
