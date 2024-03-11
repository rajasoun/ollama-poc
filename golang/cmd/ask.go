package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/rajasoun/ollama-poc/wrapper"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms/ollama"
)

// outputWriter is used as the default output, which can be overridden for testing.
var outputWriter io.Writer = os.Stdout

var prompt wrapper.Prompt = wrapper.ConsolePrompt{
	Runner: wrapper.RealPromptRunner{},
}

// Update the Run function of your Cobra command to use the outputWriter.
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask a question to the Ollama model",
	Run: func(cmd *cobra.Command, args []string) {
		// Use the input package to get the question from the user
		question, err := prompt.GetUserQuestion()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Error: %v\n", err)
			return
		}

		fullQuestion := fmt.Sprintf("Human: %s\nAssistant:", question)
		// Initialize the default OllamaCaller
		llmCaller, err := ollama.New(ollama.WithModel("llama2"))
		wrapper.AskQuestion(context.Background(), fullQuestion, cmd.OutOrStdout(), llmCaller)
	},
}

// Add the askCmd to the root command.
func init() {
	rootCmd.AddCommand(askCmd)
}
