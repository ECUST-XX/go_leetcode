package leet206

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		next := head.Next
		head.Next = prev
		prev, head = head, next
	}
	return prev
}

func reverseList2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	x := reverseList2(head.Next)
	head.Next.Next = head
	head.Next = nil

	return x

}

