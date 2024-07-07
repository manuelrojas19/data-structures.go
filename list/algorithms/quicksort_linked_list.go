package algorithms

import (
	"data-structures/node"
)

func QuickSort(head *node.Node[int]) *node.Node[int] {
	if head == nil || head.Next == nil {
		return head
	}

	pivot := head
	head = head.Next
	pivot.Next = nil

	lessHead := &node.Node[int]{}
	lessTail := lessHead
	greaterHead := &node.Node[int]{}
	greaterTail := greaterHead

	for head != nil {
		if head.Value < pivot.Value {
			lessTail.Next = head
			lessTail = lessTail.Next
		} else {
			greaterTail.Next = head
			greaterTail = greaterTail.Next
		}
		head = head.Next
	}

	lessTail.Next = nil
	greaterTail.Next = nil

	sortedLess := QuickSort(lessHead.Next)
	sortedGreat := QuickSort(greaterHead.Next)

	if sortedLess == nil {
		pivot.Next = sortedGreat
		return pivot
	} else {
		current := sortedLess
		for current.Next != nil {
			current = current.Next
		}
		current.Next = pivot
		pivot.Next = sortedGreat
		return sortedLess
	}
}
