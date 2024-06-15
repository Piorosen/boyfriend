package main

import (
	"fmt"
	"os"
)

type Environment struct {
	TelegramBotToken string
	TelegramChatId   string
	PostgresIP       string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
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
	return Environment{
		TelegramBotToken: token,
		TelegramChatId:   chat_id,
		PostgresIP:       postgres_ip,
		PostgresUser:     postgres_user,
		PostgresPassword: postgres_password,
		PostgresDB:       postgres_db,
	}, nil
}
