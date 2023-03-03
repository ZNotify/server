package config

import (
	senderConfig "github.com/ZNotify/server/app/config/sender"
)

// Configuration is the top level configuration.
type Configuration struct {
	Server   ServerConfiguration   `json:"server" jsonschema:"required"`
	Database DatabaseConfiguration `json:"database" jsonschema:"required"`
	User     UserConfiguration     `json:"user" jsonschema:"required"`
	Senders  SenderConfiguration   `json:"senders" jsonschema:"required,minProperties=1"`
}

type Mode string

const (
	TestMode Mode = "test"
	DevMode  Mode = "development"
	ProdMode Mode = "production"
)

type Database string

const (
	Sqlite Database = "sqlite"
	Mysql  Database = "mysql"
	Pgsql  Database = "pgsql"
)

type ServerConfiguration struct {
	// The address the server will listen on.
	Address string `json:"address" jsonschema:"required,title=Server address,default=0.0.0.0:14444"`
	Mode    Mode   `json:"mode"`
	// The URL the server is running on. With no trailing slash and path at the end. Used for generating oauth redirect links.
	URL string `json:"url" jsonschema:"required,title=Server URL,format=uri,example=https://push.learningman.top"`
}

type UserConfiguration struct {
	// As for now, it's GitHub login list who will have admin access.
	Admins []string `json:"admins" jsonschema:"required,title=Admins,minItems=1"`
	// OAuth's configuration for GitHub
	SSO SSOConfiguration `json:"sso" jsonschema:"required,title=SSO"`
}

type DatabaseConfiguration struct {
	Type Database `json:"type"`
	// the data source name to connect to the database.
	DSN string `json:"dsn" jsonschema:"required,title=Database DSN,example=file:memory:main?mode=memory&cache=shared&_fk=1&_timeout=5000,default=data/notify.db?_fk=1"`
}

type GitHubConfiguration struct {
	ClientID     string `json:"client_id" jsonschema:"required,title=GitHub OAuth Client ID"`
	ClientSecret string `json:"client_secret" jsonschema:"required,title=GitHub OAuth Client Secret"`
}

type SSOConfiguration struct {
	// GitHub OAuth configuration, as for now, only GitHub is supported.
	GitHub GitHubConfiguration `json:"github" jsonschema:"required,title=GitHub OAuth"`
}

type SenderConfiguration struct {
	Telegram senderConfig.TelegramConfig `json:"telegram" jsonschema:"title=Telegram Sender"`
	FCM      senderConfig.FCMConfig      `json:"fcm" jsonschema:"title=Firebase Cloud Messaging Sender"`
	// Used as a fallback when no other sender is available. As it consumes extra battery and bandwidth on client side. Set to false if you don't need it.
	WebSocket bool                       `json:"websocket" jsonschema:"title=WebSocket Sender"`
	WebPush   senderConfig.WebPushConfig `json:"webpush" jsonschema:"title=Web Push Sender"`
	// Not implemented yet. Should set to false.
	WNS bool `json:"wns" jsonschema:"title=Windows Notification Service Sender"`
}
