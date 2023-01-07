package bot

import (
	"context"
	gogpt "github.com/sashabaranov/go-gpt3"
	"log"
)

func connectAI(response string) string {
	c := gogpt.NewClient("YOUR_TOKEN")
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     "text-davinci-003",
		MaxTokens: 1000,
		Prompt:    response,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		log.Printf("Ошибка ИИ : %v", err)
	}
	return resp.Choices[0].Text
}
