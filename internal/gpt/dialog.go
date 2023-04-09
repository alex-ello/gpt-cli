package gpt

import (
    "github.com/alex-ello/gpt-cli/internal/config"
    "github.com/sashabaranov/go-openai"
)

type Dialog struct {
	History []openai.ChatCompletionMessage
	Client  *Client
}

func (c *Client) StartDialog() *Dialog {
	return &Dialog{
		History: []openai.ChatCompletionMessage{},
		Client:  c,
	}
}

func (dialog *Dialog) SendMessage(cfg *config.Config, message string) (string, error) {

	dialog.HistoryAddUser(message)

	response, err := dialog.Client.SendRequest(cfg, dialog.History)
	if err != nil {
		return "", err
	}

	dialog.HistoryAddAssistant(response)
	return response, nil
}

func (dialog *Dialog) HistoryAddSystem(message string) *Dialog {
	return dialog.historyAdd(openai.ChatMessageRoleSystem, message)
}

func (dialog *Dialog) HistoryAddUser(message string) *Dialog {
	return dialog.historyAdd(openai.ChatMessageRoleUser, message)
}

func (dialog *Dialog) HistoryAddAssistant(message string) *Dialog {
	return dialog.historyAdd(openai.ChatMessageRoleAssistant, message)
}

func (dialog *Dialog) historyAdd(role string, content string) *Dialog {
	assistantMessage := openai.ChatCompletionMessage{
		Role:    role,
		Content: content,
	}
	dialog.History = append(dialog.History, assistantMessage)

	return dialog
}
