package process

import (
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"runtime"
	"strings"
)

// FindProcessUsingPort fetches the process ID for a given port by executing the appropriate command.
func FindProcessUsingPort(port string, executor executor.CommandExecutor) string {
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = executor.Execute("netstat", "-ano")
	} else {
		out, err = executor.Execute("lsof", "-t", "-i:"+port)
	}

	if err != nil || len(out) == 0 {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// KillProcess terminates the process with the given PID.
func KillProcess(pid string, executor executor.CommandExecutor) error {
	var err error
	if pid == "" {
		return fmt.Errorf("no PID provided")
	}

	if runtime.GOOS == "windows" {
		_, err = executor.Execute("taskkill", "/PID", pid, "/F")
	} else {
		_, err = executor.Execute("kill", "-9", pid)
	}
	return err
}
