package sll

import (
	"errors"

	"github.com/diegolopes98/go-data-structs/lists"
)

const (
	emptyerr       = "empty list"
	outofboundserr = "index out of bounds"
)

type snode[T any] struct {
	value T
	next  *snode[T]
}

func newnode[T any](value T) *snode[T] {
	return &snode[T]{value, nil}
}

type list[T any] struct {
	head *snode[T]
	tail *snode[T]
	len  uint
}

func New[T any]() lists.List[T] {
	return &list[T]{nil, nil, 0}
}

func (l *list[T]) Head() (T, error) {
	if l.len == 0 {
		return *new(T), errors.New(emptyerr)
	}
	return l.head.value, nil
}

func (l *list[T]) Tail() (T, error) {
	if l.len == 0 {
		return *new(T), errors.New(emptyerr)
	}
	return l.tail.value, nil
}

func (l *list[T]) Length() uint {
	return l.len
}

func (l *list[T]) Push(value T) lists.List[T] {
	node := newnode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}
	l.len++
	return l
}

func (l *list[T]) Pop() (T, error) {
	if l.len == 0 {
		return *new(T), errors.New(emptyerr)
	}
	var curr *snode[T]
	var prev *snode[T]
	curr = l.head
	if curr.next != nil {
		for curr.next != nil {
			prev = curr
			curr = curr.next
		}
		prev.next = nil
	}
	l.tail = prev
	l.len--
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	}
	return curr.value, nil
}

func (l *list[T]) Shift() (T, error) {
	if l.len == 0 {
		return *new(T), errors.New(emptyerr)
	}
	node := l.head
	l.head = node.next
	l.len--
	if l.len == 0 {
		l.tail = nil
	}
	return node.value, nil
}

func (l *list[T]) Unshift(value T) lists.List[T] {
	node := newnode(value)
	if l.len == 0 {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head = node
	}
	l.len++
	return l
}

func (l *list[T]) Get(index uint) (T, error) {
	node := l.get(index)
	if node == nil {
		return *new(T), errors.New(outofboundserr)
	}
	return node.value, nil
}

func (l *list[T]) get(index uint) *snode[T] {
	if l.len == 0 || index > l.len-1 {
		return nil
	}
	curr := l.head
	for i := uint(0); i < index; i++ {
		curr = curr.next
	}
	return curr
}

func (l *list[T]) Set(index uint, value T) {
	node := l.get(index)
	if node != nil {
		node.value = value
	}
}

func (l *list[T]) Insert(index uint, value T) lists.List[T] {
	if index == 0 {
		l.Unshift(value)
	} else {
		prevNode := l.get(index - 1)
		if prevNode != nil {
			node := newnode(value)
			node.next = prevNode.next
			prevNode.next = node
			l.len++
		}
	}
	return l
}

func (l *list[T]) Remove(index uint) (T, error) {
	if index == 0 {
		return l.Shift()
	} else if index == l.len-1 {
		return l.Pop()
	} else {
		prev := l.get(index - 1)
		if prev != nil {
			curr := prev.next
			next := curr.next
			prev.next = next
			l.len--
			return curr.value, nil
		}
	}
	return *new(T), errors.New(outofboundserr)
}

func (l *list[T]) Reverse() {
	curr := l.head
	l.head = l.tail
	l.tail = curr
	var prev *snode[T]
	var next *snode[T]
	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
}

func (l *list[T]) ForEach(f func(*T)) {
	curr := l.head
	for curr != nil {
		f(&curr.value)
		curr = curr.next
	}
}
