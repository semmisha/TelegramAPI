package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Telegram struct {
	ApiKey     string
	ChannelKey string
	Bot        *tgbotapi.BotAPI
}

func New() *Telegram {
	return &Telegram{}
}

func (t *Telegram) SetBotKey(ApiKey string) (err error) {

	t.Bot, err = tgbotapi.NewBotAPI(ApiKey)
	return err
}
func (t *Telegram) SetChannel(ChannelKey string) (err error) {
	if len(ChannelKey) != 0 {
		t.ChannelKey = ChannelKey
		return nil
	}
	return fmt.Errorf("ApiKey is empty")
}
func (t *Telegram) Write(msg []byte) (int, error) {
	var (
		bot     = t.Bot
		channel = t.ChannelKey
	)
	_, err := bot.Send(tgbotapi.NewMessageToChannel(channel, string(msg)))

	return 0, err
}
