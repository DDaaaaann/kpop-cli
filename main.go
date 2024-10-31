package main

import (
	"flag"
	"fmt"
	"os"
)

// Flags
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

	// Check for port argument or recent error in command history
	port := ""
	if len(flag.Args()) > 0 {
		port = flag.Arg(0)
	} else {
		fmt.Println("Error: No port specified and no recent port error found.")
		usage()
		os.Exit(1)
	}

	pid := getPID(port)
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

	err := killPID(pid)
	if err != nil && !*quietFlag {
		fmt.Printf("Failed to kill process %s on port %s: %v\n", pid, port, err)
	} else if !*quietFlag {
		fmt.Printf("Killed process %s on port %s.\n", pid, port)
	}
}
