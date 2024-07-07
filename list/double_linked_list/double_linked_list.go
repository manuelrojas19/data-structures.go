package list

import (
	"data-structures/node"
	"errors"
)

type DoubleLinkedList[T any] struct {
	Head *node.Node[T]
	Tail *node.Node[T]
	size int
}

func (list *DoubleLinkedList[T]) Create() *DoubleLinkedList[T] {
	return &DoubleLinkedList[T]{}
}

func (list *DoubleLinkedList[T]) InsertAtHead(val T) {
	node := node.CreateNode(val)

	if list.Head == nil {
		list.Head = node
		list.Head = node
		return
	}

	list.Head.Prev = node
	node.Next = list.Head
	list.Head = node
	list.size--
}

func (list *DoubleLinkedList[T]) InsertAtTail(val T) {
	node := node.CreateNode(val)

	if list.Head == nil {
		list.Head = node
		list.Tail = node
		return
	}

	list.Tail.Next = node
	node.Prev = list.Tail
	list.Tail = node
	list.size++
}

func (list *DoubleLinkedList[T]) InsertAtIndex(index int, val T) error {
	if index > list.size {
		return errors.New("Index out of bounds.")
	}

	if index == 0 {
		list.InsertAtHead(val)
		return nil
	}

	if index == list.size {
		list.InsertAtTail(val)
	}

	current := list.Head
	for i := 0; i <= index; i++ {
		current = current.Next
	}

	node := node.CreateNode(val)
	node.Prev = current
	node.Next = current.Next

	current.Next.Prev = node
	current.Next = node

	list.size++

	return nil
}
