//go:build test

package config

func Load(path string) {
	Config = Configuration{
		Server: ServerConfiguration{
			Port: 14444,
			Host: "0.0.0.0",
			Mode: TestMode,
		},
		Database: DatabaseConfiguration{
			Type: Sqlite,
			DSN:  "data/notify.db",
		},
		Senders: make(map[string]SenderConfiguration),
		Users:   make(UserConfiguration, 0),
	}
	Config.Users = append(Config.Users, "test")
}
