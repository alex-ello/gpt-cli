package config

import (
	"fmt"
	"github.com/alex-ello/gpt-cli/internal/console"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	configPath         string
	Version            string
	Debug              bool    `toml:"debug,omitempty"`
	MaxTokens          int     `toml:"max_tokens,omitempty"`
	Temperature        float32 `toml:"temperature,omitempty"`
	PromptTemplate     string  `toml:"prompt_template,omitempty"`
	Model              string  `toml:"model,omitempty"`
	APIKey             string  `toml:"api_key,omitempty"`
	Color              bool    `toml:"color,omitempty"`
	SystemMessage      string  `toml:"system_message,omitempty"`
	SystemMessageDebug string  `toml:"system_message_debug,omitempty"`
}

func NewConfig(configPath string) *Config {
	return &Config{
		configPath: configPath,
	}
}

func (c *Config) LoadConfig() error {

	_, err := os.Stat(c.configPath)
	if os.IsNotExist(err) {
		console.Println("Configuration file ", c.configPath, "not found. Creating a new one.")
		return c.createDefaultConfig()
	}

	if _, err := toml.DecodeFile(c.configPath, c); err != nil {
		return err
	}

	c.applyDefaultValues()

	return nil
}

func (c *Config) createDefaultConfig() error {
	c.applyDefaultValues()

	err := c.setAPIKeyDialog()
	if err != nil {
		return err
	}

	err = c.SaveConfig()
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) SaveConfig() error {

	dir := filepath.Dir(c.configPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	file, err := os.Create(c.configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := toml.NewEncoder(file)
	if err := encoder.Encode(c); err != nil {
		return err
	}

	console.Println("Configuration saved successfully: ", c.configPath)
	console.Println("Type \"config\" to start the configuration dialog.")

	return nil
}

func (c *Config) ConfigDialog() error {

	err := c.setAPIKeyDialog()
	if err != nil {
		return fmt.Errorf("error setting API key: %w", err)
	}

	err = c.SaveConfig()
	if err != nil {
		return fmt.Errorf("error saving config: %w", err)
	}
	return nil
}

func (c *Config) setAPIKeyDialog() error {
	apiKey, err := console.Prompt("Enter your OpenAI GPT-3 API key: ")
	if err != nil {
		return fmt.Errorf("error reading API key: %w", err)
	}
	c.APIKey = apiKey

	return nil
}

func getConfigPathFromEnvVar(appName string) (string, bool) {
	envVarName := fmt.Sprintf("%s_CONFIG_PATH", strings.ToUpper(appName))
	configPath := os.Getenv(envVarName)
	return configPath, configPath != ""
}

func getConfigPathFromCWD(configName string) (string, bool) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", false
	}

	configPath := filepath.Join(cwd, configName)
	if _, err := os.Stat(configPath); err == nil {
		return configPath, true
	}

	return configPath, false
}

func getConfigPathFromHomeDir(appName, configName string) (string, bool) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", false
	}

	configPath := filepath.Join(homeDir, fmt.Sprintf(".%s", appName), configName)
	if _, err := os.Stat(configPath); err == nil {
		return configPath, true
	}

	return configPath, false
}

func GetConfigFilePath(appName, configName string) string {
	if configPath, found := getConfigPathFromEnvVar(appName); found {
		return configPath
	}

	if configPath, found := getConfigPathFromCWD(configName); found {
		return configPath
	}

	configPath, _ := getConfigPathFromHomeDir(appName, configName)

	return configPath
}

func (c *Config) GetSystemMessage() string {
	if c.Debug {
		return c.SystemMessageDebug
	}
	return c.SystemMessage
}
