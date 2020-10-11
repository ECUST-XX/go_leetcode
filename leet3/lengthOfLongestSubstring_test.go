package leet3

import (
	"fmt"
	"testing"
)

var data = []struct {
	give string
	res  int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
}

func TestLengthOfLongestSubstring(t *testing.T) {
	for k, d := range data {
		t.Run(fmt.Sprint("test", k, d), func(t *testing.T) {
			if LengthOfLongestSubstring(d.give) != d.res {
				t.Error("failed", d)
			}
		})
	}

}
