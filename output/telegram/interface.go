package telegram

type TelegramInt interface {
	SetBotKey(ApiKey string) *TelegramInt
	SetChannel(ChannelKey string) *TelegramInt
	Write([]byte) (int, error)
}
