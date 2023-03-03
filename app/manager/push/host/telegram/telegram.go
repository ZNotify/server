package telegram

import (
	"context"
	"fmt"
	"strconv"
	"time"

	tgBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	"github.com/ZNotify/server/app/config/sender"
	"github.com/ZNotify/server/app/db/dao"
	"github.com/ZNotify/server/app/global"
	"github.com/ZNotify/server/app/manager/push/enum"
	"github.com/ZNotify/server/app/manager/push/interfaces"
	"github.com/ZNotify/server/app/manager/push/item"
	"github.com/ZNotify/server/app/utils"
)

type Host struct {
	BotToken string
	Bot      *tgBot.BotAPI
}

func (h *Host) Setup() error {
	go h.commandRoutine()
	return nil
}

func (h *Host) Send(ctx context.Context, msg *item.PushMessage) error {
	tokens, ok := dao.Device.GetUserDeviceChannelTokens(ctx, msg.User, h.Name())
	if !ok {
		return errors.New("telegram get user device channel tokens failed")
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

		if msg.Priority == enum.PriorityLow {
			tgMsg.DisableNotification = true
		}

		_, err = h.Bot.Send(tgMsg)
		if err != nil {
			return errors.WithStack(err)
		}
	}
	return nil
}

func (h *Host) Init(config interfaces.Config) error {
	cfg, ok := config.(senderConfig.TelegramConfig)
	if !ok {
		return errors.New("telegram config type error")
	}
	h.BotToken = utils.YamlStringClean(cfg.BotToken)

	err := tgBot.SetLogger(loggerAdapter)
	if err != nil {
		return errors.WithStack(err)
	}

	bot, err := tgBot.NewBotAPI(h.BotToken)
	if err != nil {
		return errors.WithStack(err)
	}

	if global.IsProd() {
		bot.Debug = false
	} else {
		bot.Debug = true
	}
	h.Bot = bot
	return nil
}

func (h *Host) Name() enum.Sender {
	return enum.SenderTelegram
}
