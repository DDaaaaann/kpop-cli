package main

import (
	"os/exec"
	"testing"
)

type MockCommandExecutor struct {
	Output string
	Err    error
}

func (m *MockCommandExecutor) Execute(name string, arg ...string) ([]byte, error) {
	if m.Err != nil {
		return nil, m.Err
	}
	return []byte(m.Output), nil
}

func mockCommandExecutor(output string, err error) CommandExecutor {
	return &MockCommandExecutor{Output: output, Err: err}
}

func TestGetPID_Success(t *testing.T) {
	executor := mockCommandExecutor("12345", nil)
	pid := getPID("8080", executor)
	if pid != "12345" {
		t.Errorf("Expected PID '12345', got '%s'", pid)
	}
}

func TestGetPID_NoProcess(t *testing.T) {
	mockExec := mockCommandExecutor("", nil)
	pid := getPID("8080", mockExec)
	if pid != "" {
		t.Errorf("Expected empty PID, got '%s'", pid)
	}
}

func TestKillPID_Success(t *testing.T) {
	mockExec := mockCommandExecutor("", nil)
	err := killPID("12345", mockExec)
	if err != nil {
		t.Errorf("Expected no error, got '%v'", err)
	}
}

func TestKillPID_Failure(t *testing.T) {
	mockExec := mockCommandExecutor("", exec.ErrNotFound)
	err := killPID("12345", mockExec)
	if err == nil {
		t.Errorf("Expected an error due to process not found, got nil")
	}
}
