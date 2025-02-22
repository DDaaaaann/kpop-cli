package utils

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

// ProcessOutputFormat represents the expected format of the command output.
type ProcessOutputFormat int

const (
	FormatPIDOnly ProcessOutputFormat = iota // Direct PID output (Linux/macOS)
	FormatNetstat                            // Netstat-style output (Windows)
)

// String returns the string representation of the ProcessOutputFormat.
func (f ProcessOutputFormat) String() string {
	return [...]string{"FormatPIDOnly", "FormatNetstat"}[f]
}

// ParseFirstPID extracts the first PID based on the given output format.
func ParseFirstPID(output []byte, format ProcessOutputFormat) (int, error) {
	lines := bytes.Split(output, []byte("\n"))

	for _, line := range lines {
		lineStr := strings.TrimSpace(string(line))
		if lineStr == "" {
			return 0, fmt.Errorf("no output found")
		}

		switch format {
		case FormatPIDOnly:
			// Directly parse a single PID (Linux/macOS `lsof -t`)
			pid, err := strconv.Atoi(lineStr)
			if err != nil {
				return 0, fmt.Errorf("non pid format: '%s'", lineStr)
			}
			return pid, nil
		case FormatNetstat:
			// Windows `netstat -ano` style parsing
			fields := strings.Fields(lineStr)
			if len(fields) >= 5 && strings.Contains(fields[1], ":") { // Ensure it's a valid netstat entry
				pid, err := strconv.Atoi(fields[len(fields)-1]) // PID is the last field
				if err == nil {
					return pid, nil
				}
			}
		default:
			return 0, fmt.Errorf("unsupported format")
		}
	}

	return 0, fmt.Errorf("no process found in output")
}
