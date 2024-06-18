package leetcode

import (
	"reflect"
	"testing"
)

func TestAddTwoNumber(t *testing.T) {
	tests := []struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			l1:       &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2:       &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			expected: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		{
			l1:       &ListNode{0, nil},
			l2:       &ListNode{0, nil},
			expected: &ListNode{0, nil},
		},
		{
			l1:       &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			l2:       &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			expected: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
	}

	for _, test := range tests {
		result := AddTwoNumber(test.l1, test.l2)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("AddTwoNumber(%v, %v) = %v; expected %v", test.l1, test.l2, result, test.expected)
		}
	}
}
