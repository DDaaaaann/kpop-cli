package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
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
		port = findRecentPortError()
		if port == "" {
			fmt.Println("Error: No port specified and no recent port error found.")
			usage()
			os.Exit(1)
		} else {
			fmt.Println("Using detected port:", port)
		}
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

func findRecentPortError() string {
	// Determine the shell and history file path
	shell := os.Getenv("SHELL")
	var historyFilePath string

	if strings.Contains(shell, "bash") {
		historyFilePath = os.Getenv("HOME") + "/.bash_history"
	} else if strings.Contains(shell, "zsh") {
		historyFilePath = os.Getenv("HOME") + "/.zsh_history"
	} else {
		log.Println("Unsupported shell for history detection.")
		return ""
	}

	// Open the history file
	historyFile, err := os.Open(historyFilePath)
	if err != nil {
		log.Fatalf("Error opening history file: %v\n", err)
	}
	defer historyFile.Close()

	scanner := bufio.NewScanner(historyFile)
	var lastPortError string
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "address already in use") || strings.Contains(line, "port") {
			lastPortError = line
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return extractPort(lastPortError)
}

func extractPort(errorLine string) string {
	parts := strings.Fields(errorLine)
	for i, part := range parts {
		if part == "port" && i+1 < len(parts) {
			if _, err := strconv.Atoi(parts[i+1]); err == nil {
				return parts[i+1]
			}
		}
	}
	return ""
}
