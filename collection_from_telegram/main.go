package main

import (
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	// BotFather로부터 받은 토큰을 환경 변수에서 읽어옵니다.
	token := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if token == "" {
		log.Fatal("TELEGRAM_BOT_API_TOKEN 환경 변수를 설정해주세요.")
	}

	chat_id := os.Getenv("TELEGRAM_BOT_API_CHAT_ID")
	if chat_id == "" {
		log.Fatal("TELEGRAM_BOT_API_CHAT_ID 환경 변수를 설정해주세요.")
	}

	bot, err := tgbotapi.NewBotAPI(token)
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

		if receive_chat_id == chat_id {
			if update.Message != nil { // If we got a message
				// if update.Id
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
				msg.ReplyToMessageID = update.Message.MessageID

				// bot.Send(msg)
			}
		} else {
			log.Printf("Detect Hacker! [%s] %s", update.Message.From.UserName, update.Message.Text)
		}
	}
}
