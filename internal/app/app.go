package app

import (
	"os"
	"os/exec"
	"strings"

	"github.com/alex-ello/gpt-cli/internal/config"
	"github.com/alex-ello/gpt-cli/internal/console"
	"github.com/alex-ello/gpt-cli/internal/gpt"
	"github.com/alex-ello/gpt-cli/pkg/utils"
)

func ExecCommand(cfg *config.Config, input string) error {

	client := gpt.NewClient(cfg)
	promptSystem := utils.GeneratePrompt(cfg.GetSystemMessage(), input)

	dialog := client.StartDialog()
	dialog.HistoryAddSystem(promptSystem)

	stop, err := cfg.HandleMessage(input)
	if err != nil {
		return err
	}
	if stop {
		return nil
	}

	// Start the loader
	loader := utils.NewLoader("Loading").Start()

	promptUser := utils.GeneratePrompt(cfg.PromptTemplate, input)
	response, err := dialog.SendMessage(cfg, promptUser)

	// Stop the loader
	loader.Stop()

	if err != nil {
		return err
	}

	cmdStr := strings.TrimSpace(response)
	if cmdStr == "" {
		console.Println("No command to execute.")
		return nil
	}
	console.Printf("\n> %s\n\n", cmdStr)

	if cfg.Debug {
		return nil
	}

	confirm, err := console.Prompt("Execute? (y/n): ")
	if err != nil {
		return err
	}

	if confirm == "y" || confirm == "Y" {
		return executeShellCommand(cmdStr)
	}

	return nil
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

		stop, err := cfg.HandleMessage(input)
		if err != nil {
			return err
		}
		if stop {
			return nil
		}

		input = strings.TrimSpace(input)
		if input == "exit" || input == "quit" {
			break
		}
		if input == "" {
			continue
		}

		// Start the loader
		loader := utils.NewLoader("Loading")
		loader.Start()

		response, err := dialog.SendMessage(cfg, input)

		// Stop the loader
		loader.Stop()

		if err != nil {
			return err
		}

		console.PrintResponse("GPT: %s\n", response)
	}

	return nil
}

func executeShellCommand(cmdStr string) error {
	strings.Trim(cmdStr, "`")
	cmd := exec.Command("bash", "-c", cmdStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}