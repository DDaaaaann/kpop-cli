package unit_test

import (
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/DDaaaaann/kpop-cli/internal/process"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestGetPID_Success(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor("12345", nil)
	pid := process.GetPID("8080", mockExecutor)

	assert.Equal(t, "12345", pid, "Expected PID '12345' for process on port 8080")
}

func TestGetPID_NoProcess(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor("", nil)
	pid := process.GetPID("8080", mockExecutor)

	assert.Empty(t, pid, "Expected empty PID when no process is found on the specified port")
}

func TestKillPID_Success(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor("", nil)
	err := process.KillPID("12345", mockExecutor)

	assert.NoError(t, err, "Expected no error for valid PID")
}

func TestKillPID_WithEmptyPID(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor("", nil)
	err := process.KillPID("", mockExecutor)

	assert.Error(t, err, "Expected error for empty PID")
	assert.EqualError(t, err, "no PID provided", "Unexpected error message for empty PID")
}

func TestKillPID_Failure(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor("", exec.ErrNotFound)
	err := process.KillPID("12345", mockExecutor)

	assert.Error(t, err, "Expected error due to ErrNotFound")
	assert.EqualError(t, err, "executable file not found in $PATH", "Unexpected error message for empty PID")
}
