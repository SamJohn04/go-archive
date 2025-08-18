package internal

import (
	"path/filepath"
	"testing"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name string;
		args []string;
		expected Argument;
	}{
		{
			name: "Single argument",
			args: []string{"a"},
			expected: Argument{
				Source: absolutePath("a"),
				Destination: absolutePath("a.zip"),
				DeleteOriginal: false,
			},
		}, {
			name: "Single argument with backtrack",
			args: []string{"../b"},
			expected: Argument{
				Source: absolutePath("../b"),
				Destination: absolutePath("../b.zip"),
				DeleteOriginal: false,
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			arg, err := Parse(c.args)
			if err != nil {
				t.Errorf("Something went wrong: %v", err)
			} else if arg != c.expected {
				t.Errorf("On parsing %v; Expected %v; Got %v", c.args, c.expected, arg)
			}
		})
	}
}

func absolutePath(strToAbs string) string {
	str, _ := filepath.Abs(strToAbs)
	return str
}
