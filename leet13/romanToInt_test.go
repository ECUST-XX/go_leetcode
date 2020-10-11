package leet13

import (
	"fmt"
	"testing"
)

var data = []struct {
	give string
	res  int
}{
	{"III", 3},
	{"IV", 4},
	{"IX", 9},
	{"LVIII", 58},
	{"MCMXCIV", 1994},
}

func TestRomanToInt(t *testing.T) {
	for _, d := range data {
		t.Run(fmt.Sprint("test", d), func(t *testing.T) {
			if RomanToInt(d.give) != d.res {
				t.Error("fail")
			}
		})

	}
}
