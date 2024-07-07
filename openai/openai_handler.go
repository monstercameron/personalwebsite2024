package openai

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

func Init() error {
	apiKey := os.Getenv("OAIKEY")
	if apiKey == "" {
		return fmt.Errorf("OAIKEY not found in environment variables")
	}
	client = openai.NewClient(apiKey)
	return nil
}

func GenerateCompletion(prompt string) (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("completion error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
}

func GenerateImage(prompt string) (string, error) {
	resp, err := client.CreateImage(
		context.Background(),
		openai.ImageRequest{
			Prompt:         prompt,
			Size:           openai.CreateImageSize256x256,
			ResponseFormat: openai.CreateImageResponseFormatURL,
			N:              1,
		},
	)

	if err != nil {
		return "", fmt.Errorf("image generation error: %v", err)
	}

	return resp.Data[0].URL, nil
}
