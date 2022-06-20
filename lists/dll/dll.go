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

func NewDLL[T any]() lists.List[T] {
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

func (l *list[T]) Unshift(value T) lists.List[T] {
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
		prev := l.Get(index - 1)
		if prev != nil {
			node := NewNode(value)
			next := prev.GetNext()
			node.SetNext(next)
			next.SetPrevious(node)
			node.SetPrevious(prev)
			prev.SetNext(node)
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
			next.SetPrevious(prev)
			l.len--
		}
	}
}

func (l *list[T]) Reverse() {
	curr := l.Head()
	l.head = l.Tail()
	l.tail = curr
	var prev lists.Node[T]
	var next lists.Node[T]
	for i := uint(0); i < l.Length(); i++ {
		next = curr.GetNext()
		curr.SetNext(prev)
		curr.SetPrevious(next)
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
	nl := NewDLL[T]()
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
	nl := NewDLL[N]()
	curr := l.Head()
	for curr != nil {
		nl.Push(f(curr.GetValue()))
		curr = curr.GetNext()
	}
	return nl
}
