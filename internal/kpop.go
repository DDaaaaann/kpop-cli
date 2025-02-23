package internal

import (
	"bufio"
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/DDaaaaann/kpop-cli/internal/process"
	"io"
	"strconv"
)

func KPOP(port string, forceFlag bool, quietFlag bool, in io.Reader, out io.Writer, executor executor.CommandExecutor) {
	if _, err := strconv.Atoi(port); err != nil {
		fmt.Fprintf(out, "Port '%s' is not well formatted", port)
		return
	}

	pid, err := process.FindProcessUsingPort(port, executor)

	if err != nil {
		if !quietFlag {
			fmt.Fprintf(out, "No process found using port %s\n", port)
		}
		return
	}

	if !forceFlag {
		fmt.Fprintf(out, "Kill process on port %s (PID %d)? (y/n)\n", port, pid)
		scanner := bufio.NewScanner(in)
		scanner.Scan()
		if scanner.Text() != "y" {
			if !quietFlag {
				fmt.Fprintf(out, "Cancelled.")
			}
			return
		}

	}

	killed := process.KillProcess(pid, executor)

	if !killed && !quietFlag {
		fmt.Fprintf(out, "Failed to kill process %d on port %s: %v\n", pid, port, err)
	} else if !quietFlag {
		fmt.Fprintf(out, "Killed process %d on port %s.\n", pid, port)
	}
}
