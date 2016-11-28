package mergeTwoSortedLists


//Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func MergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var curListNode *ListNode = nil
	if l1 == nil && l2 == nil{
		return curListNode
	}else if l1 == nil {
		return l2
	}else if l2 == nil{
		return l1
	}
	if l1.Val < l2.Val {
		curListNode = l1
		l1 = l1.Next
	}else{
		curListNode = l2
		l2 = l2.Next
	}
	start := curListNode
	for l1 !=nil && l2 !=nil{
		if l1.Val < l2.Val {
			curListNode.Next = l1
			l1 = l1.Next
		}else{
			curListNode.Next = l2
			l2 = l2.Next
		}
		curListNode = curListNode.Next
	}

	if l1==nil && l2 !=nil {
		curListNode.Next = l2
	}

	if (l2==nil && l1 !=nil){
		curListNode.Next = l1
	}
	return start
}

