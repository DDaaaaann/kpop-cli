//go:build e2e

package internal

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"net"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

const BindAddress = "127.0.0.1"

func TestKpopCLI_success(t *testing.T) {
	var stdout bytes.Buffer

	port, pid, kill := startTestServer(BindAddress)
	log.Printf("Started test server on port %d and pid %d.\n", port, pid)

	exists, _ := processExists(pid)
	assert.True(t, exists, fmt.Sprintf("No process with pid %d exists", pid))

	binaryPath := getBinaryPath()
	log.Printf("Running e2e-test for binary '%s'", binaryPath)

	kpopCmd := exec.Command(binaryPath, "-f", strconv.Itoa(port))
	kpopCmd.Stdout = &stdout

	err := kpopCmd.Run()
	assert.NoError(t, err, "Failed to execute kpop command")

	expected := fmt.Sprintf("Killed process %d on port %d", pid, port)
	assert.Contains(t, stdout.String(), expected, "Output does not contain confirmation message")

	processExists, _ := processExists(pid)
	assert.False(t, processExists, fmt.Sprintf("Process with pid %d should have been killed", pid))

	defer kill()
}

func processExists(pid int) (bool, error) {
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = exec.Command("netstat", "-ano").Output()
	} else {
		out, err = exec.Command("lsof", "-p", strconv.Itoa(pid)).Output()
	}
	if err != nil || len(out) == 0 {
		return false, err
	}

	return strings.Contains(string(out), strconv.Itoa(pid)), nil
}

// getFreePort dynamically finds an available port
func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0") // Let OS pick a free port
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

// startTestServer runs a simple HTTP server on a free port
func startTestServer(bindAddress string) (int, int, func()) {
	resultCmd := make(chan *exec.Cmd)
	port, err := getFreePort()

	if err != nil {
		panic(fmt.Sprintf("Failed to find free port: %v", err))
	}

	go func() {
		var cmd *exec.Cmd
		if runtime.GOOS == "windows" {
			cmd = exec.Command("python", "-m", "http.server", strconv.Itoa(port), "--bind", bindAddress)
		} else {
			cmd = exec.Command("python3", "-m", "http.server", strconv.Itoa(port), "--bind", bindAddress)
		}
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		cmdErr := cmd.Start()
		if cmdErr != nil {
			log.Fatal(cmdErr.Error())
		}

		resultCmd <- cmd // Send the result to the channel
	}()

	cmd := <-resultCmd
	pid := cmd.Process.Pid

	if !waitForPort(bindAddress, port, 10*time.Second) {
		log.Fatal("Server never came up!")
	}

	return port, pid, func() {
		if err := cmd.Process.Kill(); err == nil {
			fmt.Printf("Force-killed PID %d\n", pid)
		} else {
			fmt.Printf("Successfully stopped PID %d gracefully\n", pid)
		}
	}
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

	return fmt.Sprintf("../dist/%s", binaryName)
}

func waitForPort(bindAddress string, port int, timeout time.Duration) bool {
	address := fmt.Sprintf("%s:%d", bindAddress, port)
	deadline := time.Now().Add(timeout)

	for time.Now().Before(deadline) {
		fmt.Printf("Waiting for port %d...\n", port)
		conn, err := net.DialTimeout("tcp", address, time.Second)
		if err == nil {
			conn.Close()
			return true
		}
		time.Sleep(500 * time.Millisecond)
	}
	return false
}
