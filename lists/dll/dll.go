package dll

import (
	"errors"

	"github.com/diegolopes98/go-data-structs/lists"
)

const (
	emptyerr       = "empty list"
	outofboundserr = "index out of bounds"
)

type dnode[T any] struct {
	value T
	next  *dnode[T]
	prev  *dnode[T]
}

func newnode[T any](value T) *dnode[T] {
	return &dnode[T]{value, nil, nil}
}

type list[T any] struct {
	head *dnode[T]
	tail *dnode[T]
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
	node.prev = l.tail
	if l.len == 0 {
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
	popped := l.tail
	if l.len > 1 {
		l.tail = popped.prev
		l.tail.next = nil
		popped.prev = nil
	} else {
		l.head = nil
		l.tail = nil
	}
	l.len--
	return popped.value, nil
}

func (l *list[T]) Shift() (T, error) {
	if l.len == 0 {
		return *new(T), errors.New(emptyerr)
	}
	node := l.head
	l.head = node.next
	if l.head != nil {
		l.head.prev = nil
	}
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
		l.head.prev = node
		node.next = l.head
		l.head = node
	}
	l.len++
	return l
}

func (l *list[T]) get(index uint) *dnode[T] {
	if l.len == 0 || index > l.len-1 {
		return nil
	}
	headstart := index < l.len/2
	var node *dnode[T]
	if headstart {
		node = l.head
		for i := uint(0); i < index; i++ {
			node = node.next
		}
	} else {
		node = l.tail
		for i := l.len - 1; i > index; i-- {
			node = node.prev
		}
	}
	return node
}

func (l *list[T]) Get(index uint) (T, error) {
	node := l.get(index)
	if node == nil {
		return *new(T), errors.New(outofboundserr)
	}
	return node.value, nil
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
		prev := l.get(index - 1)
		if prev != nil {
			node := newnode(value)
			next := prev.next
			node.next = next
			next.prev = node
			node.prev = prev
			prev.next = node
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
			next.prev = prev
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
	var prev *dnode[T]
	var next *dnode[T]
	for i := uint(0); i < l.len; i++ {
		next = curr.next
		curr.next = prev
		curr.prev = next
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
