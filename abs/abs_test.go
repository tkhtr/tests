package abs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "test #1: simple check",
			value: 3.1,
			want:  3.1,
		},
		{
			name:  "test #2: negative value check",
			value: -3.14,
			want:  3.14,
		},
		{
			name:  "test #3: negative zero check",
			value: -0,
			want:  0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Abs(test.value))
		})
	}
}
