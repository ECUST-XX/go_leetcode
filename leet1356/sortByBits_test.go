package leet1356

import (
	"fmt"
	"reflect"
	"testing"
)

var data = []struct {
	give []int
	res  []int
}{
	{
		[]int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		[]int{0, 1, 2, 4, 8, 3, 5, 6, 7},
	},
	{
		[]int{1024, 512, 256, 128, 64, 32, 16, 8, 4, 2, 1},
		[]int{1, 2, 4, 8, 16, 32, 64, 128, 256, 512, 1024},
	},
	{
		[]int{10000, 10000},
		[]int{10000, 10000},
	},
	{
		[]int{2, 3, 5, 7, 11, 13, 17, 19},
		[]int{2, 3, 5, 17, 7, 11, 13, 19},
	}, {
		[]int{10, 100, 1000, 10000},
		[]int{10, 100, 10000, 1000},
	},
}

func Test_SortByBits(t *testing.T) {
	for k, v := range data {
		t.Run(fmt.Sprint("test", k), func(t *testing.T) {
			if !reflect.DeepEqual(v.res, SortByBits(v.give)) {
				t.Error("failed", v)
			}
		})

	}
}
