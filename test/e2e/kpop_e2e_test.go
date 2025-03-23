//go:build e2e
// +build e2e

package e2e

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestKpopCLI_success(t *testing.T) {
	var stdout bytes.Buffer

	port, pid, kill := startTestServer()
	log.Printf("Started test server on port %d and pid %d.\n", port, pid)

	exists, _ := processExists(pid)
	assert.True(t, exists, fmt.Sprintf("No process with pid %d exists", pid))

	binaryPath := getBinaryPath()
	log.Printf("Running e2e-test for binary '%s'", binaryPath)

	kpopCmd := execCommand(binaryPath, port)
	kpopCmd.Stdout = &stdout

	err := kpopCmd.Run()
	assert.NoError(t, err, "Failed to execute kpop command")

	expected := fmt.Sprintf("Killed process %d on port %d", pid, port)
	assert.Contains(t, stdout.String(), expected, "Output does not contain confirmation message")

	exists, _ = processExists(pid)
	assert.False(t, exists, fmt.Sprintf("Process with pid %d should have been killed", pid))

	defer kill()
}
