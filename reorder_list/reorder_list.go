package reorder_list

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (h *ListNode) Print() {
	for h != nil {
		fmt.Printf("%d ", h.Val)
		h = h.Next
	}
	fmt.Printf("\n")
}

func NewListNode(num int) *ListNode {
	return &ListNode{num, nil}
}

func NewList(nums ...int) *ListNode {
	head := NewListNode(nums[0])
	temp := head
	for i := 1; i < len(nums); i++ {
		temp.Next = NewListNode(nums[i])
		temp = temp.Next
	}
	return head
}

type Stack struct {
	stack []*ListNode
}

func NewStack() Stack {
	return Stack{[]*ListNode{}}
}

func (s *Stack) Len() int {
	return len(s.stack)
}

func (s *Stack) MustTop() *ListNode {
	return s.stack[len(s.stack)-1]
}

func (s *Stack) MustPop() {
	s.stack = s.stack[:len(s.stack)-1]
}

func (s *Stack) MustPush(i *ListNode) {
	s.stack = append(s.stack, i)
}

func ReorderList(head *ListNode) {
	if head == nil {
		return
	}

	var slice []*ListNode
	tmp := head
	for tmp != nil {
		slice = append(slice, tmp)
		tmp = tmp.Next
	}

	l, r := 0, len(slice)-1
	for l < r {
		slice[r].Next = nil
		slice[l].Next = slice[r]
		l++
		slice[r].Next = slice[l]
		slice[l].Next = nil
		r--
	}

	head = slice[0]
}
