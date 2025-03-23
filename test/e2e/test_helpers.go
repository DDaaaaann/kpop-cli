//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func startServerWithCmd(cmd *exec.Cmd, port int) (int, int, func()) {
	resultCmd := make(chan *exec.Cmd)

	go func() {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmdErr := cmd.Start()
		if cmdErr != nil {
			log.Fatal(cmdErr.Error())
		}

		resultCmd <- cmd
	}()

	cmdResult := <-resultCmd
	pid := cmdResult.Process.Pid

	time.Sleep(500 * time.Millisecond)

	return port, pid, func() {
		if err := cmdResult.Process.Kill(); err == nil {
			fmt.Printf("Force-killed PID %d\n", pid)
		} else {
			fmt.Printf("Successfully stopped PID %d gracefully\n", pid)
		}
	}
}

func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}

func getBinaryPath() string {
	binaryName := "kpop-cli"

	switch runtime.GOOS {
	case "linux":
		binaryName += "-linux"
	case "windows":
		binaryName += "-windows"
	case "darwin":
		binaryName += "-darwin"
	default:
		panic(fmt.Sprintf("Unsupported OS: %s", runtime.GOOS))
	}

	switch runtime.GOARCH {
	case "amd64":
		binaryName += "-amd64"
	case "arm64":
		binaryName += "-arm64"
	default:
		panic(fmt.Sprintf("Unsupported architecture: %s", runtime.GOARCH))
	}

	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	return fmt.Sprintf("../../dist/%s", binaryName)
}

func processExists(pid int) (bool, error) {
	var out []byte
	var err error

	if runtime.GOOS == "windows" {
		out, err = exec.Command("tasklist", "/FI", fmt.Sprintf("PID eq %d", pid)).Output()
	} else {
		out, err = exec.Command("ps", "-p", strconv.Itoa(pid)).Output()
	}

	if err != nil || len(out) == 0 {
		return false, err
	}

	return true, nil
}

func getCommand(binaryPath string, port int) *exec.Cmd {
	return exec.Command(binaryPath, "-f", strconv.Itoa(port))
}
