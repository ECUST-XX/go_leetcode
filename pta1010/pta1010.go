package pta1010

import (
	. "fmt"
	"io"
)

func Pta() {
	var (
		S, L []int
		s    int
	)
	for {
		_, e := Scanf("%d", &s)
		if e == io.EOF {
			break
		}
		S = append(S, s)
	}
	if len(S) == 2 {
		L = append(L, 0, 0)
	} else {
		for i := 0; i < len(S); i += 2 {
			if S[i] == 0 && S[i+1] == 0 {
				L = append(L, 0, 0)
			}
			if S[i+1] != 0 {
				L = append(L, S[i]*S[i+1], S[i+1]-1)
			}
		}
	}
	for i, v := range L {
		if i != len(L)-1 {
			Printf("%d ", v)
		} else {
			Printf("%d", v)
		}
	}
}
