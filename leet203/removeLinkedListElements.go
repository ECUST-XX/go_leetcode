package leet203

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}

	firstHead := head
	last := head
	isFirst := true

	for head.Next != nil {
		if head.Val == val {
			if isFirst {
				firstHead = head.Next
				last = head.Next
				head = head.Next
				continue
			} else {
				isFirst = false
				last.Next = head.Next
				head.Next = nil
				head = last.Next
			}
		} else {
			isFirst = false
			last = head
			head = head.Next
		}
	}

	if head.Val == val {
		last.Next = nil
	}

	if firstHead.Val == val {
		firstHead = nil
	}

	return firstHead
}
