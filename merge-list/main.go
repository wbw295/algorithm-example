package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {

	var dummyNode *ListNode = &ListNode{}
	tail := dummyNode
	l, r := list1, list2
	for l != nil && r != nil {
		if l.Val < r.Val {
			tail.Next = &ListNode{Val: l.Val}
			l = l.Next
		} else {
			tail.Next = &ListNode{Val: r.Val}
			r = r.Next
		}
		tail = tail.Next
	}
	if l != nil {
		tail.Next = l
	}
	if r != nil {
		tail.Next = r
	}
	return dummyNode.Next
}

func main() {
	l1 := &ListNode{Val: 1}
	l1.Next = &ListNode{Val: 3}
	l1.Next = &ListNode{Val: 5}
	l1.Next = &ListNode{Val: 7}
	l2 := &ListNode{Val: 2}
	l2.Next = &ListNode{Val: 4}
	l2.Next = &ListNode{Val: 6}
	l2.Next = &ListNode{Val: 8}

	list := mergeTwoLists(l1, l2)
	fmt.Println(list)

}
