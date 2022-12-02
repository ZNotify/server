package telegram

import (
	"fmt"
	"strconv"
	"time"

	tgBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"notify-api/db/model"
	pushTypes "notify-api/push/types"
	"notify-api/utils"
	"notify-api/utils/config"
)

type Host struct {
	BotToken string
	Bot      *tgBot.BotAPI
}

func (h *Host) Start() error {
	go h.commandRoutine()
	return nil
}

func (h *Host) Send(msg *pushTypes.Message) error {
	tokens, err := model.TokenUtils.GetUserChannelTokens(msg.UserID, h.Name())
	if err != nil {
		return errors.WithStack(err)
	}
	if len(tokens) == 0 {
		return nil
	}

	for _, v := range tokens {
		id, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return errors.WithStack(err)
		}

		msgText := fmt.Sprintf("*%s*\n\n%s\n\n",
			msg.Title,
			msg.Content,
		)
		if msg.Long != "" {
			msgText += fmt.Sprintf("%s\n\n", msg.Long)
		}
		msgText += fmt.Sprintf("`%s`", msg.CreatedAt.Format(time.RFC3339))

		tgMsg := tgBot.NewMessage(id, msgText)
		tgMsg.ParseMode = tgBot.ModeMarkdown

		if msg.Priority == pushTypes.PriorityLow {
			tgMsg.DisableNotification = true
		}

		_, err = h.Bot.Send(tgMsg)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (h *Host) Init() error {
	cfg := config.Config.Senders[h.Name()].(Config)
	h.BotToken = utils.TokenClean(cfg.BotToken)

	err := tgBot.SetLogger(loggerAdapter)
	if err != nil {
		return errors.WithStack(err)
	}

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

func (h *Host) Name() string {
	return "TelegramHost"
}

type Config struct {
	BotToken string
}

func (h *Host) Config() any {
	return Config{}
}
