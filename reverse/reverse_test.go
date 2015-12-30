package reverse

import (
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		text, want string
	}{
		{"hello my name is Bob", "Bob is name my hello"},
		{"hello", "hello"},
		{"A", "A"},
		{"", ""},
	}
	for _, c := range cases {
		got := ReverseSentence(c.text)
		if got != c.want {
			t.Errorf("ReverseSentence(%q) -> %q want %q", c.text, got, c.want)
		}
	}
}
