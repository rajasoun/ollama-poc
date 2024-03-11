package input

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// GetUserQuestion prompts the user to enter a question and returns it.
func GetUserQuestion() (string, error) {
	prompt := promptui.Prompt{
		Label: "What is your question",
	}

	question, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return question, nil
}
