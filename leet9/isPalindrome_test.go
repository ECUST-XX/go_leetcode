package leet9

import (
	"fmt"
	"testing"
)

var data = map[int]bool{121: true, -121: false, 10: false}

func TestIsPalindrome(t *testing.T) {
	for give, res := range data {
		t.Run(fmt.Sprint("test", give, res), func(t *testing.T) {
			if res != IsPalindrome(give) {
				t.Error("fail")
			}
		})
	}
}
