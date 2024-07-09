package list

import (
	"data-structures/node"
)

type SingleLinkedList[T any] struct {
	size int
	Head *node.Node[T]
}

func createSingleLinkedList[T any]() *SingleLinkedList[T] {
	return &SingleLinkedList[T]{}
}

func (list *SingleLinkedList[T]) AddAtBeg(val T) {
	node := node.CreateNode(val)
	node.Next = list.Head
	list.Head = node
}

func (list *SingleLinkedList[T]) AddAtEnd(val T) {
	node := node.CreateNode(val)

	if list.Head == nil {
		list.Head = node
		list.size++
		return
	}

	current := list.Head
	for current.Next != nil {
	}
	current.Next = node
	list.size++
}

func (list *SingleLinkedList[T]) DeleteAtEnd() bool {
	if list.Head == nil {
		return false
	}

	current := list.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = nil
	list.size--
	return true
}
