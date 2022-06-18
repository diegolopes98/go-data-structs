package list

type SLL[T interface{}] interface {
	Head() *Node[T]
	Tail() *Node[T]
	Length() uint
	Push(value T) SLL[T]
	Pop() *Node[T]
	Shift() *Node[T]
	Unshift(value T) SLL[T]
	Get(index uint) *Node[T]
	Set(index uint, value T)
	Insert(index uint, value T) SLL[T]
	Remove(index uint)
	Reverse()
}

type Node[T interface{}] struct {
	Value T
	Next  *Node[T]
}

func NewNode[T interface{}](value T) *Node[T] {
	return &Node[T]{value, nil}
}

type list[T interface{}] struct {
	head *Node[T]
	tail *Node[T]
	len  uint
}

func NewSLL[T interface{}]() SLL[T] {
	return &list[T]{nil, nil, 0}
}

func (l *list[T]) Head() *Node[T] {
	return l.head
}

func (l *list[T]) Tail() *Node[T] {
	return l.tail
}

func (l *list[T]) Length() uint {
	return l.len
}

func (l *list[T]) Push(value T) SLL[T] {
	node := NewNode(value)
	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.Next = node
		l.tail = node
	}
	l.len++
	return l
}

func (l *list[T]) Pop() *Node[T] {
	if l.Length() == 0 {
		return nil
	}
	var curr *Node[T]
	var prev *Node[T]
	curr = l.head
	if curr.Next != nil {
		for curr.Next != nil {
			prev = curr
			curr = curr.Next
		}
		prev.Next = nil
	}
	l.tail = prev
	l.len--
	if l.Length() == 0 {
		l.head = nil
		l.tail = nil
	}
	return curr
}

func (l *list[T]) Shift() *Node[T] {
	if l.Length() == 0 {
		return nil
	}
	node := l.Head()
	l.head = node.Next
	l.len--
	if l.Length() == 0 {
		l.tail = nil
	}
	return node
}

func (l *list[T]) Unshift(value T) SLL[T] {
	node := NewNode(value)
	node.Next = l.Head()
	l.head = node
	l.len++
	return l
}

func (l *list[T]) Get(index uint) *Node[T] {
	if index > l.len-1 {
		return nil
	}
	curr := l.Head()
	for i := uint(0); i < index; i++ {
		curr = curr.Next
	}
	return curr
}

func (l *list[T]) Set(index uint, value T) {
	node := l.Get(index)
	if node != nil {
		node.Value = value
	}
}

func (l *list[T]) Insert(index uint, value T) SLL[T] {
	if index == 0 {
		l.Unshift(value)
	} else {
		prevNode := l.Get(index - 1)
		if prevNode != nil {
			node := NewNode(value)
			node.Next = prevNode.Next
			prevNode.Next = node
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
			curr := prev.Next
			next := curr.Next
			prev.Next = next
			l.len--
		}
	}
}

func (l *list[T]) Reverse() {
	curr := l.head
	l.head = l.tail
	l.tail = curr
	var prev *Node[T]
	var next *Node[T]
	for curr != nil {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
}

func ForEach[T interface{}](l SLL[T], f func(*Node[T])) {
	curr := l.Head()
	for curr != nil {
		f(curr)
		curr = curr.Next
	}
}

func Filter[T interface{}](l SLL[T], f func(T) bool) SLL[T] {
	nl := NewSLL[T]()
	curr := l.Head()
	for curr != nil {
		if f(curr.Value) {
			nl.Push(curr.Value)
		}
		curr = curr.Next
	}
	return nl
}

func Map[T, N interface{}](l SLL[T], f func(T) N) SLL[N] {
	nl := NewSLL[N]()
	curr := l.Head()
	for curr != nil {
		nl.Push(f(curr.Value))
		curr = curr.Next
	}
	return nl
}
