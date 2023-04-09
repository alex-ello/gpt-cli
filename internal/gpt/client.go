package gpt

import (
	"context"
	"github.com/alex-ello/gpt-cli/pkg/utils"

	"github.com/alex-ello/gpt-cli/internal/config"
	"github.com/sashabaranov/go-openai"
)

type Client struct {
	Client *openai.Client
}

func NewClient(cfg *config.Config) *Client {
	return &Client{
		Client: openai.NewClient(cfg.APIKey),
	}
}

func (c *Client) SendRequest(cfg *config.Config, messages []openai.ChatCompletionMessage) (string, error) {
	request := c.buildRequest(cfg, messages)

	if cfg.Debug {
		utils.PrintStruct(request)
	}

	completion, err := c.Client.CreateChatCompletion(context.Background(), *request)
	if err != nil {
		return "", err
	}

	if cfg.Debug {
		utils.PrintStruct(completion)
	}

	return completion.Choices[0].Message.Content, nil
}

func (c *Client) buildRequest(cfg *config.Config, messages []openai.ChatCompletionMessage) *openai.ChatCompletionRequest {
	return &openai.ChatCompletionRequest{
		Model:       cfg.Model,
		Temperature: cfg.Temperature,
		MaxTokens:   cfg.MaxTokens,
		Messages:    messages,
	}
}
