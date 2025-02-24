package process_test

import (
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/DDaaaaann/kpop-cli/internal/process"
	"github.com/DDaaaaann/kpop-cli/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindProcess_Error(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor(nil, fmt.Errorf("error"), utils.FormatPIDOnly)

	port, err := process.FindProcessUsingPort("8080", mockExecutor)
	assert.Error(t, err)
	assert.Equal(t, "error", err.Error())
	assert.Equal(t, port, 0)
}

func TestFindProcess_No_Result(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor([]byte(""), nil, utils.FormatPIDOnly)

	port, err := process.FindProcessUsingPort("8080", mockExecutor)
	assert.Error(t, err)
	assert.Equal(t, "no result", err.Error())
	assert.Equal(t, port, 0)
}

func TestFindProcess_Succes(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor([]byte("9999"), nil, utils.FormatPIDOnly)

	port, err := process.FindProcessUsingPort("8080", mockExecutor)
	assert.NoError(t, err)
	assert.Equal(t, port, 9999)
}

func TestKillProcess_Succes(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor(nil, nil, utils.FormatPIDOnly)
	killed := process.KillProcess(9999, mockExecutor)
	assert.True(t, killed)
}

func TestKillProcess_Failure(t *testing.T) {
	mockExecutor := executor.MockCmdExecutor(nil, fmt.Errorf("an error occured"), utils.FormatPIDOnly)
	killed := process.KillProcess(9999, mockExecutor)
	assert.False(t, killed)
}
