package pta1004

import "fmt"


func Pta() {
	var n int
	if _, e := fmt.Scan(&n); e != nil {
		fmt.Println("err input")
	}
	inputName := make([]string, n)
	inputNum := make([]string, n)
	inputScore := make([]int, n)

	min:=0
	max:=0
	for i := 0; i < n; i++ {
		var e error
		_, e = fmt.Scan(&inputName[i])
		if e != nil {
			fmt.Println("err input")
		}
		_, e = fmt.Scan(&inputNum[i])
		if e != nil {
			fmt.Println("err input")
		}
		_, e = fmt.Scan(&inputScore[i])
		if e != nil {
			fmt.Println("err input")
		}
		if inputScore[i] > inputScore[max] {
			max = i
		}
		if inputScore[i] < inputScore[min] {
			min = i
		}
	}

	fmt.Println(inputName[max],inputNum[max])
	fmt.Println(inputName[min],inputNum[min])
}

