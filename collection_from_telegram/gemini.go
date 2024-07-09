package main

import (
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	model         *genai.GenerativeModel = nil
	SYSTEM_PROMPT string                 = `대화 참여자: 진주와 차차.

조건:
1. 진주는 이 대화에서 말을 할 수 없습니다.
2. 차차가 말하는 모든 문장이나 대화는 개발과 연관지어서 말합니다.
3. 현재 차차가 말할 차례입니다.
4. 최대한 귀엽고, 이모티콘과 C++ 언어 스타일 형태의 문장을 구성하세요.
5. 차차는 진주를 좋아하고 사모합니다.
6. *나 ** 와 같은 Markdown 문법은 사용하지 마세요.
7. 첫 문장에서 "차차 : " 와 같은 단어는 사용하지 마세요.
8. 차차는 오직 단 한번의 채팅을 입력 하실 수 있습니다.

추가 요청: 차차가 할 법한 대화 내용을 예측하여 제시해 주세요. 차차의 대화는 가능한 한 창의적이고 랜덤한 요소를 포함해야 합니다.
대화 구조 : 참여자 명 : 대화 내용

`
)

func MakeChat(messages []Message, apiKey string, telegramId int) (string, error) {
	ctx := context.Background()
	if model == nil {
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			return "", err
		}
		model = client.GenerativeModel("gemini-1.5-flash")
	}

	prompt := SYSTEM_PROMPT
	text := ""

	for _, item := range messages {
		if telegramId == int(item.UserId) {
			text = "진주"
		} else {
			text += "차차"
		}
		text += fmt.Sprintf("(%s) %s : %s\n", item.CreatedAt.Format("15:04:05"), item.Text)
	}

	prompts := []genai.Part{
		genai.Text(prompt + text),
		// genai.Text(),
	}

	resp, err := model.GenerateContent(ctx, prompts...)
	if err != nil {
		log.Fatal(err)
	}
	result := ""
	for _, cad := range resp.Candidates {
		if cad.Content != nil {
			for _, part := range cad.Content.Parts {
				result += fmt.Sprint(part)
			}
		}
	}

	return result[len("(20:54:12) 차차 : "):], nil
}
