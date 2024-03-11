package cmd

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/rajasoun/ollama-poc/input"
	"github.com/spf13/cobra"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
)

// This writer will be used as the default output, which can be overridden for testing.
var outputWriter io.Writer = os.Stdout

// Update the Run function of your Cobra command to use the outputWriter.
var askCmd = &cobra.Command{
	Use:   "ask",
	Short: "Ask a question to the Ollama model",
	Run: func(cmd *cobra.Command, args []string) {
		// Use the input package to get the question from the user
		question, err := input.GetUserQuestion()
		if err != nil {
			fmt.Fprintf(cmd.OutOrStdout(), "Error: %v\n", err)
			return
		}

		fullQuestion := fmt.Sprintf("Human: %s\nAssistant:", question)
		askQuestion(context.Background(), fullQuestion, cmd.OutOrStdout())
	},
}

// askQuestion asks a question to the Ollama model and writes the response to the given writer.
func askQuestion(ctx context.Context, question string, writer io.Writer) {
	llm, err := ollama.New(ollama.WithModel("llama2"))
	if err != nil {
		fmt.Fprintln(writer, err) // Use writer for output
		return
	}
	_, err = llm.Call(ctx, question,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			_, err := writer.Write(chunk) // Use writer for output
			return err
		}),
	)
	if err != nil {
		fmt.Fprintln(writer, err) // Use writer for output
	}
}

// Add the askCmd to the root command.
func init() {
	rootCmd.AddCommand(askCmd)
}
