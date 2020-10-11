package leet9

import "strconv"

func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	maxLen := len(s)
	for i := 0; i < maxLen/2; i++ {
		if s[i] != s[maxLen-i-1] {
			return false
		}
	}
	return true
}
