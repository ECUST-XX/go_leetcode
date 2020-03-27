package leet914

import (
	"fmt"
	"testing"
)

var data = []struct {
	give []int
	res  bool
}{
	{
		[]int{1, 2, 3, 4, 4, 3, 2, 1},
		true,
	},

	{
		[]int{1, 1, 1, 2, 2, 2, 3, 3},
		false,
	},

	{
		[]int{1},
		false,
	},

	{
		[]int{1, 1},
		true,
	},

	{
		[]int{1, 1, 2, 2, 2, 2},
		true,
	},
	{
		[]int{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 2, 2, 3, 3, 4, 4},
		true,
	},
}

func Test_HasGroupsSizeX(t *testing.T) {
	for k, v := range data {
		t.Run(fmt.Sprint("test", k, v), func(t *testing.T) {
			if v.res != HasGroupsSizeX(v.give) {
				t.Error("failed", v)
			}
		})
	}
}
