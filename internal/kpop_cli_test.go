package internal

import (
	"bytes"
	"errors"
	"github.com/DDaaaaann/kpop-cli/internal/executor"
	"github.com/DDaaaaann/kpop-cli/internal/utils"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestKpopCLI(t *testing.T) {
	var stdin, stdout bytes.Buffer
	var tests = []struct {
		name string
		port string
		pid  int
		err  error
		kill bool
		want string
	}{
		{
			name: "Successfully kill process on port",
			port: "8080",
			pid:  9999,
			kill: true,
			want: `Kill process on port 8080 (PID 9999)? (y/n)
Killed process 9999 on port 8080.`,
		},
		{
			name: "Cancel kill process on port",
			port: "123",
			pid:  98765,
			kill: false,
			want: `Kill process on port 123 (PID 98765)? (y/n)
Cancelled.`,
		},
		{
			name: "Wrong port input",
			port: "abc",
			want: `Port 'abc' is not well formatted`,
		},
		{
			name: "No pid found",
			port: "123",
			err:  errors.New("could not extract pid"),
			want: `No process found using port 123`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockExecutor := executor.MockCmdExecutor([]byte(strconv.Itoa(tt.pid)), tt.err, utils.FormatPIDOnly)

			if tt.kill {
				confirm(&stdin)
			} else {
				decline(&stdin)
			}

			KPOP(tt.port, false, false, &stdin, &stdout, mockExecutor)
			assert.Contains(t, stdout.String(), tt.want)
		})
	}
}

func decline(buf *bytes.Buffer) {
	buf.WriteString("n\n")
}

func confirm(buf *bytes.Buffer) {
	buf.WriteString("y\n")
}
