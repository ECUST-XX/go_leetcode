package leet100

import (
	"fmt"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsSameTree(p *TreeNode, q *TreeNode) bool {

	// 偷鸡。。。
	//reflect.DeepEqual(p,q)
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}
	return reflect.DeepEqual(treeToArray(p), treeToArray(q))
}

func treeToArray(p *TreeNode) []int {
	res := make([]int, 0)
	res = getNode(p, res)
	fmt.Println(res)
	return res
}

func getNode(p *TreeNode, res []int) []int {

	res = append(res, p.Val)
	if p.Left != nil {
		res = getNode(p.Left, res)
	} else {
		res = append(res, -999)
	}
	if p.Right != nil {
		res = getNode(p.Right, res)
	} else {
		res = append(res, -999)
	}
	return res
}
