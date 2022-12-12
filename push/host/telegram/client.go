package telegram

import (
	"errors"
	"fmt"
	"strconv"

	tgBot "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"

	"notify-api/db/util"
	"notify-api/utils/user"
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
	_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, startMessage))
	if err != nil {
		zap.S().Errorf("failed to send start message: %v", err)
	}
}

func (h *Host) handleHelpCommand(msg *tgBot.Message) {
	_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, helpMessage))
	if err != nil {
		zap.S().Errorf("failed to send help message: %v", err)
	}
}

func (h *Host) handleRegisterCommand(msg *tgBot.Message) {
	userID := msg.CommandArguments()
	if userID == "" {
		sendMsg := tgBot.NewMessage(msg.Chat.ID, "Please provide user id.\nexample: `/register test`")
		sendMsg.ReplyToMessageID = msg.MessageID
		sendMsg.ParseMode = tgBot.ModeMarkdown

		_, err := h.Bot.Send(sendMsg)
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
		}
		return
	}

	if !user.Is(userID) {
		_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, "Invalid user id "+userID))
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
		}
		return
	}

	chatID := strconv.FormatInt(msg.Chat.ID, 10)
	deviceID := int64toUUID(msg.From.ID)

	// check if user already registered
	pt, err := util.DeviceUtil.GetDevice(deviceID)
	if err != nil {
		if !errors.Is(err, util.ErrNotFound) {
			zap.S().Errorf("failed to get device token: %v", err)
			return
		}
	} else {
		var errText string
		if pt.UserID == userID {
			errText = fmt.Sprintf("You already registered with user id `%s`.", userID)
		} else {
			errText = fmt.Sprintf("You are already registered with user id `%s`.\nYou should first call `/unregister`", pt.UserID)
		}
		errMsg := tgBot.NewMessage(msg.Chat.ID, errText)
		errMsg.ParseMode = tgBot.ModeMarkdown
		_, err := h.Bot.Send(errMsg)
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
			return
		}
		return
	}

	err = util.DeviceUtil.CreateOrUpdate(userID, deviceID, h.Name(), chatID, "Telegram")
	if err != nil {
		zap.S().Errorf("failed to create or update token: %v", err)
		errText := fmt.Sprintf("Internal Error %+v", err)
		_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, errText))
		if err != nil {
			zap.S().Errorf("failed to send message: %v", err)
		}
		return
	}
	successText := fmt.Sprintf("Successfully registered user %s with %s", userID, msg.From.UserName)

	_, err = h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, successText))
	if err != nil {
		zap.S().Errorf("failed to send message: %v", err)
	}
}

func (h *Host) handleUnregisterCommand(msg *tgBot.Message) {
	deviceID := int64toUUID(msg.From.ID)

	// check if user already registered
	_, err := util.DeviceUtil.GetDevice(deviceID)
	if err != nil {
		if errors.Is(err, util.ErrNotFound) {
			_, err := h.Bot.Send(tgBot.NewMessage(msg.Chat.ID, "You are not registered yet."))
			if err != nil {
				zap.S().Errorf("failed to send message: %v", err)
			}
			return
		} else {
			zap.S().Errorf("failed to get device token: %v", err)
			return
		}
	}
	err = util.DeviceUtil.DeleteDevice(deviceID)
	if err != nil {
		zap.S().Errorf("failed to delete device token: %v", err)
		return
	}
	msgText := fmt.Sprintf("Successfully unregistered `%s`", msg.From.UserName)
	tipMsg := tgBot.NewMessage(msg.Chat.ID, msgText)
	tipMsg.ReplyToMessageID = msg.MessageID
	tipMsg.ParseMode = tgBot.ModeMarkdown
	_, err = h.Bot.Send(tipMsg)
	if err != nil {
		zap.S().Errorf("failed to send message: %v", err)
	}
}

func (h *Host) handleUnknownCommand(msg *tgBot.Message) {
	sendMsg := tgBot.NewMessage(msg.Chat.ID, "Unknown command "+msg.Command())
	sendMsg.ReplyToMessageID = msg.MessageID
	_, err := h.Bot.Send(sendMsg)
	if err != nil {
		zap.S().Errorf("failed to send message: %v", err)
	}
}
