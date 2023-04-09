package main

import (
	"fmt"
    "github.com/alex-ello/gpt-cli/internal/console"
    "os"
    "strings"

    "github.com/alex-ello/gpt-cli/internal/app"
	"github.com/alex-ello/gpt-cli/internal/config"
)

const configName = "config.toml"
const appName = "gpt-cli"

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	path := config.GetConfigFilePath(appName, configName)
    cfg := config.NewConfig(path)
	err := cfg.LoadConfig()
	if err != nil {
		return err
	}

	console.NoColor(!cfg.Color)
	if len(os.Args) > 1 {
		return app.ExecCommand(cfg, strings.Join(os.Args[1:], " "))
	}

	return app.InteractiveDialog(cfg)
}
