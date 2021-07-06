package leet144

import (
	"fmt"
	"testing"
)

func Test_preorderTraversal(t *testing.T) {
	p := &TreeNode{
		10,
		&TreeNode{6, &TreeNode{4, nil, nil}, &TreeNode{8, nil, nil}},
		&TreeNode{14, &TreeNode{-4, &TreeNode{0, nil, nil}, nil}, &TreeNode{16, nil, nil}},
	}

	fmt.Println(preorderTraversal(p))
	//fmt.Println(preorderTraversal2(p))
	preorderTraversal3(p,20)

}
