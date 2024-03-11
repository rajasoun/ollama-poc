package wrapper

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockPromptRunner is a mock type for the PromptRunner interface
type MockPromptRunner struct {
	mock.Mock
}

// Run mocks the Run method
func (m *MockPromptRunner) Run() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

// TestConsolePrompt_GetUserQuestion tests the GetUserQuestion method
func TestConsolePrompt_GetUserQuestion(t *testing.T) {
	mockRunner := new(MockPromptRunner)
	expectedQuestion := "What is the capital of France?"
	mockRunner.On("Run").Return(expectedQuestion, nil)

	consolePrompt := ConsolePrompt{
		Runner: mockRunner,
	}

	question, err := consolePrompt.GetUserQuestion()

	assert.NoError(t, err)
	assert.Equal(t, expectedQuestion, question)
	mockRunner.AssertExpectations(t) // Verify that Run was called as expected
}
