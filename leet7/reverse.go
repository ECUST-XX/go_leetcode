package leet7

import (
	"fmt"
	"math"
	"strconv"
)

func Reverse(x int) int {
	res := ""
	if x < 0 {
		res = "-"
		x = -x
	}

	s := strconv.Itoa(x)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == 0 {
			if len(s)-1 == i || (len(s) != i && s[i+1] == 0) {
				continue
			}
		}
		res = fmt.Sprint(res, s[i]-48)
	}
	num, _ := strconv.Atoi(res)
	if num > math.MaxInt32 || num < math.MinInt32 {
		return 0
	}
	return num
}
