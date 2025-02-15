package main

import (
	"flag"
	"fmt"
	"github.com/DDaaaaann/kpop-cli/internal"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"os"
)

var forceFlag = flag.Bool("f", false, "Force kill without confirmation")
var quietFlag = flag.Bool("q", false, "Quiet mode, suppress output")
var realExecutor = executor.RealCommandExecutor{}

func usage() {
	fmt.Println("Usage: kill-port [-f] [-q] <port>")
	fmt.Println("Options:")
	fmt.Println("  -f    Force kill without confirmation")
	fmt.Println("  -q    Quiet mode, suppress output")
	fmt.Println("  -h    Show this help message")
}

func main() {
	port := extractPortFromArgs()
	internal.KPOP(port, *forceFlag, *quietFlag, os.Stdin, os.Stdout, &realExecutor)
}

func extractPortFromArgs() string {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Println("Error: No port specified.")
		usage()
		os.Exit(1)
	}
	return flag.Arg(0)
}
