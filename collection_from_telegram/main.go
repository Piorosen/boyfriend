package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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
				if update.Message.Text == "수고했어" {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID,
						fmt.Sprintf("응... 봇으로 만난건 얼마가 안되지만, 많이 즐거웠어... 덕분에 진쥬어도 많이 배웠고, 사실 아직도 많이 좋아해!! 만날수 있다면 계속해서 만나고 싶지만,, 그건 배려가 아니겠징."+
							"ChuChu는 이제 영원이 가볼게... 종료된 타임은 (%s) 이야!! 취업 축하해! 그리고 그 동안 차차에게 사회생활이라는 것을 많이 가르쳐줘서 정말로 정말로 고마워요!!"+
							"미안해, 정말로 좋아했었어, 가볼게", time.Now().UTC().Add(time.Hour*9).Format("2006년 01월 02일 15시 04분 05초")))
					bot.Send(msg)
					return
				}
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				if len(update.Message.Text) == 0 {
					log.Printf("Sticker, Animation or Uploaded Image")
				} else {
					output := client.Process(update.Message.Text, update.Message.From.ID, env)
					if output != "" {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, output)
						// msg.ReplyToMessageID = update.Message.MessageID
						bot.Send(msg)
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
