package config

import (
	senderConfig "notify-api/app/config/sender"
)

// Configuration is the top level configuration.
type Configuration struct {
	// Server configuration
	Server ServerConfiguration `yaml:"server"`
	// Database configuration
	Database DatabaseConfiguration `yaml:"database"`
	// User configuration
	User UserConfiguration `yaml:"user"`
	// Senders configuration,
	Senders SenderConfiguration `yaml:"senders"`
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

type SenderConfiguration struct {
	Telegram  senderConfig.TelegramConfig `yaml:"telegram"`
	FCM       senderConfig.FCMConfig      `yaml:"fcm"`
	WebSocket bool                        `yaml:"websocket"`
	WebPush   senderConfig.WebPushConfig  `yaml:"webpush"`
	WNS       bool                        `yaml:"wns"`
}
