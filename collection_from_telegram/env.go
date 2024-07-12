package main

import (
	"fmt"
	"os"
	"strconv"
)

type Environment struct {
	TelegramBotToken string
	TelegramChatId   string
	TelegramJubuId   int
	GeminiApiKey     string
}

func GetEnvironment() (Environment, error) {
	token := os.Getenv("TELEGRAM_BOT_API_TOKEN")
	if token == "" {
		return Environment{}, fmt.Errorf("TELEGRAM_BOT_API_TOKEN 환경 변수를 설정해주세요")
	}

	chat_id := os.Getenv("TELEGRAM_BOT_API_CHAT_ID")
	if chat_id == "" {
		return Environment{}, fmt.Errorf("TELEGRAM_BOT_API_CHAT_ID 환경 변수를 설정해주세요")
	}

	gemini_api_key := os.Getenv("GEMINI_API_KEY")
	if gemini_api_key == "" {
		return Environment{}, fmt.Errorf("GEMINI_API_KEY 환경 변수를 설정해주세요")
	}

	jubu_telegram_id := os.Getenv("JUBU_TELEGRAM_ID")
	if jubu_telegram_id == "" {
		return Environment{}, fmt.Errorf("JUBU_TELEGRAM_ID 환경 변수를 설정해주세요")
	}
	id, _ := strconv.ParseInt(jubu_telegram_id, 10, 32)

	return Environment{
		TelegramBotToken: token,
		TelegramChatId:   chat_id,
		GeminiApiKey:     gemini_api_key,
		TelegramJubuId:   int(id),
	}, nil
}
