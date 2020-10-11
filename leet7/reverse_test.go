package leet7

import (
	"fmt"
	"testing"
)

var data = map[int]int{123: 321, -123: -321, 120: 21}

func TestReverse(t *testing.T) {
	for give, res := range data {
		t.Run(fmt.Sprint("test", give, res), func(t *testing.T) {
			if res != Reverse(give) {
				t.Error("fail")
			}
		})
	}
}
