package main

import (
	"os"

	"github.com/go-chat-bot/bot/telegram"
	_ "github.com/Jonkus/as-chatbot/critical"
	_ "github.com/Jonkus/as-chatbot/dice"
	_ "github.com/Jonkus/as-chatbot/motivehit"
	_ "github.com/Jonkus/as-chatbot/unitsearch"
)

func main() {
	telegram.Run(os.Getenv("TELEGRAM_TOKEN"), os.Getenv("DEBUG") != "")

}
