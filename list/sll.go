package list

type SLL[T interface{}] interface {
	Head() *Node[T]
	Tail() *Node[T]
	Length() uint
	Push(value T) *List[T]
	Pop() *Node[T]
	Shift() *Node[T]
}

type Node[T interface{}] struct {
	Data T
	Next *Node[T]
}

type List[T interface{}] struct {
	head *Node[T]
	tail *Node[T]
	len  uint
}

func NewSLL[T interface{}]() SLL[T] {
	return &List[T]{nil, nil, 0}
}

func (l *List[T]) Head() *Node[T] {
	return l.head
}

func (l *List[T]) Tail() *Node[T] {
	return l.tail
}

func (l *List[T]) Length() uint {
	return l.len
}

func (l *List[T]) Push(value T) *List[T] {
	node := &Node[T]{value, nil}
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

func (l *List[T]) Pop() *Node[T] {
	if l.len == 0 {
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
	if l.len == 0 {
		l.head = nil
		l.tail = nil
	}
	return curr
}

func (l *List[T]) Shift() *Node[T] {
	if l.len == 0 {
		return nil
	}
	node := l.Head()
	l.head = node.Next
	l.len--
	if l.len == 0 {
		l.tail = nil
	}
	return node
}
