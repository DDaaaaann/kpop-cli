//go:build !windows && e2e
// +build !windows,e2e

package e2e

import (
	"os/exec"
	"strconv"
)

func startTestServer() (int, int, func()) {
	port, err := getFreePort()
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("python3", "-m", "http.server", strconv.Itoa(port))

	return startServerWithCmd(cmd, port)
}
