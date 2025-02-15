package testutils

import (
	"fmt"
	"log"
	"net"
	"os/exec"
	"strconv"
	"time"
)

// getFreePort dynamically finds an available port
func getFreePort() (int, error) {
	listener, err := net.Listen("tcp", ":0") // Let OS pick a free port
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

// StartTestServer runs a simple HTTP server on a free port
func StartTestServer() (int, int, func()) {
	resultCmd := make(chan *exec.Cmd)
	port, err := getFreePort()

	if err != nil {
		panic(fmt.Sprintf("Failed to find free port: %v", err))
	}

	go func() {
		cmd := exec.Command("python3", "-m", "http.server", strconv.Itoa(port))

		cmdErr := cmd.Start()
		if cmdErr != nil {
			log.Fatal(cmdErr)
		}

		resultCmd <- cmd // Send the result to the channel
	}()

	cmd := <-resultCmd
	pid := cmd.Process.Pid

	log.Printf("Just started subprocess %d.\n", pid)

	time.Sleep(500 * time.Millisecond) // Allow server to start

	return port, pid, func() {
		err := cmd.Process.Kill()
		if err != nil {
			fmt.Printf("Killing subprocess with pid: %d", pid)
			return
		}
	}
}
