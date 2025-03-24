//go:build windows && e2e
// +build windows,e2e

package e2e

import (
	"os/exec"
	"strconv"
	"syscall"
)

func startTestServer() (int, int, func()) {
	port, err := getFreePort()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("python", "-m", "http.server", "--bind", "127.0.0.1", strconv.Itoa(port))
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: syscall.CREATE_NEW_PROCESS_GROUP,
	}

	return startServerWithCmd(cmd, port)
}
