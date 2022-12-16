package pkg

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// ConnectionConfig stores the various settings for a database.
type ConnectionConfig struct {
	Type     string
	Protocol string
	Hostname string
	Port     string
	Name     string
	User     string
	Password string
}

// DefaultDataConfig returns a ConnectionConfig struct with default values for the fields.
func DefaultDataConfig() ConnectionConfig {
	return ConnectionConfig{
		Type:     "mysql",
		Protocol: "tcp",
		Hostname: "localhost",
		Port:     "3306",
		Name:     "database",
		User:     "user",
		Password: "password",
	}
}

// CheckType checks if the Type field in the given ConnectionConfig is set from the DATABASE_TYPE environment variable.
// If it is not set, it sets it to the given default value.
func CheckType(database *ConnectionConfig, defaults string) {
	database.Type = os.Getenv("DATABASE_TYPE")
	if database.Type == "" {
		database.Type = defaults
	}
}

// CheckProtocol checks if the Protocol field in the given ConnectionConfig is set from the DATABASE_PROTOCOL environment variable.
// If it is not set, it sets it to the given default value.
func CheckProtocol(database *ConnectionConfig, defaults string) {
	database.Protocol = os.Getenv("DATABASE_PROTOCOL")
	if database.Protocol == "" {
		database.Protocol = defaults
	}
}

// CheckHostname checks if the Hostname field in the given ConnectionConfig is set from the DATABASE_HOST environment variable.
// If it is not set, it sets it to the given default value.
func CheckHostname(database *ConnectionConfig, defaults string) {
	database.Hostname = os.Getenv("DATABASE_HOST")
	if database.Hostname == "" {
		database.Hostname = defaults
	}
}

// CheckPort checks if the Port field in the given ConnectionConfig is set from the DATABASE_PORT environment variable.
// If it is not set, it sets it to the given default value.
func CheckPort(database *ConnectionConfig, defaults string) {
	database.Port = os.Getenv("DATABASE_PORT")
	if database.Port == "" {
		database.Port = defaults
	}
}

// CheckName checks if the Name field in the given ConnectionConfig is set from the DATABASE_NAME environment variable.
// If it is not set, it sets it to the given default value.
func CheckName(database *ConnectionConfig, defaults string) {
	database.Name = os.Getenv("DATABASE_NAME")
	if database.Name == "" {
		database.Name = defaults
	}
}

// CheckUser checks if the User field in the given
// ConnectionConfig is set from the DATABASE_USER environment variable.
// If it is not set, it sets it to the given default value.
func CheckUser(database *ConnectionConfig, defaults string) {
	database.User = os.Getenv("DATABASE_USER")
	if database.User == "" {
		database.User = defaults
		return
	}
}

// CheckPassword checks if the Password field in the given ConnectionConfig is set from the DATABASE_PASSWORD environment variable.
// If it is not set, it sets it to the given default value.
func CheckPassword(database *ConnectionConfig, defaults string) {
	database.Password = os.Getenv("DATABASE_PASSWORD")
	if database.Password == "" {
		database.Password = defaults
	}
}

// CheckDatabase checks if the Type, Protocol, Hostname, Port, Name, User, and Password fields in the given ConnectionConfig are set from the DATABASE_TYPE, DATABASE_PROTOCOL, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME, DATABASE_USER, and DATABASE_PASSWORD environment variables.
// If they are not set, it sets them to the given default values.
func CheckDatabase(database *ConnectionConfig, defaults ConnectionConfig) {
	CheckType(database, defaults.Type)
	CheckProtocol(database, defaults.Protocol)
	CheckHostname(database, defaults.Hostname)
	CheckPort(database, defaults.Port)
	CheckName(database, defaults.Name)
	CheckUser(database, defaults.User)
	CheckPassword(database, defaults.Password)
}

// GetDatabase returns a ConnectionConfig struct with the Type, Protocol, Hostname, Port, Name, User, and Password fields set from the DATABASE_TYPE, DATABASE_PROTOCOL, DATABASE_HOST, DATABASE_PORT, DATABASE_NAME, DATABASE_USER, and DATABASE_PASSWORD environment variables.
// If they are not set, it sets them to the given default values.
func GetDatabase(defaults ConnectionConfig) ConnectionConfig {
	var database ConnectionConfig
	CheckDatabase(&database, defaults)
	return database
}
