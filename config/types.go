package config

type Configuration struct {
	Server   ServerConfiguration            `yaml:"server"`
	Database DatabaseConfiguration          `yaml:"database"`
	User     UserConfiguration              `yaml:"user"`
	Senders  map[string]SenderConfiguration `yaml:"senders"`
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
	Address string `yaml:"address"`
	Mode    string `yaml:"mode"`
	URL     string `yaml:"url"`
}

type UserConfiguration struct {
	// Admins
	Admins []string         `yaml:"admins"`
	SSO    SSOConfiguration `yaml:"sso"`
}

type DatabaseConfiguration struct {
	Type string `yaml:"type"`
	DSN  string `yaml:"dsn"`
}

type GitHubConfiguration struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type SSOConfiguration struct {
	GitHub GitHubConfiguration `yaml:"github"`
}

type SenderConfiguration = map[string]string
