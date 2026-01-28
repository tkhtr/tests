package fullname

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullName(t *testing.T) {
	tests := []struct {
		name  string
		value User
		want  string
	}{
		{
			name: "test #1",
			value: User{
				FirstName: "Tokhtar",
				LastName:  "Aubakirov",
			},
			want: "Tokhtar Aubakirov",
		},
		{
			name: "test #2",
			value: User{
				FirstName: "Aigerim",
				LastName:  "Aubakirova",
			},
			want: "Aigerim Aubakirova",
		},
		{
			name: "test #2",
			value: User{
				FirstName: "Jone",
				LastName:  "Jones",
			},
			want: "Jone Jones",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, test.value.FullName())
		})
	}
}
