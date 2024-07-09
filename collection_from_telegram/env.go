package main

import (
	"fmt"
	"os"
	"strconv"
)

type Environment struct {
	TelegramBotToken string
	TelegramChatId   string
	PostgresIP       string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	PreviousTextSize int
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
	postgres_ip := os.Getenv("POSTGRES_IP")
	if chat_id == "" {
		return Environment{}, fmt.Errorf("POSTGRES_IP 환경 변수를 설정해주세요")
	}

	postgres_user := os.Getenv("POSTGRES_USER")
	if postgres_user == "" {
		return Environment{}, fmt.Errorf("POSTGRES_USER 환경 변수를 설정해주세요")
	}
	postgres_password := os.Getenv("POSTGRES_PASSWORD")
	if postgres_password == "" {
		return Environment{}, fmt.Errorf("POSTGRES_PASSWORD 환경 변수를 설정해주세요")
	}
	postgres_db := os.Getenv("POSTGRES_DB")
	if postgres_db == "" {
		return Environment{}, fmt.Errorf("POSTGRES_DB 환경 변수를 설정해주세요")
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

	previous_text_size := os.Getenv("FROM_DATABASE_READ_SIZE")
	if previous_text_size == "" {
		return Environment{}, fmt.Errorf("FROM_DATABASE_READ_SIZE 환경 변수를 설정해주세요")
	}
	size, _ := strconv.ParseInt(previous_text_size, 10, 32)

	return Environment{
		TelegramBotToken: token,
		TelegramChatId:   chat_id,
		PostgresIP:       postgres_ip,
		PostgresUser:     postgres_user,
		PostgresPassword: postgres_password,
		PostgresDB:       postgres_db,
		GeminiApiKey:     gemini_api_key,
		TelegramJubuId:   int(id),
		PreviousTextSize: int(size),
	}, nil
}
