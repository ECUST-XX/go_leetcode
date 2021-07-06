package leet144

import (
	"log"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	if root != nil {
		res = append(res, root.Val)
		res = append(res, preorderTraversal(root.Left)...)
		res = append(res, preorderTraversal(root.Right)...)
	}

	return res
}

func preorderTraversal2(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	res = append(res, root.Val)
	res = append(res, pp(root)...)

	return res
}

func pp(root *TreeNode) []int {
	res := make([]int, 0, 2)
	if root.Left != nil {
		res = append(res, root.Left.Val)
	}
	if root.Right != nil {
		res = append(res, root.Right.Val)
	}

	if root.Left != nil {
		res = append(res, pp(root.Left)...)
	}
	if root.Right != nil {
		res = append(res, pp(root.Right)...)
	}
	return res
}

func preorderTraversal3(root *TreeNode, sum int) {
	res := make([]int, 0)

	sumPath(root, sum, res)
}

func sumPath(root *TreeNode, sum int, v []int) (over bool) {
	over = true
	if root == nil {
		over = true
		return
	}

	v = append(v, root.Val)
	if sum-root.Val == 0 {
		log.Print(v)
	}

	if root.Left == nil && root.Right == nil {
		over = true
	}

	lo := sumPath(root.Left, sum-root.Val, v)
	ro := sumPath(root.Right, sum-root.Val, v)
	if lo {
		root.Left = nil
	}
	if ro {
		root.Right = nil
	}

	return
}
