# go leetcode

## 一，相关知识

### 1.1数据结构知识

### 1.2算法知识

### 1.3时间复杂度

## 二，解题

### 0001～0099

#### [0002. Add Two Numbers](https://leetcode.com/problems/add-two-numbers/)

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order**, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/10/02/addtwonumber1.jpg)

```
Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
```

**Example 2:**

```
Input: l1 = [0], l2 = [0]
Output: [0]
```

**Example 3:**

```
Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]
```

思路：

这个问题的解决方案实际上就是模拟了我们手动相加两个数的过程。首先，我们创建一个空链表来存储每一位的相加结果。在每一步中，我们将对应的节点值相加，并加上可能存在的进位（carry）。然后，我们将相加的结果对10取余数，这个余数就是新的节点的值。最后，我们通过将相加的结果除以10，得到的商就是下一步的进位。如果最后还有进位，我们会在链表的末尾添加一个新的节点

答案：

```go
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := &ListNode{}
    current , n1 ,n2 , carray := head, 0,0,0
    for l1 != nil || l2 != nil || carray != nil{
        if l1 == nil{
            n1 = 0
        }else{
            n1 = l1.Val
            l1 = l1.Next
        }
        if l2 == nil{
            n2 = 0
        }else{
            n2 = l2.Val
            l2 = l2.Next
        }
        current.Next = &ListNode{Val:(n1 + n2 +carray)%10}
        carray := (n1 + n2 +carray) /10
    }
    return current.Next
}
```

#### [0021. Merge Two Sorted Lists](https://leetcode.com/problems/merge-two-sorted-lists/)

Merge the two lists into one **sorted** list. The list should be made by splicing together the nodes of the first two lists.

Return *the head of the merged linked list*.



解法一：

创建空的结果链表，然后每次比较l1 和l2 中的节点val大小，将结果指针节点指向结果较小的节点，当某个list为空的时候，将结果指针指向其余的节点

```go
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}
	return head.Next
}
```

解法二： 使用递归

```go
func mergeTwoList(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoList(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoList(l1, l2.Next)
		return l2
	}
}
```

