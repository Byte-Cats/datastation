package pkg

func DefaultPostgres() DatabaseConfig {
	return DatabaseConfig{
		Type:     "postgres",
		Protocol: "tcp",
		Hostname: "localhost",
		Port:     "5432",
		Name:     "database",
		Username: "user",
		Password: "password",
	}
}
