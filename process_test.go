package main

import (
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func TestGetPID_Success(t *testing.T) {
	executor := mockCommandExecutor("12345", nil)
	pid := getPID("8080", executor)

	assert.Equal(t, "12345", pid, "Expected PID '12345' for process on port 8080")
}

func TestGetPID_NoProcess(t *testing.T) {
	executor := mockCommandExecutor("", nil)
	pid := getPID("8080", executor)

	assert.Empty(t, pid, "Expected empty PID when no process is found on the specified port")
}

func TestKillPID_Success(t *testing.T) {
	executor := mockCommandExecutor("", nil)
	err := killPID("12345", executor)

	assert.NoError(t, err, "Expected no error for valid PID")
}

func TestKillPID_WithEmptyPID(t *testing.T) {
	executor := mockCommandExecutor("", nil)
	err := killPID("", executor)

	assert.Error(t, err, "Expected error for empty PID")
	assert.EqualError(t, err, "no PID provided", "Unexpected error message for empty PID")
}

func TestKillPID_Failure(t *testing.T) {
	executor := mockCommandExecutor("", exec.ErrNotFound)
	err := killPID("12345", executor)

	assert.Error(t, err, "Expected error due to ErrNotFound")
	assert.EqualError(t, err, "executable file not found in $PATH", "Unexpected error message for empty PID")
}
