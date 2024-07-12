// openai/openai_handler.go

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	openai "github.com/sashabaranov/go-openai"
)

var client *openai.Client

type QuoteOfTheDay struct {
	Date  string `json:"date"`
	Quote string `json:"quote"`
}

func Init() error {
	apiKey := os.Getenv("OAIKEY")
	if apiKey == "" {
		return fmt.Errorf("OAIKEY not found in environment variables")
	}
	client = openai.NewClient(apiKey)
	return nil
}

func GetQuoteOfTheDay() (string, error) {
	today := time.Now().Format("2006-01-02")
	cacheDir := "cache"
	cacheFile := filepath.Join(cacheDir, "qotd.json")

	// Ensure cache directory exists
	if err := os.MkdirAll(cacheDir, 0755); err != nil {
		return "", fmt.Errorf("failed to create cache directory: %v", err)
	}

	// Check if we have a cached quote for today
	if quote, err := getCachedQuote(cacheFile, today); err == nil {
		return quote, nil
	}

	// Generate new quote
	quote, err := generateQuote()
	if err != nil {
		return "", err
	}

	// Cache the new quote
	if err := cacheQuote(cacheFile, today, quote); err != nil {
		fmt.Printf("Warning: Failed to cache quote: %v\n", err)
	}

	return quote, nil
}

func getCachedQuote(cacheFile, today string) (string, error) {
	data, err := os.ReadFile(cacheFile)
	if err != nil {
		return "", err
	}

	var qotd QuoteOfTheDay
	if err := json.Unmarshal(data, &qotd); err != nil {
		return "", err
	}

	if qotd.Date == today {
		return qotd.Quote, nil
	}

	return "", fmt.Errorf("no quote for today")
}

func cacheQuote(cacheFile, date, quote string) error {
	qotd := QuoteOfTheDay{
		Date:  date,
		Quote: quote,
	}

	data, err := json.Marshal(qotd)
	if err != nil {
		return err
	}

	return os.WriteFile(cacheFile, data, 0644)
}

func generateQuote() (string, error) {
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Generate an inspirational quote from a famous person. Quote and cite the source.",
				},
			},
		},
	)

	if err != nil {
		return "", fmt.Errorf("quote generation error: %v", err)
	}

	return resp.Choices[0].Message.Content, nil
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
