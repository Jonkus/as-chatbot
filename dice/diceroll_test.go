package asdiceroll

import (
	"github.com/go-chat-bot/bot"
	"math/rand"
	"testing"
)

func TestDicerollString(t *testing.T) {
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

func TestDicerollAtASHelperBotString(t *testing.T) {
	arg := ""
	want := "(5 + 5) = 10"
	bot := &bot.Cmd{
		Command: "diceroll@ASHelperBot",
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
