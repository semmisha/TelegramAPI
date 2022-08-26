package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

type Telegram struct {
	ApiKey     string
	ChannelKey string
	Bot        *tgbotapi.BotAPI
	Logger     *logrus.Logger
}

func NewTelegram(apiKey string, channelKey string) *Telegram {
	return &Telegram{
		ApiKey:     apiKey,
		ChannelKey: channelKey,
	}
}
