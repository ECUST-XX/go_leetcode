package leet1

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := []int{0, 1}

	if !reflect.DeepEqual(res, TwoSum(nums, target)) {
		t.Error("failed")
	}
}
