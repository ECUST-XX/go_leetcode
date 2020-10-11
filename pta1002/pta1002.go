package pta1002

import (
	"fmt"
	"strconv"
)

var zhMap = [10]string{"ling", "yi", "er", "san", "si", "wu", "liu", "qi", "ba", "jiu"}

func Pta(input string) string {
	intArr := []byte(input)
	intSum := 0
	for _, v := range intArr {
		intSum = intSum - 48 + int(v)
	}

	resInt := strconv.Itoa(intSum)
	resIntLen := len(resInt) - 1
	result := ""
	for k, v := range []byte(resInt) {
		if k == resIntLen {
			result += zhMap[v-48]
		} else {
			result += zhMap[v-48] + " "
		}
	}

	return result
}

func ptaIn() string {
	var input string
	_, e := fmt.Scan(&input)
	if e != nil {
		fmt.Println("err input")
	}
	return input
}

func patOut(result string) {
	fmt.Println(result)
}

func PatMain() {
	patOut(Pta(ptaIn()))
}
