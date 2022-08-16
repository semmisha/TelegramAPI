package main

import (
	"TelegramCLII/logging"
	"TelegramCLII/output/telegram"

	"flag"
)

var (
	tlgrm = telegram.New()
)

func main() {
	logger := logging.Logger()
	//TODO ----- SetUp flags ----- //

	apiKey := flag.String("a", "5203368767:AAH5fdcTNV2Zj41Cp3SbTXYYWftY7H5eze0", "Enter BotFather API key")
	chanellKey := flag.String("ch", "@test111_111", "Enter channel name")
	message := flag.String("msg", "Test", "Enter desired message")
	flag.Parse()

	//TODO ----- SetUp Bot and Chanell -----//

	err := tlgrm.SetBotKey(*apiKey)
	if err != nil {
		logger.Fatalf("Unable to set bot", err)
	}
	err = tlgrm.SetChannel(*chanellKey)
	if err != nil {
		logger.Fatal(err)
	}

	//TODO ----- process message ----- //

	_, err = tlgrm.Write([]byte(*message))
	if err != nil {

		logger.Fatal(err)

	}

}
