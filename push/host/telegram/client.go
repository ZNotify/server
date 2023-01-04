package telegram

import (
	"context"
	"fmt"
	"strconv"

	tgBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"

	"notify-api/ent/dao"
	"notify-api/ent/helper"
)

func (h *Host) setCommand() {
	commands := tgBot.NewSetMyCommands(tgBot.BotCommand{
		Command:     "start",
		Description: "Start the bot",
	}, tgBot.BotCommand{
		Command:     "help",
		Description: "Get available commands",
	}, tgBot.BotCommand{
		Command:     "register",
		Description: "Register your telegram account with user id.",
	}, tgBot.BotCommand{
		Command:     "unregister",
		Description: "Unregister your telegram account.",
	})
	_, err := h.Bot.Request(commands)
	if err != nil {
		zap.S().Errorf("failed to set command: %v", err)
	}
}

func (h *Host) commandRoutine() {
	h.setCommand()

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

		command := update.Message.Command()

		switch command {
		case "start":
			h.handleStartCommand(update.Message)
		case "help":
			h.handleHelpCommand(update.Message)
		case "register":
			h.handleRegisterCommand(update.Message)
		case "unregister":
			h.handleUnregisterCommand(update.Message)
		default:
			h.handleUnknownCommand(update.Message)
		}
	}
}

func (h *Host) handleStartCommand(msg *tgBot.Message) {
	h.send(tgBot.NewMessage(msg.Chat.ID, startMessage))
}

func (h *Host) handleHelpCommand(msg *tgBot.Message) {
	h.send(tgBot.NewMessage(msg.Chat.ID, helpMessage))
}

func (h *Host) handleRegisterCommand(msg *tgBot.Message) {
	ctx := context.TODO()

	userSecret := msg.CommandArguments()
	if userSecret == "" {
		sendMsg := tgBot.NewMessage(msg.Chat.ID, "Please provide user secret.\nexample: `/register secret`")
		sendMsg.ReplyToMessageID = msg.MessageID
		sendMsg.ParseMode = tgBot.ModeMarkdown

		h.send(sendMsg)
		return
	}

	u, exist := dao.User.GetUserBySecret(ctx, userSecret)

	if !exist {
		_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, "Invalid user secret."))
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
		}
		return
	}

	chatID := strconv.FormatInt(msg.Chat.ID, 10)
	deviceIdentifier := strconv.FormatInt(msg.From.ID, 10)

	// check if user already registered
	du, exist := dao.User.GetDeviceUser(ctx, deviceIdentifier)
	if exist {
		var errText string
		if du.ID == u.ID {
			errText = fmt.Sprintf("You already registered with user id `%s`.", userSecret)
		} else {
			errText = fmt.Sprintf("You are already registered with user `%s`.\nYou should first call `/unregister`", helper.GetReadableName(du))
		}
		errMsg := tgBot.NewMessage(msg.Chat.ID, errText)
		errMsg.ParseMode = tgBot.ModeMarkdown
		_, err := h.Bot.Send(errMsg)
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
		}
		return
	}
	_, ok := dao.Device.EnsureDevice(ctx,
		deviceIdentifier,
		h.Name(),
		"",
		chatID,
		"Telegram",
		msg.From.String(),
		u)
	if !ok {
		_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, "Internal Error"))
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
		}
		return
	}
	successText := fmt.Sprintf("Successfully registered user %s with %s", userSecret, msg.From.UserName)

	h.send(tgBot.NewMessage(msg.Chat.ID, successText))
}

func (h *Host) handleUnregisterCommand(msg *tgBot.Message) {
	ctx := context.TODO()
	deviceIdentifier := strconv.FormatInt(msg.From.ID, 10)

	// check if user already registered
	_, exist := dao.User.GetDeviceUser(ctx, deviceIdentifier)

	if !exist {
		_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, "You are not registered yet."))
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
		}
		return
	}

	ok := dao.Device.DeleteDeviceByIdentifier(ctx, deviceIdentifier)
	if !ok {
		h.send(tgBot.NewMessage(msg.Chat.ID, "Internal Error"))
		return
	}
	msgText := fmt.Sprintf("Successfully unregistered `%s`", msg.From.UserName)

	tipMsg := tgBot.NewMessage(msg.Chat.ID, msgText)
	tipMsg.ReplyToMessageID = msg.MessageID
	tipMsg.ParseMode = tgBot.ModeMarkdown
	h.send(tipMsg)
}

func (h *Host) handleUnknownCommand(msg *tgBot.Message) {
	sendMsg := tgBot.NewMessage(msg.Chat.ID, "Unknown command "+msg.Command())
	sendMsg.ReplyToMessageID = msg.MessageID
	h.send(sendMsg)
}

func (h *Host) send(msg tgBot.MessageConfig) {
	_, err := h.Bot.Send(msg)
	if err != nil {
		zap.S().Errorf("failed to send message: %v", err)
	}
}
