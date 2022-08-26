package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (t *Telegram) SetBotKey(ApiKey string) *Telegram {

	t.Bot, _ = tgbotapi.NewBotAPI(ApiKey)
	return t
}
func (t *Telegram) SetChannel(ChannelKey string) *Telegram {
	if len(ChannelKey) != 0 {
		t.ChannelKey = ChannelKey
		return nil
	}
	return t
}

func (t *Telegram) Write(msg []byte) (int, error) {

	t.SetBotKey(t.ApiKey)
	t.SetChannel(t.ChannelKey)

	var (
		bot     = t.Bot
		channel = t.ChannelKey
	)
	_, err := bot.Send(tgbotapi.NewMessageToChannel(channel, string(msg)))

	return 0, err
}
