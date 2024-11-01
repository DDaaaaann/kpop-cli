package executor

import "os/exec"

type CommandExecutor interface {
	Execute(name string, arg ...string) ([]byte, error)
}

type RealCommandExecutor struct{}

func (r *RealCommandExecutor) Execute(name string, arg ...string) ([]byte, error) {
	cmd := exec.Command(name, arg...)
	return cmd.Output()
}

func MockCmdExecutor(output string, err error) CommandExecutor {
	return &MockCommandExecutor{Output: output, Err: err}
}

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
