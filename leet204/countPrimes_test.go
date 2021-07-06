package leet204

import (
	"fmt"
	"testing"
)

var data = [][]int{{10, 4}, {100, 25}, {2, 0}}

func Test_RemoveElements(t *testing.T) {
	for _, d := range data {
		t.Run(fmt.Sprintln("test", d), func(t *testing.T) {
			if d[1] != countPrimes(d[0]) {
				t.Error("fail")
			}
		})
	}
}
