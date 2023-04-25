package IntLinkedList

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToLinkedList(nums []int) *ListNode {

	root := new(ListNode)
	parent := root
	for i, v := range nums {
		parent.Val = v

		// except last node
		if i+1 < len(nums) {
			parent.Next = new(ListNode)
			parent = parent.Next
		}
	}
	return root
}

// アクセス不能になるわけではないが破壊的変更
func Reverse(head *ListNode) *ListNode {
	cur, prev := head, (*ListNode)(nil)

	for cur != nil {
		next := cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}
	return prev

}

func (head *ListNode) String() string {
	s := ""
	for node := head; node != nil; node = node.Next {
		s += fmt.Sprintf("&{%d %p},", node.Val, node.Next)
	}
	return s[:len(s)-1]
}

func Equals(l1, l2 *ListNode) bool {

	for n1, n2 := l1, l2; ; n1, n2 = n1.Next, n2.Next {
		if n1 == nil || n2 == nil {
			if n1 == nil && n2 == nil {
				return true
			} else {
				return false
			}
		}
		if n1.Val != n2.Val {
			return false
		}
	}
}
