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
		}, {
			name: "Single complicated argument",
			args: []string{"../a/b"},
			expected: Argument{
				Source: absolutePath("../a/b"),
				Destination: absolutePath("../a/b.zip"),
				DeleteOriginal: false,
			},
		}, {
			name: "Single + Delete Original",
			args: []string{"-d", "../a/b"},
			expected: Argument{
				Source: absolutePath("../a/b"),
				Destination: absolutePath("../a/b.zip"),
				DeleteOriginal: true,
			},
		}, {
			name: "Source and Destination",
			args: []string{"-o", "../c.zip", "../a/b"},
			expected: Argument{
				Source: absolutePath("../a/b"),
				Destination: absolutePath("../c.zip"),
				DeleteOriginal: false,
			},
		}, {
			name: "Source and Destination + Delete Original",
			args: []string{"-d", "-o", "../c.zip", "../a/b"},
			expected: Argument{
				Source: absolutePath("../a/b"),
				Destination: absolutePath("../c.zip"),
				DeleteOriginal: true,
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
