package leetcode

import "testing"

func TestIsSameTree(t *testing.T) {
	// 创建两棵相同的二叉树
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   3,
			Left:  nil,
			Right: nil,
		},
	}

	// 调用 isSameTree 函数，期望结果为 true
	result := isSameTree(p, q)

	if !result {
		t.Errorf("isSameTree(p, q) = %v; want true", result)
	}

	// 创建两棵不同的二叉树
	p = &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

	q = &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	// 调用 isSameTree 函数，期望结果为 false
	result = isSameTree(p, q)

	if result {
		t.Errorf("isSameTree(p, q) = %v; want false", result)
	}
}
