package app

import (
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/alex-ello/gpt-cli/internal/config"
	"github.com/alex-ello/gpt-cli/internal/console"
	"github.com/alex-ello/gpt-cli/internal/gpt"
	"github.com/alex-ello/gpt-cli/pkg/utils"
)

func ExecCommand(cfg *config.Config, input string) error {

	client := gpt.NewClient(cfg)
	promptSystem := utils.GeneratePrompt(cfg.GetSystemMessage(), input)
	input = utils.GeneratePrompt(cfg.PromptTemplate, input)

	dialog := client.StartDialog()
	dialog.HistoryAddSystem(promptSystem)

	for {
		if err := handleMessage(cfg, input); err != nil {
			return err
		}

		response, err := getResponse(dialog, cfg, input)
		if err != nil {
			return err
		}
		console.PrintResponse("\n> %s\n\n", response)

		input, err = console.Prompt("Execute? (y/n) or type for a correction: ")
		if err != nil {
			return err
		}

		if stop := executeOrContinue(input, response); stop {
			return nil
		}
	}
}

func handleMessage(cfg *config.Config, input string) error {
	stop, err := cfg.HandleMessage(input)
	if err != nil || stop {
		return err
	}
	return nil
}

func getResponse(dialog *gpt.Dialog, cfg *config.Config, promptUser string) (string, error) {
	// Start the loader
	loader := utils.NewLoader("Loading").Start()

	response, err := dialog.SendMessage(cfg, promptUser)

	// Stop the loader
	loader.Stop()

	if err != nil {
		return "", err
	}

	return response, nil
}

func executeOrContinue(input, response string) bool {
	switch input {
	case "y":
		cmdStr := extractCommand(response)
		return executeShellCommand(cmdStr) == nil
	case "n":
		return true
	}
	return false
}

func extractCommand(response string) string {
	if !strings.Contains(response, "```") {
		return response
	}
	re := regexp.MustCompile("(?s)```(.*?)```")
	matches := re.FindStringSubmatch(response)

	if len(matches) > 1 {
		return strings.Trim(matches[0], "`")
	}

	return ""
}

func InteractiveDialog(cfg *config.Config) error {
	client := gpt.NewClient(cfg)

	dialog := client.StartDialog()

	console.Println("You have entered into an interactive dialog with ChatGPT. Type \"quit\" to exit.\n")
	for {
		input, err := console.Prompt("You > ")
		if err != nil {
			return err
		}

		if err := handleMessage(cfg, input); err != nil {
			return err
		}

		input = strings.TrimSpace(input)
		if input == "exit" || input == "quit" {
			return nil
		}
		if input == "" {
			continue
		}

		response, err := getResponse(dialog, cfg, input)
		if err != nil {
			return err
		}

		console.PrintResponse("GPT: %s\n", response)
	}
}

func executeShellCommand(cmdStr string) error {
	if cmdStr == "" {
		console.Println("No command to execute.")
		return nil
	}
	cmd := exec.Command("sh", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
