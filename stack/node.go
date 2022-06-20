package stack

type Node[T any] interface {
	GetValue() T
	SetValue(value T)
	GetNext() Node[T]
	SetNext(value Node[T])
}

type node[T any] struct {
	value T
	next  Node[T]
}

func NewNode[T any](value T) Node[T] {
	return &node[T]{value, nil}
}

func (n *node[T]) SetValue(value T) {
	n.value = value
}

func (n *node[T]) GetValue() T {
	return n.value
}

func (n *node[T]) SetNext(node Node[T]) {
	n.next = node
}

func (n *node[T]) GetNext() Node[T] {
	return n.next
}
