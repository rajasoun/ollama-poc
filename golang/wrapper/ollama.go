package wrapper

import (
	"context"
	"fmt"
	"io"

	"github.com/tmc/langchaingo/llms"
)

// OllamaCaller defines an interface for calling the Ollama model.
type OllamaCaller interface {
	Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error)
}

// AskQuestion asks a question to the Ollama model and writes the response to the given writer.
func AskQuestion(ctx context.Context, question string, writer io.Writer, llmCaller OllamaCaller) {
	_, err := llmCaller.Call(ctx, question,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			_, err := writer.Write(chunk)
			return err
		}),
	)
	if err != nil {
		fmt.Fprintln(writer, err) // Use writer for output
	}
}
