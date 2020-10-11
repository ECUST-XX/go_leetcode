package pta1007

import (
	"fmt"
	"math"
)

func Pta() {
	var n int
	if _, e := fmt.Scan(&n); e != nil {
		fmt.Println("err input")
	}
	if 4 > n {
		fmt.Println(0)
		return
	}
	resArr := make([]int, n/2+3)
	resArr[0] = 2
	res := 0
	a := 1
	for i := 3; i <= n; i += 2 {
		flag := false
		for m := 3; m <= int(math.Sqrt(float64(i))); m += 2 {
			if 0 == i%m {
				flag = true
				break
			}
		}
		if !flag {
			if a > 0 && resArr[a-1]+2 == i {
				//fmt.Println(resArr[a-1],i)

				res++
			}
			resArr[a] = i
			a++
		}
	}
	//fmt.Println(resArr)

	fmt.Println(res)
}
