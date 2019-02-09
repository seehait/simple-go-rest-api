package configs

// DatabaseConfig represents database configurations.
type DatabaseConfig struct {
	Name     string
	UserName string
	Password string
	Host     string
	Port     uint
}

// GetDatabaseConfig provides database configurations.
func GetDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{Name: "simple-go-rest-api", UserName: "user", Password: "password", Host: "127.0.0.1", Port: 3306}
}
