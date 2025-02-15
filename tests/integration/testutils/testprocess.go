package testutils

import (
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"runtime"
	"strconv"
	"strings"
)

func ProcessExists(pid int) (bool, error) {
	realExecutor := &executor.RealCommandExecutor{}
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = realExecutor.Execute("netstat", "-ano")
	} else {
		out, err = realExecutor.Execute("lsof", "-p", strconv.Itoa(pid))
	}
	if err != nil || len(out) == 0 {
		return false, err
	}

	return strings.Contains(string(out), strconv.Itoa(pid)), nil
}
