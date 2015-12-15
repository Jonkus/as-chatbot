package asdiceroll

import (
	"fmt"
	"github.com/go-chat-bot/bot"
	"math/rand"
)

func diceroll(command *bot.Cmd) (string, error) {

	d1 := rand.Intn(5) + 1
	d2 := rand.Intn(5) + 1

	return fmt.Sprintf("(%d + %d) = %d", d1, d2, d1+d2), nil
}

func init() {
	bot.RegisterCommand(
		"diceroll",
		"rolls 2D6",
		"",
		diceroll)
	bot.RegisterCommand(
		"diceroll@ASHelperBot",
		"rolls 2D6",
		"",
		diceroll)
}
