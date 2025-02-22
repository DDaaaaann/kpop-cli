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
			name:     "Pid-only format success",
			output:   []byte("9999\n"),
			format:   utils.FormatPIDOnly,
			expected: 9999,
		},
		{
			name:        "Pid-only format failure",
			output:      []byte("abc"),
			format:      utils.FormatPIDOnly,
			expectedErr: "non pid format: 'abc'",
		},
		{
			name:        "Pid-only format failure",
			output:      []byte("\n"),
			format:      utils.FormatPIDOnly,
			expectedErr: "no output found",
		},
		{
			name: "Netstat format success",
			output: []byte(`  TCP    0.0.0.0:12345          0.0.0.0:0              LISTENING       9999\n
				  TCP    [::]:12345             [::]:0                 LISTENING       9999`),
			format:   utils.FormatNetstat,
			expected: 9999,
		},
		{
			name:        "Unsupported format failure",
			output:      []byte("9999"),
			format:      -1,
			expectedErr: "unsupported format",
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
