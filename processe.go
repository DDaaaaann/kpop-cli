package main

import (
	"os/exec"
	"runtime"
	"strings"
)

func getPID(port string) string {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("netstat", "-ano", "|", "findstr", ":"+port)
	} else {
		cmd = exec.Command("lsof", "-t", "-i:"+port)
	}

	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return ""
	}
	return strings.TrimSpace(string(out))
}

func killPID(pid string) error {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("taskkill", "/PID", pid, "/F")
	} else {
		cmd = exec.Command("kill", "-9", pid)
	}
	return cmd.Run()
}
