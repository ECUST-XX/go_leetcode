package leet914

import (
	"sort"
)

func HasGroupsSizeX(deck []int) bool {
	if 1 == len(deck) {
		return false
	}
	sort.Ints(deck)

	tempX := 1
	res := make([]int, 0)
	for k, v := range deck {
		if 0 == k {
			continue
		}
		if v == deck[k-1] {
			tempX++

		} else {
			res = append(res, tempX)
			tempX = 1
		}
		if k == len(deck)-1 {
			res = append(res, tempX)
		}
	}

	if 1 == len(res) && res[0] > 1 {
		return true
	}
	sort.Ints(res)

	if 1 == res[0] {
		return false
	}
	flag := true
	for i := 0; i < len(res); i++ {
		for x := i+1; x < len(res); x++ {
			if res[x]%res[i] != 0 {
				flag = false
			}
		}
	}
	if flag {
		return flag
	}

	for i := 2; i <= res[0]; i++ {
		for x := 0; x < len(res); x++ {
			if res[x]%i != 0 {
				flag = false
				break
			}
			flag = true
		}
		if flag {
			return flag
		}
	}

	return false
}
