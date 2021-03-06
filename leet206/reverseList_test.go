package leet206

import (
	"fmt"
	"testing"
)

var l6 = &ListNode{6, nil}
var l5 = &ListNode{5, l6}
var l4 = &ListNode{4, l5}
var l3 = &ListNode{3, l4}
var l2 = &ListNode{2, l3}
var l1 = &ListNode{1, l2}

func Test_reverseList(t *testing.T) {
	head := reverseList(l1)
	i := 6
	for head != nil {
		fmt.Println(head.Val)
		if i != head.Val {
			t.Error("fail")
		}
		i--
		head = head.Next
	}

	head = reverseList2(l6)
	i = 1
	for head != nil {
		fmt.Println(head.Val)
		if i != head.Val {
			t.Error("fail")
		}
		i++
		head = head.Next
	}
}
