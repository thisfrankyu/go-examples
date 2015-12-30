package mutest

import (
	"testing"
)


func TestContains(t *testing.T) {
	cases := []struct {
		pattern, text string
		want          bool
	}{
		{"hello", "hello", true},
		{"hell", "hello", true},
		{"elh", "hello", false},
		{"h.ll", "hello", true},
		{".", "hello", true},
		{"......", "hello", false},
		{"", "anything", true},
	}
	for _, c := range cases {
		got := Matches(c.pattern, c.text)
		if got != c.want {
			t.Errorf("Matches(%q, %q) -> %q want %q", c.pattern, c.text, got, c.want)
		}
	}
}

