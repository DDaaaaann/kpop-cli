package internal

import (
	"bufio"
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/DDaaaaann/kpop-cli/internal/process"
	"io"
)

func KPOP(port string, forceFlag bool, quietFlag bool, in io.Reader, out io.Writer, executor *executor.RealCommandExecutor) {
	pid := process.FindProcessUsingPort(port, executor)

	if pid == "" {
		if !quietFlag {
			fmt.Fprintf(out, "No process found using port", port)
		}
		return
	}

	if !forceFlag {
		fmt.Fprintf(out, "Kill process on port %s (PID %s)? (y/n) \n", port, pid)
		scanner := bufio.NewScanner(in)
		scanner.Scan()
		if scanner.Text() != "y" {
			if !quietFlag {
				fmt.Fprintf(out, "Cancelled.")
			}
			return
		}

	}

	err := process.KillProcess(pid, executor)
	if err != nil && !quietFlag {
		fmt.Fprintf(out, "Failed to kill process %s on port %s: %v\n", pid, port, err)
	} else if !quietFlag {
		fmt.Fprintf(out, "Killed process %s on port %s.\n", pid, port)
	}
}
