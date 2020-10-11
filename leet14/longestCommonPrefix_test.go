package leet14

import (
	"fmt"
	"testing"
)

var data = []struct {
	give []string
	res  string
}{
	{[]string{"flower", "flow", "flight"}, "fl"},
	{[]string{"dog", "racecar", "car"}, ""},
	{[]string{"ss"},"ss"},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, d := range data {
		t.Run(fmt.Sprint("test", d), func(t *testing.T) {
			if d.res != LongestCommonPrefix(d.give) {
				t.Error("fail")
			}
		})
	}

}
