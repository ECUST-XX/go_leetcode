package leet704

import (
	"fmt"
	"testing"
)

func Test_search(t *testing.T) {
	nums := []int{-1, 12, 13, 40, 55, 610}
	target := 13

	fmt.Println(search2(nums, target))
}
