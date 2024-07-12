package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"strings"
)

var (
	BUILD_DATE_TIME = "null"
	DEPLOY_VERSION  = "0.0.0"
	ALL_DATA_REMOVE = GenerateRandomHex(16)
)

func GenerateRandomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(bytes)
}

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}
func (client *Client) Process(text string, user_id int64, env Environment) string {
	if len(text) == 0 {
		return ""
	}
	if text[0] == '!' {
		// jsonData := fmt.Sprintf(`{"size": %d, "jubu_id": %d}`, 100, env.TelegramJubuId)
		// message := client.GetText(1)
		result, err := MakeChat(text[1:], user_id, env.GeminiApiKey, env.TelegramJubuId)
		if err != nil {
			return err.Error()
		} else {
			return result
		}
	}

	if text[0] == '/' {
		switch strings.Split(strings.ToLower(text[1:]), " ")[0] {
		case "get_instruction":
			return GetSystemInstruction()
		case "set_instrctuin":
			SetSystemInstruction(text[len("/set_instrctuin "):])
			return GetSystemInstruction()
		case "history":
			return GetChatHistoryFromGemini()
		case "version":
			return fmt.Sprintf("Version : %s\nBuild Time : %s", DEPLOY_VERSION, BUILD_DATE_TIME)
		case "clear":
			data := strings.Split(text, " ")
			if len(data) != 2 {
				return fmt.Sprintf("아래의 명령어를 사용하시면 데이터베이스에 있는 모든 기록을 삭제합니다.\n/clear %s", ALL_DATA_REMOVE)
			}
			if data[1] == ALL_DATA_REMOVE {
				ClearChatHistory()
				return fmt.Sprintf("모든 처리를 수행하였습니다.")
			} else {
				return fmt.Sprintf("아래의 명령어를 사용하시면 데이터베이스에 있는 모든 기록을 삭제합니다.\n/clear %s", ALL_DATA_REMOVE)
			}
		// case "help":
		default:
			return "clear\nhistory\nversion\nget_instruction\nset_instruction"
		}
	}
	return ""
}
