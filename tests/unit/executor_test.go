package unit_test

import (
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestMockExecutor_Success(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor("12345", nil)
	output, err := mockExecutor.Execute("dummy command")

	assert.NoError(t, err)
	assert.Equal(t, "12345", string(output))
}

func TestMockExecutor_Error(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor("", exec.ErrNotFound)
	_, err := mockExecutor.Execute("dummy command")

	assert.Error(t, err)
	assert.EqualError(t, err, "executable file not found in $PATH")
}
