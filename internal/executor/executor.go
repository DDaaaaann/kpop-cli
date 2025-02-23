package executor

import (
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal/utils"
	"os/exec"
	"runtime"
	"strconv"
)

type CommandExecutor interface {
	FindProcessForPort(port string) ([]byte, error, *utils.ProcessOutputFormat)
	KillProcess(pid int) error
}

type RealCommandExecutor struct{}

func (r *RealCommandExecutor) KillProcess(pid int) error {
	var err error

	switch runtime.GOOS {
	case "darwin", "linux":
		_, err = exec.Command("kill", "-9", strconv.Itoa(pid)).Output()
	case "windows":
		_, err = exec.Command("taskkill", "/PID", strconv.Itoa(pid), "/F").Output()
	default:
		panic(fmt.Sprintf("Unsupported OS: %s", runtime.GOOS))
	}

	return err
}

func (r *RealCommandExecutor) FindProcessForPort(port string) ([]byte, error, *utils.ProcessOutputFormat) {
	switch runtime.GOOS {
	case "darwin", "linux":
		output, err := exec.Command("lsof", "-t", "-i:"+port).Output()
		format := utils.FormatPIDOnly
		return output, err, &format
	case "windows":
		output, err := exec.Command("netstat", "-ano").Output()
		format := utils.FormatNetstat
		return output, err, &format
	default:
		panic(fmt.Sprintf("Unsupported OS: %s", runtime.GOOS))
	}
}

func MockCmdExecutor(output []byte, err error, format utils.ProcessOutputFormat) CommandExecutor {
	return &MockCommandExecutor{Output: output, Err: err, Format: &format}
}

type MockCommandExecutor struct {
	Output []byte
	Err    error
	Format *utils.ProcessOutputFormat
}

func (m *MockCommandExecutor) FindProcessForPort(port string) ([]byte, error, *utils.ProcessOutputFormat) {
	if m.Err != nil {
		return nil, m.Err, nil
	}
	return m.Output, nil, m.Format
}

func (m *MockCommandExecutor) KillProcess(pid int) error {
	return m.Err
}
