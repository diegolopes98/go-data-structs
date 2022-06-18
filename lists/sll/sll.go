package sll

import "github.com/diegolopes98/go-data-structs/lists"

type node[T any] struct {
	value T
	next  lists.Node[T]
}

func NewNode[T any](value T) lists.Node[T] {
	return &node[T]{value, nil}
}

func (n *node[T]) SetValue(value T) {
	n.value = value
}

func (n *node[T]) GetValue() T {
	return n.value
}

func (n *node[T]) SetNext(node lists.Node[T]) {
	n.next = node
}

func (n *node[T]) GetNext() lists.Node[T] {
	return n.next
}

func (n *node[T]) SetPrevious(node lists.Node[T]) {
	// TODO: review how to prevent empty implementation
}

func (n *node[T]) GetPrevious() lists.Node[T] {
	return nil
}

type list[T any] struct {
	head lists.Node[T]
	tail lists.Node[T]
	len  uint
}

func NewSLL[T any]() lists.List[T] {
	return &list[T]{nil, nil, 0}
}

func (l *list[T]) Head() lists.Node[T] {
	return l.head
}

func (l *list[T]) Tail() lists.Node[T] {
	return l.tail
}

func (l *list[T]) Length() uint {
	return l.len
}

func (l *list[T]) Push(value T) lists.List[T] {
	node := NewNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.SetNext(node)
		l.tail = node
	}
	l.len++
	return l
}

func (l *list[T]) Pop() lists.Node[T] {
	if l.Length() == 0 {
		return nil
	}
	var curr lists.Node[T]
	var prev lists.Node[T]
	curr = l.Head()
	if curr.GetNext() != nil {
		for curr.GetNext() != nil {
			prev = curr
			curr = curr.GetNext()
		}
		prev.SetNext(nil)
	}
	l.tail = prev
	l.len--
	if l.Length() == 0 {
		l.head = nil
		l.tail = nil
	}
	return curr
}

func (l *list[T]) Shift() lists.Node[T] {
	if l.Length() == 0 {
		return nil
	}
	node := l.Head()
	l.head = node.GetNext()
	l.len--
	if l.Length() == 0 {
		l.tail = nil
	}
	return node
}

func (l *list[T]) Unshift(value T) lists.List[T] {
	node := NewNode(value)
	node.SetNext(l.Head())
	l.head = node
	l.len++
	return l
}

func (l *list[T]) Get(index uint) lists.Node[T] {
	if index > l.len-1 {
		return nil
	}
	curr := l.Head()
	for i := uint(0); i < index; i++ {
		curr = curr.GetNext()
	}
	return curr
}

func (l *list[T]) Set(index uint, value T) {
	node := l.Get(index)
	if node != nil {
		node.SetValue(value)
	}
}

func (l *list[T]) Insert(index uint, value T) lists.List[T] {
	if index == 0 {
		l.Unshift(value)
	} else {
		prevNode := l.Get(index - 1)
		if prevNode != nil {
			node := NewNode(value)
			node.SetNext(prevNode.GetNext())
			prevNode.SetNext(node)
			l.len++
		}
	}
	return l
}

func (l *list[T]) Remove(index uint) {
	if index == 0 {
		l.Shift()
	} else if index == l.Length()-1 {
		l.Pop()
	} else {
		prev := l.Get(index - 1)
		if prev != nil {
			curr := prev.GetNext()
			next := curr.GetNext()
			prev.SetNext(next)
			l.len--
		}
	}
}

func (l *list[T]) Reverse() {
	curr := l.head
	l.head = l.tail
	l.tail = curr
	var prev lists.Node[T]
	var next lists.Node[T]
	for curr != nil {
		next = curr.GetNext()
		curr.SetNext(prev)
		prev = curr
		curr = next
	}
}

func ForEach[T any](l lists.List[T], f func(lists.Node[T])) {
	curr := l.Head()
	for curr != nil {
		f(curr)
		curr = curr.GetNext()
	}
}

func Filter[T any](l lists.List[T], f func(T) bool) lists.List[T] {
	nl := NewSLL[T]()
	curr := l.Head()
	for curr != nil {
		if f(curr.GetValue()) {
			nl.Push(curr.GetValue())
		}
		curr = curr.GetNext()
	}
	return nl
}

func Map[T, N any](l lists.List[T], f func(T) N) lists.List[N] {
	nl := NewSLL[N]()
	curr := l.Head()
	for curr != nil {
		nl.Push(f(curr.GetValue()))
		curr = curr.GetNext()
	}
	return nl
}
