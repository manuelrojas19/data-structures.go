package queue

import (
	"data-structures/node"
	"errors"
)

type Queue[T any] struct {
	Head *node.Node[T]
	Tail *node.Node[T]
	size int
}

func (queue *Queue[T]) Enqueue(val T) {
	node := node.CreateNode(val)

	if queue.Head == nil && queue.Tail == nil {
		queue.Head = node
		queue.Tail = node
		return
	}

	queue.Tail.Next = node
	queue.Tail = node
	queue.size++
}

func (queue *Queue[T]) Dequeue() (T, error) {
	if queue.size == 0 {
		var zero T
		return zero, errors.New("Queue is empty")

	}
	node := queue.Head
	queue.Head = queue.Head.Next
	queue.size--
	if queue.size == 0 {
		queue.Head = queue.Tail
	}
	return node.Value, nil
}
