package main

import (
	"log"
	"os"
	"time"

	"github.com/jonkus/as-chatbot/aslib"
	"github.com/tucnak/telebot"
)

func main() {
	token := os.Getenv("TELEGRAM_TOKEN")
	log.Printf("connecting to bot: %v", token)
	bot, err := telebot.NewBot(telebot.Settings{
		Token:  token,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	})
	if err != nil {
		log.Fatalln(err)
	}

	messages(bot)

	bot.Start()
}

func messages(bot *telebot.Bot) {

	bot.Handle("/diceroll", func(m *telebot.Message) {
		go func() {
			response, _ := aslib.Diceroll()
			log.Printf("Req: '%s', Resp: %#v, Sender: %s", m.Text, response, m.Sender.Username)
			bot.Send(m.Sender, response)
		}()
	})

	bot.Handle("/motive", func(m *telebot.Message) {
		go func() {
			response := aslib.Motivehit(m.Payload)
			log.Printf("Req: '%s', Resp: %#v, Sender: %s", m.Text, response, m.Sender.Username)
			bot.Send(m.Sender, response)
		}()
	})

	bot.Handle("/critical", func(m *telebot.Message) {
		go func() {
			response := aslib.Critical(m.Payload)
			log.Printf("Req: '%s', Resp: %#v, Sender: %s", m.Text, response, m.Sender.Username)
			bot.Send(m.Sender, response)
		}()
	})

	bot.Handle("/search", func(m *telebot.Message) {
		go func() {
			response := aslib.Unitsearch(m.Payload)
			log.Printf("Req: '%s', Resp: %#v, Sender: %s", m.Text, response, m.Sender.Username)
			bot.Send(m.Sender, response)
		}()
	})

	//go bot.SendMessage(message.Chat, "Unkown Command", nil)
}
