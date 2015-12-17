package asdiceroll

import (
	"github.com/go-chat-bot/bot"
	"math/rand"
	"os"
	"testing"
)

func TestDiceroll(t *testing.T) {
	arg := ""
	want := "(5 + 5) = 10"
	bot := &bot.Cmd{
		Command: "diceroll",
		RawArgs: arg,
	}

	rand.Seed(0)
	got, error := diceroll(bot)

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}

	if error != nil {
		t.Errorf("Expected '%v' got '%v'", nil, error)
	}
}

func TestDicerollAtBotname(t *testing.T) {
	arg := ""
	want := "(5 + 5) = 10"
	bot := &bot.Cmd{
		Command: "diceroll@" + os.Getenv("TELEGRAM_BOTNAME"),
		RawArgs: arg,
	}

	rand.Seed(0)
	got, error := diceroll(bot)

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}

	if error != nil {
		t.Errorf("Expected '%v' got '%v'", nil, error)
	}
}
