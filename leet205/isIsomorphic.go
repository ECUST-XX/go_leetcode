package leet205

import "fmt"

func isIsomorphic(s string, t string) bool {
	ss := []byte(s)
	tt := []byte(t)

	if s == t {
		return true
	}

	rule1 := map[int]int{}
	rule2 := map[int]int{}
	for k, v := range ss {
		tmp1, ok := rule1[int(v)]
		if ok {
			if tmp1 != int(tt[k]) {
				fmt.Println("xxx", rule1, rule2)
				return false
			}
		} else {
			rule1[int(v)] = int(tt[k])
		}

		tmp2, ok := rule2[int(tt[k])]
		if ok {
			if tmp2 != int(v) {
				fmt.Println("yy", rule1, rule2)

				return false
			}
		} else {
			rule2[int(tt[k])] = int(v)
		}
	}
	fmt.Println("zz", rule1, rule2)

	return true
}
