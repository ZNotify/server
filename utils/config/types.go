package config

var Config Configuration

type Configuration struct {
	// Server config
	Server ServerConfiguration `yaml:"server"`
	// Database config
	Database DatabaseConfiguration `yaml:"database"`
	// Senders config
	Senders map[string]SenderConfiguration `yaml:"senders"`
}

const (
	TestMode = "test"
	DevMode  = "development"
	ProdMode = "production"
)

const (
	Sqlite = "sqlite"
	Mysql  = "mysql"
	Pgsql  = "pgsql"
)

type ServerConfiguration struct {
	// Server port
	Port int `yaml:"port"`
	// Server host
	Host string `yaml:"host"`
	// Server mode
	Mode string `yaml:"mode"`
}

type DatabaseConfiguration struct {
	// Database type
	Type string `yaml:"type"`
	// Database DSN
	DSN string `yaml:"dsn"`
}

type SenderConfiguration = map[string]string
