package senderConfig

type TelegramConfig struct {
	// Bot token from @BotFather
	BotToken string `json:"bot_token" jsonschema:"required,title=Telegram bot token"`
}
