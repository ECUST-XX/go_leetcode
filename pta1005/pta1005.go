package pta1005

import (
	"fmt"
	"sort"
)

func Pta() {
	var n int
	if _, e := fmt.Scan(&n); e != nil {
		fmt.Println("err input")
	}

	input := make([]int, n)
	result := make(map[int]bool, 2*n)
	for i := 0; i < n; i++ {
		var e error
		_, e = fmt.Scan(&input[i])
		if e != nil {
			fmt.Println("err input")
		}

		num := input[i]
		for {
			if 1 == num {
				break
			}
			if input[i] != num && num < 101 {
				result[num] = true
			}

			if 0 == num%2 {
				num /= 2
			} else {
				num = (num*3 + 1) / 2
			}
		}
	}
	outRes := make([]int, 0)
	for _, v := range input {
		if _, ok := result[v]; ok {
			continue
		}
		outRes = append(outRes, v)
	}
	sort.Ints(outRes)
	for i := len(outRes) - 1; i > 0; i-- {
		fmt.Print(outRes[i], " ")
	}
	fmt.Print(outRes[0])
}
