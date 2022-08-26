package processor

import "TelegramAPI/output/telegram"

func ProceedMessage(message string, channel string) {

	t := telegram.New(message, channel)

}
