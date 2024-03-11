package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ollama-pos",
	Short: "ollama-pos is a tool for interacting with the Ollama model.",
	Long:  `ollama-pos is a tool for interacting with the Ollama model. It can be used to ask questions to the model and get responses.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
