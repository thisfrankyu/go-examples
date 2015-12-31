package next_greatest_integer

import (
	"testing"
)


func TestContains(t *testing.T) {
	cases := []struct {
		num int
		want int
	}{
		{12345, 12354},
		{12, 21},
		{425116, 425161},
		{5693013, 5693031},
		//specifically leaving off testing the 0 case
	}
	for _, c := range cases {
		got := GetNextGreatestInt(c.num)
		if got != c.want {
			t.Errorf("GetNextGreatestInt(%v) -> got %v want %v", c.num, got, c.want)
		}
	}
}

