package leet3

import "fmt"

// todo 滑动窗口
func LengthOfLongestSubstring(s string) int {

	p1 := 0
	p2 := 0
	strMap := make(map[string]int)

	maxStrMap:=make(map[byte]int)

	maxStr:=""

	for k,v:=range s{
		maxStrMap[byte(v)]++
		if maxStrMap[byte(v)]>1{
			// 分割
			ts:=s[p1:p2+1]
			for tsK,tsV:=range ts{
				if tsV==v{
					if tsK-p1+1>p2-tsK+1{

					}
				}

			}


		}



	}
	fmt.Println(strMap, p2, p1)

	return p1 - p2
}
