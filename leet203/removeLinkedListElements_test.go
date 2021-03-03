package leet203

import (
	"fmt"
	"testing"
)

var l6 = &ListNode{6, nil}
var l5 = &ListNode{5, l6}
var l4 = &ListNode{4, l5}
var l3 = &ListNode{3, l4}
var lx = &ListNode{6, l3}
var l2 = &ListNode{2, lx}
var l1 = &ListNode{1, l2}

var data = []struct {
	val       int
	nodeValue []int
}{
	//{6, []int{1, 2, 3, 4, 5}},
	{1, []int{2, 6, 3, 4, 5, 6}},
}

func Test_RemoveElements(t *testing.T) {
	for i, d := range data {
		t.Run(fmt.Sprint("test", i), func(t *testing.T) {
			head := l1
			res := removeElements(head, d.val)
			i := 0
			for res.Next != nil {
				fmt.Println(res.Val)
				if res.Val != d.nodeValue[i] {
					t.Error("fail")
				}
				res = res.Next
				i++
			}
			fmt.Println(res.Val)
			if res!=nil && res.Val != d.nodeValue[i] {
				t.Error("fail")
			}
		})
	}

	// 边界条件好恶心.....
	a := &ListNode{1, nil}
	fmt.Println(removeElements(a, 1))

	b := &ListNode{1, &ListNode{1, nil}}
	fmt.Println(removeElements(b, 1))

	c := &ListNode{1, &ListNode{1, &ListNode{1, nil}}}
	fmt.Println(removeElements(c, 1))

	d := &ListNode{1, &ListNode{2, nil}}
	fmt.Println(removeElements(d, 1))

	e := &ListNode{1, &ListNode{2,  &ListNode{1, nil}}}
	fmt.Println(removeElements(e, 2))
}
