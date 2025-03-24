//go:build !windows && e2e
// +build !windows,e2e

package e2e

import (
	"os/exec"
	"strconv"
)

func startTestServer(bindAddress string) (int, int, func()) {
	port, err := getFreePort()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("python", "-m", "http.server", strconv.Itoa(port), "--bind", bindAddress)

	return startServerWithCmd(cmd, bindAddress, port)
}
