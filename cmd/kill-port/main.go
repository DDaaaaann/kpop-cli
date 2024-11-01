package main

import (
	"flag"
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/DDaaaaann/kpop-cli/internal/process"
	"os"
)

var forceFlag = flag.Bool("f", false, "Force kill without confirmation")
var quietFlag = flag.Bool("q", false, "Quiet mode, suppress output")

func usage() {
	fmt.Println("Usage: kill-port [-f] [-q] <port>")
	fmt.Println("Options:")
	fmt.Println("  -f    Force kill without confirmation")
	fmt.Println("  -q    Quiet mode, suppress output")
	fmt.Println("  -h    Show this help message")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Error: No port specified.")
		usage()
		os.Exit(1)
	}
	port := flag.Arg(0)

	realExecutor := &executor.RealCommandExecutor{}

	pid := process.GetPID(port, realExecutor)
	if pid == "" {
		if !*quietFlag {
			fmt.Println("No process found using port", port)
		}
		return
	}

	if !*forceFlag {
		var confirm string
		fmt.Printf("Kill process using port %s (PID %s)? (y/n) ", port, pid)
		fmt.Scanln(&confirm)
		if confirm != "y" {
			if !*quietFlag {
				fmt.Println("Cancelled.")
			}
			return
		}
	}

	err := process.KillPID(pid, realExecutor)
	if err != nil && !*quietFlag {
		fmt.Printf("Failed to kill process %s on port %s: %v\n", pid, port, err)
	} else if !*quietFlag {
		fmt.Printf("Killed process %s on port %s.\n", pid, port)
	}
}
