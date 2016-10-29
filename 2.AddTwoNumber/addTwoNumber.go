package addTwoNumber

/*
You are given two linked lists representing two non-negative numbers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	start := &ListNode{Val:0, Next:nil}
	last := start
	inc := 0
	sum := 0
	for {
		if l1 == nil && l2 == nil {
			if inc == 1{
				last.Next = &ListNode{Val:1}
			}
			return  start
		}else  if (l1 == nil){
			l1 = nil
			sum = l2.Val +inc
			l2 = l2.Next
		}else if (l2 == nil){
			l2 = nil
			sum = l1.Val + inc
			l1= l1.Next
		}else{
			sum = l1.Val + l2.Val + inc
			l1 = l1.Next
			l2 = l2.Next
		}
		last.Val = sum % 10
		inc = 0
		if sum / 10 == 1 {
			inc =1
		}
		if l1 != nil || l2 != nil{
			next := &ListNode{}
			last.Next = next
			last = next
		}
	}
}