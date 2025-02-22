package unit

import (
	"github.com/DDaaaaann/kpop-cli/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseProcessOutput(t *testing.T) {
	tests := []struct {
		name        string
		output      []byte
		format      utils.ProcessOutputFormat
		expected    int
		expectedErr string
	}{
		{
			name:     "Pid Only success",
			output:   []byte("9999\n"),
			format:   utils.FormatPIDOnly,
			expected: 9999,
		},
		{
			name:        "Pid Only failure",
			output:      []byte("abc"),
			format:      utils.FormatPIDOnly,
			expectedErr: "non pid format: 'abc'",
		},
		{
			name:        "Pid Only failure",
			output:      []byte("\n"),
			format:      utils.FormatPIDOnly,
			expectedErr: "no output found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out, err := utils.ParseFirstPID(tt.output, tt.format)

			if len(tt.expectedErr) > 0 {
				assert.Error(t, err)
				assert.Equal(t, tt.expectedErr, err.Error())
			}

			assert.Equal(t, tt.expected, out, "Expected '%d' as output.", tt.expected)
		})
	}
}
