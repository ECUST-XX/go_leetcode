package leet1356

import (
	"fmt"
	"sort"
)

type countX struct {
	val int
	cnt int
}

func SortByBits(arr []int) []int {
	count1 := make([]countX, len(arr))
	for k, v := range arr {
		str := fmt.Sprintf("%b", v)
		count1[k] = countX{
			v,
			0,
		}
		for _, s := range str {
			if s == '1' {
				count1[k].cnt ++
			}
		}
	}
	//fmt.Println(count1)
	m2 := make(map[int][]int, len(arr))
	for _, c := range count1 {
		m2[c.cnt] = append(m2[c.cnt], c.val)
	}
	//fmt.Println(m2)

	kSort := make([]int, len(m2))
	i := 0
	for k := range m2 {
		kSort[i] = k
		sort.Ints(m2[k])
		i++
	}
	sort.Ints(kSort)
	//fmt.Println(m2,kSort)

	res := make([]int, 0)
	for _, v := range kSort {
		//fmt.Println(res,"====",v,m2[v],"------",append(res,m2[v]...))
		res = append(res, m2[v]...)
	}
	//fmt.Println(arr,"-----",res)

	return res
}
