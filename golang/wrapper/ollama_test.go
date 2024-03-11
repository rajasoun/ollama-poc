package wrapper_test

import (
	"bytes"
	"context"
	"testing"

	"github.com/rajasoun/ollama-poc/wrapper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/tmc/langchaingo/llms"
)

// MockOllamaCaller mocks the OllamaCaller interface for testing
type MockOllamaCaller struct {
	mock.Mock
}

func (m *MockOllamaCaller) Call(ctx context.Context, prompt string, options ...llms.CallOption) (string, error) {
	args := m.Called(ctx, prompt, options)

	// Manually invoke the streaming function for testing
	for _, option := range options {
		opt := llms.CallOption(option)
		cfg := &llms.CallOptions{}
		opt(cfg)

		if cfg.StreamingFunc != nil {
			// Simulate streaming data to the writer
			cfg.StreamingFunc(ctx, []byte("Jupiter is the largest planet."))
		}
	}

	return args.String(0), args.Error(1)
}

func TestAskQuestion(t *testing.T) {
	mockOllamaCaller := new(MockOllamaCaller)
	testQuestion := "What is the largest planet?"
	expectedResponse := "Jupiter is the largest planet."

	mockOllamaCaller.On("Call", mock.Anything, testQuestion, mock.Anything).Return(expectedResponse, nil)

	buffer := &bytes.Buffer{} // Using bytes.Buffer to implement io.Writer for capturing output

	// Calling the function under test
	wrapper.AskQuestion(context.Background(), testQuestion, buffer, mockOllamaCaller)

	// Asserting the function's behavior
	assert.Contains(t, buffer.String(), expectedResponse, "The response should contain the expected output")
	mockOllamaCaller.AssertExpectations(t) // Ensure that the mock was called as expected
}
