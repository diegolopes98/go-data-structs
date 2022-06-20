package dll

import "github.com/diegolopes98/go-data-structs/lists"

type node[T any] struct {
	value T
	next  lists.Node[T]
	prev  lists.Node[T]
}

func NewNode[T any](value T) lists.Node[T] {
	return &node[T]{value, nil, nil}
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
	n.prev = node
}

func (n *node[T]) GetPrevious() lists.Node[T] {
	return n.prev
}

type list[T any] struct {
	head lists.Node[T]
	tail lists.Node[T]
	len  uint
}

// TODO: change every *list[T] return type to lists.List[T] when implementation complete
func NewDLL[T any]() *list[T] {
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

func (l *list[T]) Push(value T) *list[T] {
	node := NewNode(value)
	node.SetPrevious(l.Tail())
	if l.Length() == 0 {
		l.head = node
		l.tail = node
	} else {
		l.Tail().SetNext(node)
		l.tail = node
	}
	l.len++
	return l
}

func (l *list[T]) Pop() lists.Node[T] {
	if l.Length() == 0 {
		return nil
	}
	popped := l.Tail()
	if l.Length() > 1 {
		l.tail = popped.GetPrevious()
		l.Tail().SetNext(nil)
		popped.SetPrevious(nil)
	} else {
		l.head = nil
		l.tail = nil
	}
	l.len--
	return popped
}

func (l *list[T]) Shift() lists.Node[T] {
	if l.Length() == 0 {
		return nil
	}
	shifted := l.Head()
	if l.Length() > 1 {
		l.head = shifted.GetPrevious()
		l.Head().SetPrevious(nil)
		shifted.SetNext(nil)
	} else {
		l.head = nil
		l.tail = nil
	}
	l.len--
	return shifted
}

func (l *list[T]) Unshift(value T) *list[T] {
	node := NewNode(value)
	if l.Length() == 0 {
		l.head = node
		l.tail = node
	} else {
		l.Head().SetPrevious(node)
		node.SetNext(l.head)
		l.head = node
	}
	l.len++
	return l
}

func (l *list[T]) Get(index uint) lists.Node[T] {
	if index > l.Length()-1 {
		return nil
	}
	headstart := index <= l.Length()/2
	var node lists.Node[T]
	if headstart {
		node = l.Head()
		for i := uint(0); i < index; i++ {
			node = node.GetNext()
		}
	} else {
		node = l.Tail()
		for i := l.Length() - 1; i > index; i-- {
			node = node.GetPrevious()
		}
	}
	return node
}
