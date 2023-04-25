package IntLinkedList

import (
	"testing"
)

func Test_Equals(t *testing.T) {

	succ_case := []*ListNode{

		SliceToLinkedList([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
		SliceToLinkedList([]int{1e1, 1e2, 1e3, 1e4, 1e4, 1e5, 1e6, 1e7, 1e8, 1e9}),
		SliceToLinkedList([]int{1e10, 1e11, 1e12, 1e13, 1e14, 1e15, 1e16, 1e17, 1e18}),
		SliceToLinkedList([]int{-1e10, -1e11, -1e12, -1e13, -1e14, -1e15, -1e16, -1e17, -1e18}),
	}
	for _, v := range succ_case {
		if !Equals(v, v) {
			t.Errorf("%v should be same", v)
		}
	}

	fail_case := []struct {
		in1, in2 *ListNode
	}{
		{SliceToLinkedList([]int{1, 2, 3, 4, 2, 1}),
			SliceToLinkedList([]int{2, 2, 3, 4, 2, 1})},

		{SliceToLinkedList([]int{100, 2, -2, 3, 4, 2, 1}),
			SliceToLinkedList([]int{2, 2, 3, 4, 2, 1})},

		{SliceToLinkedList([]int{2, 2, 3, 4, 2, 1, 100}),
			SliceToLinkedList([]int{2, 2, 3, 4, 2, 1})},
	}

	for _, v := range fail_case {
		if Equals(v.in1, v.in2) {
			t.Errorf("%v,%v should be different", v.in1, v.in2)
		}
	}
}
func Test_sliceToLinkedList(t *testing.T) {
	inout := []struct {
		in   []int
		want *ListNode
	}{
		{[]int{1, 2, 2, 1},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 2,
						Next: &ListNode{Val: 1,
							Next: nil}}}}},
		{[]int{1, 2},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: nil}}},

		{[]int{1},
			&ListNode{Val: 1,
				Next: nil}},

		{[]int{1, 2, 3, 2, 1},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 2,
							Next: &ListNode{Val: 1,
								Next: nil}}}}}},
		{[]int{1, 2, 3, 4, 2, 1},
			&ListNode{Val: 1,
				Next: &ListNode{Val: 2,
					Next: &ListNode{Val: 3,
						Next: &ListNode{Val: 4,
							Next: &ListNode{Val: 2,
								Next: &ListNode{Val: 1,
									Next: nil}}}}}}},
	}
	for _, v := range inout {
		got := SliceToLinkedList(v.in)
		for nodeg, nodew := got, v.want; ; nodeg, nodew = nodeg.Next, nodew.Next {
			if nodeg == nil || nodew == nil {
				if !(nodeg == nil && nodew == nil) {
					t.Errorf("length mismatch")
				}
				break
			}
			if nodeg.Val != nodew.Val {
				t.Errorf("got:%v,want:%v", nodeg.Val, nodew.Val)
			}
		}
	}
}

func Test_Reverse(t *testing.T) {
	inout := []struct {
		in   *ListNode
		want *ListNode
	}{
		{SliceToLinkedList([]int{1, 2, 2, 3}),
			SliceToLinkedList([]int{3, 2, 2, 1})},

		{SliceToLinkedList([]int{1, 2, 3, 4}),
			SliceToLinkedList([]int{4, 3, 2, 1})},

		{SliceToLinkedList([]int{}),
			SliceToLinkedList([]int{})},
	}
	for _, v := range inout {

		got := Reverse(v.in)
		for nodeg, nodew := got, v.want; ; nodeg, nodew = nodeg.Next, nodew.Next {
			if nodeg == nil || nodew == nil {
				if !(nodeg == nil && nodew == nil) {
					t.Errorf("length mismatch")
				}
				break
			}
			if nodeg.Val != nodew.Val {
				t.Errorf("got:%v,want:%v", nodeg.Val, nodew.Val)
			}
		}
	}

}
