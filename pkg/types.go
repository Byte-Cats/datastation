package pkg

type DatabaseConfig struct {
	Type     string
	Protocol string
	Hostname string
	Port     string
	Name     string
	Username string
	Password string
}

func DefaultMongoDB() DatabaseConfig {
	return DatabaseConfig{
		Type:     "mongodb",
		Protocol: "tcp",
		Hostname: "localhost",
		Port:     "27017",
		Name:     "mydb",
		Username: "user",
		Password: "password",
	}
}
func DefaultRedis() DatabaseConfig {
	return DatabaseConfig{
		Type:     "redis",
		Protocol: "tcp",
		Hostname: "localhost",
		Port:     "6379",
		Name:     "",
		Username: "",
		Password: "",
	}
}

// DefaultMysql returns a DatabaseConfig struct with default values for the fields.
func DefaultMySql() DatabaseConfig {
	return DatabaseConfig{
		Type:     "mysql",
		Protocol: "tcp",
		Hostname: "localhost",
		Port:     "3306",
		Name:     "database",
		Username: "user",
		Password: "password",
	}
}
