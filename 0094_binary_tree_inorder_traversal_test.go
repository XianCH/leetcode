package leetcode

import (
	"reflect"
	"testing"
)

func InorderTraversalTest(t *testing.T) {
	root := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   3,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}
	result := InorderTraversal(root)
	expect := []int{1, 3, 2}

	if !reflect.DeepEqual(result, expect) {
		t.Errorf("InorderTravel() = %v,want = %v", result, expect)
	}
}
