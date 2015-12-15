package asdiceroll

import (
	"github.com/go-chat-bot/bot"
	"testing"
)

func TestReverseString(t *testing.T) {
	arg := ""
	want := "match"
	bot := &bot.Cmd{
		Command: "diceroll",
		RawArgs: arg,
	}

	got, error := reverse(bot)

	if got != want {
		t.Errorf("Expected '%v' got '%v'", want, got)
	}

	if error != nil {
		t.Errorf("Expected '%v' got '%v'", nil, error)
	}
}
