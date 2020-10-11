package leet2

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{2, nil}
	l1.Next = &ListNode{4, nil}
	l1.Next.Next = &ListNode{3, nil}
	l2 := &ListNode{5, nil}
	l2.Next = &ListNode{6, nil}
	l2.Next.Next = &ListNode{4, nil}

	res := &ListNode{7, nil}
	res.Next = &ListNode{0, nil}
	res.Next.Next = &ListNode{8, nil}

	x := AddTwoNumbers(l1, l2)

	for {
		if res.Val != x.Val {
			t.Error("fail")
		}
		fmt.Println(x.Val, x)
		if x.Next == nil {
			break
		}
		x = x.Next
		res = res.Next
	}
}
