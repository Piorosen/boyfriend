package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

	env, err := GetEnvironment()
	if err != nil {
		log.Fatal(err)
	}
	client := NewClient()
	err = client.Connect(env.PostgresIP, env.PostgresPort, env.PostgresDB, env.PostgresUser, env.PostgresPassword)
	if err != nil {
		log.Panic(err)
	}

	bot, err := tgbotapi.NewBotAPI(env.TelegramBotToken)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		receive_chat_id := strconv.FormatInt(update.Message.Chat.ID, 10)
		if receive_chat_id == env.TelegramChatId {
			if update.Message != nil { // If we got a message
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				if len(update.Message.Text) == 0 {
					log.Printf("Sticker, Animation or Uploaded Image")
				} else {
					if update.Message.Text[0] == '!' {
						err = client.Insert(update.Message.From.FirstName,
							update.Message.From.LastName,
							update.Message.From.UserName,
							update.Message.Text,
							int64(update.Message.MessageID),
							update.Message.From.ID,
						)
					}
					output := client.Process(update.Message.Text, env)
					if output != "" {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
						msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)
					} else if client.Run() {
						err = client.Insert(update.Message.From.FirstName,
							update.Message.From.LastName,
							update.Message.From.UserName,
							update.Message.Text,
							int64(update.Message.MessageID),
							update.Message.From.ID,
						)
					}
					if err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, err.Error())
						msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)
					}
				}
			}
		} else {
			log.Printf("Detect Hacker! [%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}
}
