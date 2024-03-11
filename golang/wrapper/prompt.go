package wrapper

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

// Prompt defines the interface for getting user input
type Prompt interface {
	GetUserQuestion() (string, error)
}

// PromptRunner defines the interface for running a prompt and getting a response
type PromptRunner interface {
	Run() (string, error)
}

// ConsolePrompt is a struct that implements the Prompt interface
// It depends on something that can run a prompt, rather than directly on promptui
type ConsolePrompt struct {
	Runner PromptRunner
}

// GetUserQuestion prompts the user to enter a question and returns it.
func (cp ConsolePrompt) GetUserQuestion() (string, error) {
	question, err := cp.Runner.Run()
	if err != nil {
		return "", fmt.Errorf("prompt failed: %w", err)
	}

	return question, nil
}

// RealPromptRunner uses promptui to run real prompts
type RealPromptRunner struct{}

func (RealPromptRunner) Run() (string, error) {
	prompt := promptui.Prompt{
		Label: "What is your question",
	}

	return prompt.Run()
}
