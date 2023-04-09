package console

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

// Prompt asks the user for input with a given message and returns the input as a string.
func Prompt(message string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(message)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(input), nil
}

// NoColor disables color output if noColor is true.
func NoColor(noColor bool) {
	color.NoColor = noColor
}

// Println is a convenient helper function to print.
func Println(a ...any) {
	fmt.Println(a...)
}

// Printf is a convenient helper function to print with colors.
func Printf(format string, a ...any) {
	fmt.Printf(format, a...)
}

// PrintResponse prints a formatted response message with color.
func PrintResponse(format string, a ...any) error {
	responseColor := color.New(color.FgYellow, color.Bold)
	_, err := responseColor.Printf(format, a...)
	if err != nil {
		return err
	}
	return nil
}

// PrintrRequest prints a formatted request message with color.
func PrintrRequest(format string, a ...any) error {
	requestColor := color.New(color.FgCyan, color.Bold)
	_, err := requestColor.Printf(format, a...)
	if err != nil {
		return err
	}
	return nil
}

func Error(a ...any) {
	fmt.Println(a...)
}
