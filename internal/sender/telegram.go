package sender

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"go_cli/pkg/interfaceManager"
	"go_cli/pkg/logger"
	"os"
	"strconv"
)

type TelegramSender struct {
	bot    *tgbotapi.BotAPI
	chatID int64
}

func NewTelegramSender() *TelegramSender {
	token := os.Getenv("TELEGRAM_BOT_TOKEN")
	if token == "" {
		logger.Error("TELEGRAM_BOT_TOKEN environment variable is not set")
		return nil
	}
	chatID, err := strconv.Atoi(os.Getenv("TELEGRAM_CHAT_ID"))
	if err != nil {
		logger.Error("Error in TELEGRAM_CHAT_ID environment variable")
		return nil
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		logger.Error(err)
	}
	return &TelegramSender{bot: bot, chatID: int64(chatID)}
}

func (t *TelegramSender) Configure() error {
	interfaceManager.Clear()

	interfaceManager.SetText("Enter the Telegram bot token: ")
	token, err := interfaceManager.GetEnteredText()
	if err != nil {
		return err
	}
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return err
	}
	t.bot = bot

	interfaceManager.SetText("Enter the Telegram chat Id: ")
	chatIDStr, err := interfaceManager.GetEnteredText()
	if err != nil {
		return err
	}
	chatID, err := strconv.Atoi(chatIDStr)
	t.chatID = int64(chatID)

	return nil
}

func (t *TelegramSender) Send(msg string) error {
	message := tgbotapi.NewMessage(t.chatID, msg)

	_, err := t.bot.Send(message)
	if err != nil {
		return err
	}

	return nil
}

func (t *TelegramSender) GetName() string {
	return "Telegram"
}
