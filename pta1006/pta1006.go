package pta1006

import (
	"fmt"
)

func Pta() {
	var n int
	if _, e := fmt.Scan(&n); e != nil {
		fmt.Println("err input")
	}
	bai := n / 100
	shi := (n - bai*100) / 10
	ge := n%10
	for i := 0; i < bai; i++ {
		fmt.Print("B")
	}
	for i := 0; i < shi; i++ {
		fmt.Print("S")
	}
	for i := 1; i <= ge; i++ {
		fmt.Print(i)
	}
}
