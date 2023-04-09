package utils

import (
    "os"
    "runtime"
    "strings"
)

func GeneratePrompt(promptTemplate string, input string) string {
	operatingSystem := runtime.GOOS
	shell := os.Getenv("SHELL")
	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = "unknown"
	}

	fileList, err := GetFileList(workingDir, ", ", 100)
	if err != nil {
		fileList = "unknown"
	}

	prompt := promptTemplate
	prompt = strings.Replace(prompt, "{shell}", shell, 1)
	prompt = strings.Replace(prompt, "{input}", input, 1)
	prompt = strings.Replace(prompt, "{os}", operatingSystem, 1)
	prompt = strings.Replace(prompt, "{path}", workingDir, 1)
	prompt = strings.Replace(prompt, "{files}", fileList, 1)

	return prompt
}
