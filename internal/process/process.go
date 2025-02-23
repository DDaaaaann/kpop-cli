package process

import (
	"errors"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/DDaaaaann/kpop-cli/internal/utils"
)

func FindProcessUsingPort(port string, executor executor.CommandExecutor) (int, error) {
	out, err, format := executor.FindProcessForPort(port)

	if err != nil {
		return 0, err
	}

	if len(out) == 0 {
		return 0, errors.New("no result")
	}

	return utils.ParseFirstPID(out, *format)
}

func KillProcess(pid int, executor executor.CommandExecutor) bool {
	return executor.KillProcess(pid) == nil
}
