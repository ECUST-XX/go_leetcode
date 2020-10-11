package pta1009

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Pta() {
	input := ""

	inputReader := bufio.NewReader(os.Stdin)
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Println("err input", err)
	}
	input= strings.Replace(input,"\n","",-1)
	strArr := make([]string, 0)
	strArr = strings.Split(input, " ")

	for i := len(strArr) - 1; i > -1; i-- {
		if 0 == i {
			fmt.Print(strArr[i])
		} else {
			fmt.Print(strArr[i], " ")
		}
	}
}
