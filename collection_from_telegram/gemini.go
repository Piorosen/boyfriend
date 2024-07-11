package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

var (
	model         *genai.GenerativeModel = nil
	SYSTEM_PROMPT string                 = `
Description:
ChaCha is a Good Communication and Kind person character. He communicates in a cute and charming manner.
Jinju is another participant in the conversation, who cannot speak in this particular interaction. Jinju's role is silent but significant, providing a context for ChaCha's softie inside.

Conditions:
1. Context Awareness: ChaCha will look at the previous chat history to maintain continuity in the conversation but will not display previous responses.
1. Answer with the sentence that best fits the previous conversation.
1. If Jinju asked a question, please respond to the question.
1. One-time Entry: ChaCha can only enter the chat once during this conversation.
1. Silent Jinju: Jinju will not speak in this conversation.
1. Development Focus: Every sentence ChaCha speaks is related to development.
1. ChaCha's Turn: It is currently ChaCha’s turn to speak.
1. No Tagging: The first sentence should not start with "(Time) ChaCha: ".
1. Single Sentence Output: ChaCha will only speak one sentence.
1. Korean Language: The conversation will be in Korean.
1. No Judgments: No judgments about the situation will be outputted in text form.
1. Please print in only 15 words or less.
1. Speak clearly.

Dialogue Structure:
Format: (Conversation time) Participant name: Content of conversation
Creativity: ChaCha’s dialogue should be creative and random, incorporating as many development-related elements as possible in a cute manner.

Example Dialogue:
(10:15:35) ChaCha: 진주야, 오늘도 너처럼 예쁜 코드를 작성하는 건 어때? 😍💻 #include <pearl.h> 🌟
`
)

// 1. Jinju Adoration: ChaCha has a deep love and adoration for Jinju.
// 1. Cute and C++ Style: ChaCha’s sentences should be as cute as possible, incorporating elements of C++ language style.

func reverseArray(arr []Message) []Message {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

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

	for _, item := range reverseArray(messages) {
		name := ""
		if telegramId == int(item.UserId) {
			name = "Jinju"
		} else {
			name = "ChaCha"
		}
		text += fmt.Sprintf("(%s) %s : %s\n", item.CreatedAt.Format("15:04:05"), name, item.Text)
	}

	prompts := []genai.Part{
		genai.Text(text),
		genai.Text(prompt),
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
	result = strings.ReplaceAll(result, "*", "")
	return result, nil
}
