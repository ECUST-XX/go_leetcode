package leet2

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{0, nil}
	add(l1, l2, res, 0)
	return res

}

func add(l1 *ListNode, l2 *ListNode, res *ListNode, pNext int) *ListNode {

	v1 := 0
	v2 := 0
	if l1 != nil {
		v1 = l1.Val
	}
	if l2 != nil {
		v2 = l2.Val
	}

	res.Val = v1 + v2 + pNext
	pNext = 0
	if res.Val > 9 {
		pNext = 1
		res.Val -= 10
	}

	next := &ListNode{0, nil}
	res.Next = next

	if nil != l1 {
		l1 = l1.Next
	}
	if nil != l2 {
		l2 = l2.Next
	}
	if l1 == nil && l2 == nil {
		if 1 == pNext {
			next.Val = 1
			return next
		}
		res.Next = nil
		return res
	}
	return add(l1, l2, next, pNext)
}
