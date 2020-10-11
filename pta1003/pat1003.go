package pta1003

import (
	"fmt"
)

func Pta(input []string) {
	for _, v := range input {
		vArr := []byte(v)
		vLen := len(v)
		var hA, mA, tA, P, T int
		var A *int

		flag := true
		for i := 0; i < vLen; i++ {
			if !flag {
				break
			}
			if P == 0 && T == 0 {
				A = &hA
			} else if P == 1 && T == 1 {
				A = &tA
			} else if P == 1 && T == 0 {
				A = &mA
			}

			switch vArr[i] {
			case 'P':
				P++
			case 'T':
				T++
			case 'A':
				*A++
			default:
				flag = false
			}
		}

		if T == 1 && P == 1 && mA >0 && hA*mA == tA{
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}

	}
}

func ptaIn() []string {
	var n int
	if _, e := fmt.Scan(&n); e != nil {
		fmt.Println("err input")
	}
	input := make([]string, n)
	for i := 0; i < n; i++ {
		_, e := fmt.Scan(&input[i])
		if e != nil {
			fmt.Println("err input")
		}
	}

	return input
}

func PatMain() {
	Pta(ptaIn())
}
