package pta1008

import (
	"fmt"
)

func Pta() {
	var n, m int
	if _, e := fmt.Scan(&n); e != nil {
		fmt.Println("err input")
	}
	if _, e := fmt.Scan(&m); e != nil {
		fmt.Println("err input")
	}
	m = m % n

	list := make([]int, n)
	for i := 0; i < n; i++ {
		if _, e := fmt.Scan(&list[i]); e != nil {
			fmt.Println("err input")
		}
	}

	for i := range list {
		if i < m {
			fmt.Printf("%d", list[n-m+i])
		} else {
			fmt.Printf("%d", list[i-m])
		}
		if i != n-1 {
			fmt.Printf(" ")
		}
	}
}

