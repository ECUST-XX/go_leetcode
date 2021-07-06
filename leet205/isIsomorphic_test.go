package leet205

import "testing"

var data = []struct {
	s string
	t string
	r bool
}{
	{"aa", "bb", true},
	{"aqa", "bcb", true},
	{"paper", "title", true},
	{"a", "a", true},
	{"ab", "ab", true},
	{"foo", "bar", false},
	{"badc", "baba", false},
}

func Test_isIsomorphic(t *testing.T) {
	for _, d := range data {
		t.Run(d.s, func(t *testing.T) {
			if d.r != isIsomorphic(d.s, d.t) {
				t.Fatal("fail")
			}
		})
	}
}
