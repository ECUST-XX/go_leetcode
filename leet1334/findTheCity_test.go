package leet1334

import (
	"fmt"
	"testing"
)

type give struct {
	num               int
	edges             [][]int
	distanceThreshold int
}

var data = []struct {
	giveValue give
	res       int
}{
	{
		give{
			4,
			[][]int{
				{0, 1, 3},
				{1, 2, 1},
				{1, 3, 4},
				{2, 3, 1},
			},
			4,
		},
		3,
	},
	{
		give{
			5,
			[][]int{
				{0, 1, 2},
				{0, 4, 8},
				{1, 2, 3},
				{1, 4, 2},
				{2, 3, 1},
				{3, 4, 1},
			},
			2,
		},
		0,
	},
}

func Test_FindTheCity(t *testing.T) {
	for k, v := range data {
		t.Run(fmt.Sprint("test", k, v), func(t *testing.T) {
			if v.res != FindTheCity(v.giveValue.num, v.giveValue.edges, v.giveValue.distanceThreshold) {
				t.Error("failed", v.giveValue)
			}
		})
	}
}
