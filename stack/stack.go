package stack

import "errors"

type Stack[T any] struct {
	items []T
}

func (stack *Stack[T]) Push(val T) {
	stack.items = append(stack.items, val)
}

func (stack *Stack[T]) Pop(val T) (T, error) {
	if len(stack.items) == 0 {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	item := stack.items[len(stack.items)-1]
	stack.items = stack.items[:len(stack.items)-1]

	return item, nil
}
