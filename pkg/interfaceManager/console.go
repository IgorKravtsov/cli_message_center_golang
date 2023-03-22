package interfaceManager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func printBoldOption(text string) {
	fmt.Printf("\x1b[1m> %s\x1b[0m\n", text)
}

func printRegularOption(text string) {
	fmt.Printf("  %s\n", text)
}

func Clear() {
	fmt.Print("\033[2J")
	fmt.Print("\033[H")
}

func PrintOptions(options []string, selectedIndex int) {
	for i, option := range options {
		if i == selectedIndex {
			printBoldOption(option)
		} else {
			printRegularOption(option)
		}
	}
}

func SetText(text string) {
	fmt.Print(text)
}

func GetEnteredText() (string, error) {
	message := ""
	reader := bufio.NewReader(os.Stdin)
	message = ""
	for {
		input, _, err := reader.ReadRune()
		if err != nil {
			return "", errors.New(fmt.Sprintf("An error occurred while reading input: %v", err))
		}
		if input == '\n' {
			break
		}
		message += string(input)
	}
	return strings.TrimSpace(message), nil
}
