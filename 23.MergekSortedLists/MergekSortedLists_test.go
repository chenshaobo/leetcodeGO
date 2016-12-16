package MergekSortedLists

import (
	"testing"
	"fmt"
	"math/rand"
)

func TestMergeKLists(t *testing.T) {

	var s []*ListNode
	for i := 0; i < 2000; i++ {
		var t *ListNode
		//if i ==0{
		//	t = nil
		//}else {
			j := int(rand.Int63n(1000))
			t = &ListNode{
				Val:j,
			}

			//t1 := &ListNode{
			//	Val:j*2,
			//
			//}
			//t.Next = t1
		//}

		s = append(s, t)

		//fmt.Print(s)
	}
	r := MergeKLists(s)
	for r != nil {
		fmt.Printf("%v\n", r.Val)
		r = r.Next
	}
}
