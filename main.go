package main

import (
    "log"
    "time"
    "os"
    "strings"
    "github.com/tucnak/telebot"
)

func main() {
    bot, err := telebot.NewBot(os.Getenv("TELEGRAM_TOKEN"))
    if err != nil {
        log.Fatalln(err)
    }

    bot.Messages = make(chan telebot.Message, 100)
    bot.Queries = make(chan telebot.Query, 1000)

    go messages(bot)
    go queries(bot)

    bot.Start(1 * time.Second)
}

func messages(bot *telebot.Bot) {
    for message := range bot.Messages {
	command := strings.Split(strings.Fields(message.Text)[0], "@")[0]
	args := strings.Fields(message.Text)[1:]
	log.Printf("Message from %s: \"%s\", Command: %v, Args:%q", message.Sender, message.Text, command, args)
	
	switch command {
	case "/diceroll":
	    response, err := diceroll()
	    if (err==nil) {
		log.Printf("Response: %s, Chat: %d", response, message.Chat)
            	go bot.SendMessage(message.Chat, response, nil)
	    } else {
		log.Printf("Error: %s", err)
	    }

        case "/motive":
	    response := motivehit(args)
	    log.Printf("Response: %s, Chat: %d", response, message.Chat)
	    go bot.SendMessage(message.Chat, response, nil)

        case "/critical":
	    response := critical(args)
	    log.Printf("Response: %s, Chat: %d", response, message.Chat)
	    go bot.SendMessage(message.Chat, response, nil)

        case "/search":
	    go func () {
	        response := unitsearch(args)
		log.Printf("Response: %s, Chat: %d", response, message.Chat)
	        bot.SendMessage(message.Chat, response, nil)
	    } ()

	default:
	    go bot.SendMessage(message.Chat, "Unkown Command", nil)
	}
    }
}

func queries(bot *telebot.Bot) {
    for query := range bot.Queries {
        log.Println("--- new query ---")
        log.Println("from:", query.From.Username)
        log.Println("text:", query.Text)

        // Create an article (a link) object to show in results.
        article := &telebot.InlineQueryResultArticle{
            Title: "Telebot",
            URL:   "https://github.com/tucnak/telebot",
            InputMessageContent: &telebot.InputTextMessageContent{
                Text:          "Telebot is a Telegram bot framework.",
                DisablePreview: false,
            },
        }

        // Build the list of results (make sure to pass pointers!).
        results := []telebot.InlineQueryResult{article}

        // Build a response object to answer the query.
        response := telebot.QueryResponse{
            Results:    results,
            IsPersonal: true,
        }

        // Send it.
        if err := bot.AnswerInlineQuery(&query, &response); err != nil {
            log.Println("Failed to respond to query:", err)
        }
    }
}
