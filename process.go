package main

import (
	"fmt"
	"runtime"
	"strings"
)

// getPID fetches the process ID for a given port by executing the appropriate command.
func getPID(port string, executor CommandExecutor) string {
	var out []byte
	var err error
	if runtime.GOOS == "windows" {
		out, err = executor.Execute("netstat", "-ano")
	} else {
		out, err = executor.Execute("lsof", "-t", "-i:"+port)
	}

	if err != nil || len(out) == 0 {
		return ""
	}
	return strings.TrimSpace(string(out))
}

// killPID terminates the process with the given PID.
func killPID(pid string, executor CommandExecutor) error {
	var err error

	if pid == "" {
		return fmt.Errorf("no PID provided")
	}

	if runtime.GOOS == "windows" {
		_, err = executor.Execute("taskkill", "/PID", pid, "/F")
	} else {
		_, err = executor.Execute("kill", "-9", pid)
	}
	return err
}
