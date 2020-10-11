package leet14

import "fmt"

func LongestCommonPrefix(strs []string) string {
	res := ""
	strsLen := len(strs)
	if strsLen == 1 {
		return strs[0]
	} else if strsLen == 0 {
		return ""
	}
	s := strs[0]
	flag := false
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		for j := 1; j < strsLen; j++ {
			if i < len(strs[j]) && strs[j][i] == s[i] {
				continue
			}
			flag = true
			break
		}
		if flag {
			break
		}
		res = fmt.Sprint(res, string(s[i]))
	}
	return res
}
