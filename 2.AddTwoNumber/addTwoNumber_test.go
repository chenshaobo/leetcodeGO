package addTwoNumber

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := SliceToListNode([]int{1, 2, 3, 4, 5, 6})
	l2 := SliceToListNode([]int{4, 5, 6, 7, 8, 9, 9})
	r := AddTwoNumbers(l1, l2)
	t.Logf("%v\n",ListToSlice(r))
	if ! Equal(ListToSlice(r),[]int{5,7,9,1,4,6,0,1}) {
		t.Fatal("Not match")
	}
}

func SliceToListNode(s []int) *ListNode {
	var start,next *ListNode
	for i, v := range s {
		cur := &ListNode{Val:v}
		if i == 0 {
			start = cur
			next = cur
		} else {
			next.Next = cur
			next = cur
		}

	}
	return start
}

func ListToSlice(l *ListNode) []int {
	s := make([]int, 0)
	for {
		s = append(s, l.Val)
		if l.Next == nil {
			return s
		}
		l = l.Next
	}
}

func Equal(a,b []int) bool{
	for i,v :=range a {
		if b[i] != v{
			return false
		}
	}
	return true
}