package main

import (
	"TelegramAPI/input/http"
	"TelegramAPI/logging"
)

func main() {
	logger := logging.Logger()
	tlgrmAPI := http.NewApi(logger)
	http.NewRouter(tlgrmAPI)

	//TODO ----- SetUp flags ----- //

}
