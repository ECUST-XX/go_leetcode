package pta1001

import "fmt"

func Pta(input int) int {
	result := 0
	for {
		if 1 == input {
			break
		}
		if 0 == input%2 {
			input /= 2
			result++
		} else {
			input = input*3 + 1
		}
	}
	return result
}

func ptaIn() int {
	var input int
	_, e := fmt.Scan(&input)
	if e != nil {
		fmt.Println("err input")
	}
	return input
}

func patOut(result int) {
	fmt.Println(result)
}

func PatMain() {
	patOut(Pta(ptaIn()))
}
