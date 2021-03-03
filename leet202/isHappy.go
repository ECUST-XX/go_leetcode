package leet202

import (
	"strconv"
)

var squareArr = [10]int{0, 1, 4, 9, 16, 25, 36, 49, 64, 81}

func IsHappy(n int) bool {
	var resMap = make(map[int]int)

	for {
		res := getSquare(n)
		if 1 == res {
			return true
		}
		if _, ok := resMap[res]; ok {
			return false
		}
		resMap[res] = n

		n = res
	}
}

func getSquare(n int) int {
	numStr := strconv.Itoa(n)

	res := 0
	for _, num := range numStr {
		numInt := int(num) - 48
		res += squareArr[numInt]
	}
	return res
}
