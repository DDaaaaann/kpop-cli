package main

import (
	"bytes"
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal"
	"github.com/DDaaaaann/kpop-cli/tests/integration/testutils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestKpopCLI_success(t *testing.T) {
	var stdin, stdout bytes.Buffer
	port, pid, kill := testutils.StartTestServer()

	stdin.WriteString("y\n")

	exists, _ := testutils.ProcessExists(pid)
	assert.True(t, exists, fmt.Sprintf("No process with pid %d exists", pid))

	internal.KPOP(strconv.Itoa(port), false, false, &stdin, &stdout, nil)

	expected := fmt.Sprintf("Kill process on port %d (PID %d)? (y/n)", port, pid)
	assert.Contains(t, stdout.String(), expected, "Output does not contain confirmation message")

	processExists, _ := testutils.ProcessExists(pid)
	assert.False(t, processExists, fmt.Sprintf("Process with pid %d should have been killed", pid))

	defer kill()
}
