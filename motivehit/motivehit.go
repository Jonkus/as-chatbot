package asmotivehit

import (
	"fmt"
	"github.com/go-chat-bot/bot"
	"math/rand"
)

func motivehit(command *bot.Cmd) (string, error) {

	d1 := rand.Intn(5) + 1
	d2 := rand.Intn(5) + 1

	roll := d1 + d2
	movementMode := ""
	if len(command.Args) > 0 {
		movementMode = command.Args[0]
	}

	mod := 0
	switch movementMode {
	case "t", "s", "n":
		mod = 0
	case "w", "h":
		mod = 1
	case "v", "g":
		mod = 2
	}

	effect := "No Effect"

	if (roll + mod) > 12 {
		effect = "Unit Immobilized"
	} else if (roll + mod) > 10 {
		effect = "minus 50% Move (round down)"
	} else if (roll + mod) > 8 {
		effect = "minus 2 inch Move (round down)"
	}

	return fmt.Sprintf("[ (%d + %d) + %d = %d ]\n--->%s<----", d1, d2, mod, roll+mod, effect), nil
}

func init() {
	bot.RegisterCommand(
		"motive",
		"makes a motive-systems-damage-roll",
		"[h|n|s|t|v|w|g] (movement mode)",
		motivehit)
	bot.RegisterCommand(
		"motive@ASHelperBot",
		"makes a motive-systems-damage-roll",
		"[h|n|s|t|v|w|g] (movement mode)",
		motivehit)
}
