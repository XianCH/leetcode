package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoList(t *testing.T) {
	tests := []struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			l1:       &ListNode{1, &ListNode{2, &ListNode{4, nil}}},
			l2:       &ListNode{1, &ListNode{3, &ListNode{4, nil}}},
			expected: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, nil}}}}}},
		}, {
			l1:       nil,
			l2:       &ListNode{0, nil},
			expected: &ListNode{0, nil},
		},
	}

	for _, test := range tests {
		result := mergeTwoList(test.l1, test.l2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("mergeTwoList(%v, %v) = %v; expected %v", test.l1, test.l2, result, test.expected)
		}
	}
}
