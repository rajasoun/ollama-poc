package wrapper

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Prompt defines the interface for getting user input
type Prompt interface {
	GetUserQuestion() (string, error)
}

// ConsolePrompt is a struct that implements the Prompt interface
type ConsolePrompt struct{}

// GetUserQuestion prompts the user to enter a question and returns it.
func (cp ConsolePrompt) GetUserQuestion() (string, error) {
	prompt := promptui.Prompt{
		Label: "What is your question",
	}

	question, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return question, nil
}
