package leet100

import (
	"reflect"
	"testing"
)

func Test_IsSameTree(t *testing.T) {
	p := TreeNode{
		1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	q := TreeNode{
		1,
		&TreeNode{2, nil, nil},
		&TreeNode{3, nil, nil},
	}
	if IsSameTree(&p, &q) != reflect.DeepEqual(p, q) {
		t.Log("failed", p, q)
	}

}
