package MergekSortedLists

import (
	"container/heap"
	//"fmt"
	//"sort"
)


/**
 * Definition for singly-linked list.
 *
 */
type ListNode struct {
   Val int
   Next *ListNode
}

type MergeListNode struct {
	From  int
	Val int
}

type MergeListNodeHeap []MergeListNode

func (l MergeListNodeHeap ) Len() int{
	return len(l)
}

func (l MergeListNodeHeap ) Less(i,j int) bool{
	return l[i].Val < l[j].Val
}

func (l MergeListNodeHeap ) Swap(i,j int){
	length := len(l)
	if i <length && j <length {
		l[i], l[j] = l[j], l[i]
	}
}


func (l *MergeListNodeHeap ) Push(x interface{}){
	*l = append(*l,x.(MergeListNode))
}

func (l *MergeListNodeHeap ) Pop() interface{}{
	old := *l
	n := len(old)
	x := old[n-1]
	*l = old[0 : n-1]
	return x
}

func MergeKLists(lists []*ListNode) *ListNode {
	l := len(lists)
	if l == 0{
		return nil
	}
	mergeListNodes := new(MergeListNodeHeap)

	var firstNode,curMergeListNode *ListNode
	var nextMergeListNode MergeListNode
	overNum := 0

	var isAllNil = true

	for i:=0;i<l;i++{
		if lists[i] != nil {
			*mergeListNodes = append(*mergeListNodes,MergeListNode{
				Val:lists[i].Val,
				From:i,
			})
			//fmt.Printf("list:%v\n",*mergeListNodes)
			isAllNil = false
		}
	}
	if isAllNil{
		return nil
	}
	heap.Init(mergeListNodes)

	firstNode = curMergeListNode
	for overNum < l && len(*mergeListNodes) >0 {
		//sort.Sort(mergeListNodes)
		nextListNode := &ListNode{}
		if curMergeListNode == nil{
			nextListNode = nextListNode
			firstNode = nextListNode
		}else{
			curMergeListNode.Next = nextListNode
		}
		nextMergeListNode = heap.Pop(mergeListNodes).(MergeListNode)
		//fmt.Printf("pop val:%v,mergeListNodes:%v\n",nextMergeListNode,mergeListNodes)
		nextListNode.Val = nextMergeListNode.Val

		curMergeListNode = nextListNode


		from := nextMergeListNode.From

		if lists[from].Next !=nil{
			lists[from] = lists[from].Next
			m := MergeListNode{
				Val: lists[from].Val,
				From: from,
			}
			//fmt.Printf("lists[from]:%v,mergeListNodes:%v,from:%v,m:%v\n",lists[from],mergeListNodes,from,m)
			heap.Push(mergeListNodes,m)
		}else{
			overNum++
		}

	}
	return firstNode

}
