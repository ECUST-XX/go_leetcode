package leet202

import (
	"fmt"
	"testing"
)

var data = []struct {
	give int
	res  bool
}{
	{19, true},
	{1, true},
	{13, true},


}
func Test_IsHappy(t *testing.T)  {
	for _, d := range data {
		t.Run(fmt.Sprint("test", d), func(t *testing.T) {
			if d.res != IsHappy(d.give) {
				t.Error("fail")
			}
		})
	}
}
