package removeNthNodeFromEnd

import "fmt"

type ListNode struct {
	     Val int
     Next *ListNode
}
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	var s []*ListNode
	tmp  := head
	for tmp !=nil{
		s = append(s,tmp)
		tmp = tmp.Next
	}
	length := len(s)
	fmt.Printf("%v,%v\n",s,s[length-n])

	if length  -n < 0 {
		return nil
	}
	delIndex := length -n
	var start int
	if delIndex == 0 {
		if length > 1{
			start =1
			return s[start]
		}else {
			return nil
		}
	}else if delIndex ==1 {
		start = 0
		if length >2{
			s[0].Next = s[2]
		}else{
			s[0].Next = nil
		}
		return s[start]
	}else if delIndex >= length -1{
		start = 0
		s[delIndex-1].Next = nil
		return s[start]
	}else{
		start = 0
		s[delIndex-1].Next = s[delIndex+1]
		return s[start]
	}
}