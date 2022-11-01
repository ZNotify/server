package host

import (
	"github.com/pkg/errors"

	pushTypes "notify-api/push/types"
	"notify-api/utils/config"

	tgBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramHost struct {
	BotToken string
	Bot      *tgBot.BotAPI
}

func (h *TelegramHost) commandRoutine() {
	// setCommand

	u := tgBot.NewUpdate(0)
	u.Timeout = 60

	updates := h.Bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if !update.Message.IsCommand() {
			continue
		}

	}
}

func (h *TelegramHost) Start() error {
	go h.commandRoutine()
	return nil
}

func (h *TelegramHost) Send(msg *pushTypes.Message) error {
	//TODO implement me
	panic("implement me")
}

func (h *TelegramHost) Init() error {
	bot, err := tgBot.NewBotAPI(h.BotToken)
	if err != nil {
		return errors.WithStack(err)
	}

	if config.Config.Server.Mode == config.ProdMode {
		bot.Debug = false
	} else {
		bot.Debug = true
	}
	h.Bot = bot
	return nil
}

func (h *TelegramHost) Name() string {
	return "TelegramHost"
}

func (h *TelegramHost) Check(auth pushTypes.SenderAuth) error {
	token, ok := auth["BotToken"]
	if !ok {
		return errors.New("BotToken not found")
	}
	h.BotToken = token
	return nil
}
