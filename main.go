package main

import (
	"os"

	"github.com/go-chat-bot/bot/telegram"
	_ "jonku/alpha-strike-bot/critical"
	_ "jonku/alpha-strike-bot/dice"
	_ "jonku/alpha-strike-bot/motivehit"
	_ "jonku/alpha-strike-bot/unitsearch"
)

func main() {
	telegram.Run(os.Getenv("TELEGRAM_TOKEN"), os.Getenv("DEBUG") != "")

}
